package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
)

func TestSuccessEvmAdapterMultisigWalletTestMethods(t *testing.T) {
	stub, state, ev := GetEvmAdapterAndMockStub("msp1", TestUser1ClientCert, TestUser1ClientAddr)

	contractBytecode := LoadBytecodeFromFile("./bin/compilationTests/MultiSigWallet/MultiSigWalletFactory.bin")

	stub.MockTransactionStart("tx1")
	contractAcc, err := ev.InstallContractByOwner(contractBytecode, []byte{})
	stub.MockTransactionEnd("tx1")

	require.NoError(t, err)
	assert.NotNil(t, contractAcc.Code)

	// Call create multisigwallet with 2 owners by multisigwallet factory
	result, err := CallEvmMethod(stub, ev, contractAcc, 0,
		"create(address[],uint256)",
		word256.Int64ToWord256(32*2).Bytes(),
		word256.Int64ToWord256(2).Bytes(),
		word256.Int64ToWord256(2).Bytes(),
		TestUser1ClientAddr.Bytes(), TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	require.NotEqual(t, word256.Zero256, result)
	require.Equal(t, 32, len(result))

	// Retrive contract account with new code
	contractAcc = state.GetAccount(word256.LeftPadWord256(result))
	contractAcc.Balance = 100
	stub.MockTransactionStart("tx")
	state.UpdateAccount(contractAcc)
	stub.MockTransactionEnd("tx")

	// Get owners
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"getOwners()")
	require.NoError(t, err)
	assert.Equal(t, JoinBytesArgs(
		word256.Int64ToWord256(32).Bytes(),
		word256.Int64ToWord256(2).Bytes(),
		TestUser1ClientAddr.Bytes(), TestUser2ClientAddr.Bytes()), result)

	// Submit and confirm transaction by user1
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"submitTransaction(address,uint256,bytes)",
		TestUser3ClientAddr.Bytes(),
		word256.Int64ToWord256(55).Bytes(),
		word256.Int64ToWord256(32*3).Bytes(),
		word256.Int64ToWord256(4).Bytes(),
		word256.RightPadWord256([]byte("TEST")).Bytes())
	require.NoError(t, err)
	// First transactionId = 0
	require.Equal(t, word256.Zero256.Bytes(), result)

	// Check confirmation count
	// First transactionId = 0
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"getConfirmationCount(uint256)",
		word256.Int64ToWord256(0).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(1).Bytes(), result)

	// Approve by user2, and execute transaction
	stub.SetCreator("msp1", TestUser2ClientCert)
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"confirmTransaction(uint256)",
		word256.Int64ToWord256(0).Bytes())
	require.NoError(t, err)
	assert.Equal(t, 0, len(result))

	// Check confirmation count
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"getConfirmationCount(uint256)",
		word256.Int64ToWord256(0).Bytes())
	require.NoError(t, err)
	assert.Equal(t, word256.Int64ToWord256(2).Bytes(), result)

	// Check balances
	assert.Equal(t, int64(100-55), state.GetAccount(contractAcc.Address).Balance)
	assert.Equal(t, int64(300+55), state.GetAccount(TestUser3ClientAddr).Balance)
}

func TestFailEvmAdapterMultisigWalletWrongArgs(t *testing.T) {
	stub, state, ev := GetEvmAdapterAndMockStub("msp1", TestUser1ClientCert, TestUser1ClientAddr)

	contractBytecode := LoadBytecodeFromFile("./bin/compilationTests/MultiSigWallet/MultiSigWalletFactory.bin")

	stub.MockTransactionStart("tx1")
	contractAcc, err := ev.InstallContractByOwner(contractBytecode, []byte{
		1, 2, 3, 4, 5, // ignored
	})
	stub.MockTransactionEnd("tx1")

	require.NoError(t, err)
	assert.NotNil(t, contractAcc.Code)

	// Call create multisigwallet with 2 owners by multisigwallet factory
	_, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"create(address[],uint256)",
		word256.Int64ToWord256(2).Bytes(),
		word256.Int64ToWord256(32).Bytes(),
		word256.Int64ToWord256(2).Bytes(),
		word256.Int64ToWord256(2).Bytes(),
		TestUser1ClientAddr.Bytes(), TestUser2ClientAddr.Bytes())
	require.EqualError(t, err, "Insufficient gas")

	// Call create multisigwallet with 2 owners by multisigwallet factory
	result, err := CallEvmMethod(stub, ev, contractAcc, 0,
		"create(address[],uint256)",
		word256.Int64ToWord256(32*2).Bytes(),
		word256.Int64ToWord256(2).Bytes(),
		word256.Int64ToWord256(2).Bytes(),
		TestUser1ClientAddr.Bytes(), TestUser2ClientAddr.Bytes())
	require.NoError(t, err)
	require.NotEqual(t, word256.Zero256, result)
	require.Equal(t, 32, len(result))

	// Retrive contract account with new code
	contractAcc = state.GetAccount(word256.LeftPadWord256(result))
	stub.MockTransactionStart("tx")
	state.UpdateAccount(contractAcc)
	stub.MockTransactionEnd("tx")

	// Submit and confirm transaction by user1
	result, err = CallEvmMethod(stub, ev, contractAcc, 0,
		"submitTransaction(address,uint256,bytes)",
		TestUser3ClientAddr.Bytes(),
		word256.Int64ToWord256(55).Bytes(),
		word256.Int64ToWord256(16).Bytes(),
		word256.Int64ToWord256(4).Bytes(),
		word256.RightPadWord256([]byte("TEST")).Bytes())
	require.EqualError(t, err, "Insufficient gas")
}
