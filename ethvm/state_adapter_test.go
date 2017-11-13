package ethvm

import (
	"encoding/base64"
	"testing"

	evm "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm"
	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var accountAddress1Word256 = word256.RightPadWord256([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
var accountAddress1Str = "AQAAAAAAAAAAAAAAAAAAAAAAAAA="
var accountAddress2Word256 = word256.RightPadWord256([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2})
var accountAddress2Str = "AgAAAAAAAAAAAAAAAAAAAAAAAAA="
var accountAddress3Word256 = word256.RightPadWord256([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3})
var accountAddress3Str = "AwAAAAAAAAAAAAAAAAAAAAAAAAA="

func getStubWithAccounts() (*shim.MockStub, *EvmStateAdapter) {
	stub := shim.NewMockStub("hyperledger-fabric-evmcc", &ChainCode{})
	evmStateAdapter := NewEvmStateAdapter(stub)
	data, _ := MarshalEvmAccountToMsgPack(&evm.Account{
		Address: accountAddress1Word256,
		Balance: 1000,
	})
	stub.State["evm:"+accountAddress1Str] = data
	data, _ = MarshalEvmAccountToMsgPack(&evm.Account{
		Address: accountAddress2Word256,
		Balance: 2000,
	})
	stub.State["evm:"+accountAddress2Str] = data
	data, _ = MarshalEvmAccountToMsgPack(&evm.Account{
		Address: accountAddress3Word256,
		Balance: 3000,
	})
	stub.State["evm:"+accountAddress3Str] = data
	return stub, evmStateAdapter
}

func TestSuccessEvmStateAdapterCreateEvmAccountAsset(t *testing.T) {
	stub := shim.NewMockStub("hyperledger-fabric-evmcc", &ChainCode{})
	evmStateAdapter := NewEvmStateAdapter(stub)

	require.NotNil(t, evmStateAdapter)

	account := evmStateAdapter.CreateAccount(&evm.Account{
		Address: word256.One256,
		Nonce:   1,
	})

	require.NotNil(t, account)
	assert.Equal(t, []byte{
		0xc4, 0x6a, 0x48, 0xc9, 0x8d, 0xc0, 0xe5, 0x8b, 0xb1, 0x1f,
		0x27, 0x8c, 0x1b, 0x8b, 0x49, 0x12, 0x27, 0xdf, 0x0b, 0x86,
	}, account.Address[12:])
	assert.Equal(t, int64(2), account.Nonce)
	assert.Equal(t, int64(0), account.Balance)
	assert.Nil(t, account.Code)
}

func TestSuccessEvmStateAdapterSaveNewEvmAccountAsset(t *testing.T) {
	stub, evmStateAdapter := getStubWithAccounts()

	account := evmStateAdapter.CreateAccount(&evm.Account{
		Address: word256.One256,
		Balance: 1000,
		Nonce:   0,
	})
	account.Balance = 1234

	stub.MockTransactionStart("tx1")
	evmStateAdapter.UpdateAccount(account)
	stub.MockTransactionEnd("tx1")

	data := stub.State["evm:HDMBqe0726eOKHwTjXvmMUdgR2k="]
	require.NotNil(t, data)

	account, err := UnmarshalEvmAccountFromMsgPack(data)
	require.NoError(t, err)
	assert.Equal(t, int64(1234), account.Balance)
	assert.Equal(t, int64(1), account.Nonce)
}

func TestSuccessEvmStateAdapterGetAccount(t *testing.T) {
	_, evmStateAdapter := getStubWithAccounts()
	account := evmStateAdapter.GetAccount(accountAddress1Word256)
	require.NotNil(t, account)
	assert.Equal(t, account.Address, accountAddress1Word256)
}

func TestNotFoundEvmStateAdapterGetAccount(t *testing.T) {
	_, evmStateAdapter := getStubWithAccounts()
	assert.PanicsWithValue(t,
		"Get account by addr: 0000000000000000000000000000000000000000000000000000000000000000, with error <nil>",
		func() {
			evmStateAdapter.GetAccount(word256.Zero256)
		})
}

func TestSuccessEvmStateAdapterUpdateAccount(t *testing.T) {
	stub, evmStateAdapter := getStubWithAccounts()
	account := evmStateAdapter.GetAccount(accountAddress1Word256)
	require.NotNil(t, account)
	account.Balance = 1234

	stub.MockTransactionStart("tx1")
	evmStateAdapter.UpdateAccount(account)
	stub.MockTransactionEnd("tx1")
	account = evmStateAdapter.GetAccount(accountAddress1Word256)
	require.NotNil(t, account)
	assert.NotNil(t, 1234, account.Balance)
}

func TestSuccessEvmStateAdapterRemoveAccount(t *testing.T) {
	stub, evmStateAdapter := getStubWithAccounts()
	account := evmStateAdapter.GetAccount(accountAddress1Word256)
	require.NotNil(t, account)
	assert.NotEqual(t, 0, len(stub.State["evm:"+accountAddress1Str]))

	evmStateAdapter.RemoveAccount(account)

	assert.Equal(t, 0, len(stub.State["evm:"+accountAddress1Str]))
}

func TestNotExistsEvmStateAdapterRemoveAccount(t *testing.T) {
	_, evmStateAdapter := getStubWithAccounts()

	assert.PanicsWithValue(t, "Remove account with addr: 0000000000000000000000000000000000000000000000000000000000000000, error <nil>", func() {
		evmStateAdapter.RemoveAccount(&evm.Account{
			Address: word256.Zero256,
		})
	})
}

func TestSuccessEvmAccountGetStorage(t *testing.T) {
	stub, evmStateAdapter := getStubWithAccounts()
	stub.State["evm:"+accountAddress2Str+":s:"+base64.StdEncoding.EncodeToString(word256.One256.Bytes())] = word256.One256.Bytes()

	value := evmStateAdapter.GetStorage(accountAddress2Word256, word256.One256)

	assert.Equal(t, word256.One256, value)
}

func TestNotExistsEvmAccountGetStorage(t *testing.T) {
	_, evmStateAdapter := getStubWithAccounts()

	assert.NotPanics(t, func() {
		evmStateAdapter.GetStorage(word256.Zero256, word256.One256)
	})
}

func TestSuccessEvmAccountSetStorage(t *testing.T) {
	stub, evmStateAdapter := getStubWithAccounts()

	stub.MockTransactionStart("tx1")
	evmStateAdapter.SetStorage(accountAddress2Word256, word256.One256, word256.One256)
	stub.MockTransactionEnd("tx1")

	assert.Equal(t, word256.One256.Bytes(), stub.State["evm:"+accountAddress2Str+":s:"+base64.StdEncoding.EncodeToString(word256.One256.Bytes())])
}

func TestNotExistsEvmAccountSetStorage(t *testing.T) {
	stub, evmStateAdapter := getStubWithAccounts()

	assert.NotPanics(t, func() {
		stub.MockTransactionStart("tx1")
		evmStateAdapter.SetStorage(word256.Zero256, word256.One256, word256.One256)
	})
}
