package tests

import (
	// "encoding/json"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
)

func TestSuccessEvmAdapterJcrTokenTestMethods(t *testing.T) {
	stub, _, ev := GetEvmAdapterAndMockStub("msp1", TestUser1ClientCert, TestUser1ClientAddr)

	contractBytecode := LoadBytecodeFromFile("./bin/jincor/JincorToken.bin")

	stub.MockTransactionStart("tx1")
	contractAcc, err := ev.InstallContractByOwner(contractBytecode, JoinBytesArgs())
	stub.MockTransactionEnd("tx1")

	require.NoError(t, err)
	assert.NotNil(t, contractAcc.Code)

	// Call getBalance for user1
	result, err := CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser1ClientAddr.Bytes())
	require.NoError(t, err)
	// 35000000000000000000000000L
	assert.Equal(t, word256.LeftPadWord256([]byte{0x1c, 0xf3, 0x89, 0xcd, 0x46,
		0x04, 0x7d, 0x03, 0x00, 0x00, 0x00}).Bytes(), result)

	// Call getBalance for user2
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(0).Bytes(), result)

	// Set release agent user1
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"setReleaseAgent(address)", TestUser1ClientAddr.Bytes())
	require.NoError(t, err)

	// Set transfer agent user1
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"setTransferAgent(address,bool)", TestUser1ClientAddr.Bytes(), word256.Int64ToWord256(1).Bytes())
	require.NoError(t, err)

	// Transfer by user1 to user2 2000000
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"transfer(address,uint256)", TestUser2ClientAddr.Bytes(), word256.Int64ToWord256(200000).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Call getBalance for user2
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(200000).Bytes(), result)

	// Approve for user3 50000
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"approve(address,uint256)", TestUser3ClientAddr.Bytes(), word256.Int64ToWord256(50000).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Check approved value
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"allowance(address,address)", TestUser1ClientAddr.Bytes(), TestUser3ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(50000).Bytes(), result)

	// transfer approved from user1 by user3
	stub.SetCreator("msp1", TestUser3ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"transferFrom(address,address,uint256)",
		TestUser1ClientAddr.Bytes(), TestUser3ClientAddr.Bytes(),
		word256.Int64ToWord256(50000).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// check user3 balance
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser3ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(50000).Bytes(), result)

	// set released
	stub.SetCreator("msp1", TestUser1ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"release()")
	require.NoError(t, err)

	// transfer user2 to user3 10000
	stub.SetCreator("msp1", TestUser2ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"transfer(address,uint256)", TestUser3ClientAddr.Bytes(), word256.Int64ToWord256(10000).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// check user3 balance
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser3ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(60000).Bytes(), result)

	// burn user3 tokens
	stub.SetCreator("msp1", TestUser3ClientCert)
	// approve tokens for user1
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"approve(address,uint256)", TestUser1ClientAddr.Bytes(), word256.Int64ToWord256(5000).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// burn user's 3 tokens by owner user1
	stub.SetCreator("msp1", TestUser1ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"burnFrom(address,uint256)", TestUser3ClientAddr.Bytes(),
		word256.Int64ToWord256(5000).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// check user3 balance
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"balanceOf(address)", TestUser3ClientAddr.Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(55000).Bytes(), result)
}
