package tests

import (
	// "encoding/json"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
)

func TestSuccessEvmAdapterStdTokenTestMethods(t *testing.T) {
	stub, _, ev := GetEvmAdapterAndMockStub("msp1", TestUser1ClientCert, TestUser1ClientAddr)

	contractBytecode := LoadBytecodeFromFile("./bin/std/StandardToken.bin")

	stub.MockTransactionStart("tx1")
	contractAcc, err := ev.InstallContractByOwner(contractBytecode, JoinBytesArgs(
		TestUser1ClientAddr.Bytes(), word256.Int64ToWord256(20000).Bytes()))
	stub.MockTransactionEnd("tx1")

	require.NoError(t, err)
	assert.NotNil(t, contractAcc.Code)

	// Call getBalance for user1
	result, err := CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser1ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(20000).Bytes(), result)

	// Call getBalance for user2
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(0).Bytes(), result)

	// Call transfer by super user
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"transfer(address,uint256)", TestUser2ClientAddr.Bytes(), word256.Int64ToWord256(200).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Call getBalance for user2
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(200).Bytes(), result)

	// Approve transfer from user2 -> user1
	stub.SetCreator("msp1", TestUser2ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"approve(address,uint256)",
		TestUser3ClientAddr.Bytes(), word256.Int64ToWord256(100).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Transfer from .. to for not approved
	stub.SetCreator("msp1", TestUser3ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"transferFrom(address,address,uint256)",
		TestUser3ClientAddr.Bytes(), TestUser2ClientAddr.Bytes(), word256.Int64ToWord256(101).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(0).Bytes(), result)

	// Transfer from .. to too many than approved
	stub.SetCreator("msp1", TestUser3ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"transferFrom(address,address,uint256)",
		TestUser2ClientAddr.Bytes(), TestUser3ClientAddr.Bytes(), word256.Int64ToWord256(101).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(0).Bytes(), result)

	// Transfer from .. to that approved
	stub.SetCreator("msp1", TestUser3ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"transferFrom(address,address,uint256)",
		TestUser2ClientAddr.Bytes(), TestUser3ClientAddr.Bytes(), word256.Int64ToWord256(100).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

}
