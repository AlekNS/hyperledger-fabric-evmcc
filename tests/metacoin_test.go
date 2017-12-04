package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ethvm "github.com/JincorTech/hyperledger-fabric-evmcc/ethvm"
	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
)

func TestSuccessEvmAdapterMetaCoinInstallContract(t *testing.T) {
	stub, _, ev := GetEvmAdapterAndMockStub("msp1", TestUser1ClientCert, TestUser1ClientAddr)

	contractBytecode := LoadBytecodeFromFile("./bin/std/MetaCoin.bin")

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

func TestSuccessEvmAdapterMetaCoinTestMethods(t *testing.T) {
	stub, _, ev := GetEvmAdapterAndMockStub("msp1", TestUser1ClientCert, TestUser1ClientAddr)

	contractBytecode := LoadBytecodeFromFile("./bin/std/MetaCoin.bin")

	stub.MockTransactionStart("tx1")
	// contract initiate with 10000 coins
	contractAcc, err := ev.InstallContractByOwner(contractBytecode, []byte{})
	stub.MockTransactionEnd("tx1")

	require.NoError(t, err)
	assert.NotNil(t, contractAcc.Code)

	// Send coins
	result, err := CallEvmMethod(stub, ev, contractAcc, 0,
		"sendCoin(address,uint256)", TestUser2ClientAddr.Bytes(), word256.Int64ToWord256(45).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Catch events
	eventData, ok := stub.Events["EVM:LOG"]
	require.True(t, ok)

	eventLogVal := ethvm.EvmDataLogEvent{}
	require.NoError(t, json.Unmarshal(eventData[0], &eventLogVal))
	assert.Empty(t, eventLogVal.Error)
	assert.Equal(t, contractAcc.Address.Hex()[24:], eventLogVal.ContractAddress)
	assert.Equal(t, word256.Int64ToWord256(45).Bytes(), eventLogVal.Payload.Data)

	// Call getBalance for user1
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"getBalance(address)", TestUser1ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(10000-45).Bytes(), result)

	// Call getBalance for user2
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"getBalance(address)", TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(45).Bytes(), result)

	// Call getBalance for user3
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"getBalance(address)", TestUser3ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(0).Bytes(), result)
}
