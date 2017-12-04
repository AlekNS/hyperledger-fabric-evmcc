package tests

import (
	"encoding/json"
	"testing"
	//"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/txs"
	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
)

func TestSuccessEvmAdapterPayrollInstallContract(t *testing.T) {
	stub, _, ev := GetEvmAdapterAndMockStub("msp1", TestUser1ClientCert, TestUser1ClientAddr)

	contractBytecode := LoadBytecodeFromFile("./bin/other/Payroll.bin")

	stub.MockTransactionStart("tx1")
	contractAcc1, err := ev.InstallContractByOwner(contractBytecode, []byte{})
	stub.MockTransactionEnd("tx1")

	require.NoError(t, err)
	assert.Equal(t, int64(1), contractAcc1.Nonce)
	assert.NotNil(t, contractAcc1.Code)

	stub.MockTransactionStart("tx2")
	contractAcc2, err := ev.InstallContractByOwner(contractBytecode, []byte{})
	stub.MockTransactionEnd("tx2")

	require.NoError(t, err)
	assert.Equal(t, int64(2), contractAcc2.Nonce)
	assert.NotNil(t, contractAcc2.Code)

	assert.NotEqual(t, contractAcc1.Address, contractAcc2.Address)
}

func TestSuccessEvmAdapterPayrollTestMethods(t *testing.T) {
	stub, state, ev := GetEvmAdapterAndMockStub("msp1", TestUser1ClientCert, TestUser1ClientAddr)

	contractBytecode := LoadBytecodeFromFile("./bin/other/Payroll.bin")

	stub.MockTransactionStart("tx1")
	// contract initiate with 10000 coins
	contractAcc, err := ev.InstallContractByOwner(contractBytecode, []byte{})
	stub.MockTransactionEnd("tx1")

	require.NoError(t, err)
	assert.NotNil(t, contractAcc.Code)

	// Retrive contract account with new code
	contractAcc = state.GetAccount(contractAcc.Address)
	contractAcc.Balance = 1000
	stub.MockTransactionStart("tx")
	state.UpdateAccount(contractAcc)
	stub.MockTransactionEnd("tx")

	// Get employee id
	result, err := CallEvmMethod(stub, ev, contractAcc, 0,
		"getEmployeeIdByAddress(address)", TestUser1ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Get employee comments
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"getEmployeeCommentsByAddress(address)", TestUser1ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(32).Bytes(), result[:32])
	assert.Equal(t, word256.Int64ToWord256(5).Bytes(), result[32:64])
	assert.Equal(t, word256.RightPadWord256([]byte("Owner")).Bytes(), result[64:96])

	// Add employee
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"addEmployee(address,uint256,string)", TestUser2ClientAddr.Bytes(),
		word256.Int64ToWord256(100).Bytes(),
		word256.Int64ToWord256(0x60).Bytes(),
		word256.Int64ToWord256(8).Bytes(),
		word256.RightPadWord256([]byte("Designer")).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Catch second events
	eventData, ok := stub.Events["EVM:LOG:"+contractAcc.Address.Hex()[24:]+":OK"]
	require.True(t, ok)
	eventLogVal := txs.EventDataLog{}
	require.NoError(t, json.Unmarshal(eventData[1], &eventLogVal))

	// Check second event (first fired in the constructor)
	assert.Equal(t, GetKeccakHashFromString("onAddEmployee(address,uint256,string)"), eventLogVal.Topics[0].Bytes())
	assert.Equal(t, TestUser2ClientAddr, eventLogVal.Topics[1])
	assert.Equal(t, word256.Int64ToWord256(2).Bytes(), eventLogVal.Data[:32])
	assert.Equal(t, word256.Int64ToWord256(0x40).Bytes(), eventLogVal.Data[32:64])
	assert.Equal(t, word256.Int64ToWord256(8).Bytes(), eventLogVal.Data[64:96])
	assert.Equal(t, word256.RightPadWord256([]byte("Designer")).Bytes(), eventLogVal.Data[96:128])

	// Test try to withdraw something by user2
	stub.SetCreator("msp1", TestUser2ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"withdrawPayroll()")
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(0).Bytes(), result)

	// Pay first time (wages=100)
	stub.SetCreator("msp1", TestUser1ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"payEmployee(address)", TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Change wages to 50 for user2
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"modifyEmployeeWages(address,uint256)", TestUser2ClientAddr.Bytes(),
		word256.Int64ToWord256(50).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Pay again (wages=50)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"payEmployee(address)", TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Try again to withdraw something by user2
	stub.SetCreator("msp1", TestUser2ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"withdrawPayroll()")
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Check balances of user2
	userAcc := state.GetAccount(TestUser2ClientAddr)
	require.NoError(t, err)
	assert.Equal(t, int64(200+150), userAcc.Balance)

	contractAcc = state.GetAccount(contractAcc.Address)
	assert.Equal(t, int64(1000-150), contractAcc.Balance)
}
