package main

import (
	"encoding/json"
	"encoding/hex"
	"testing"

	evm "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm"

	"github.com/JincorTech/hyperledger-fabric-evmcc/tests"
	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hyperledger/fabric/bccsp/factory"
)

func TestInit(t *testing.T) {
	factory.InitFactories(nil)
	stub := shim.NewMockStub("hyperledger-fabric-evmcc", &JincorBaseChaincode{})
	response := stub.MockInit("uuid", [][]byte{})
	assert.Equal(t, shim.OK, int(response.GetStatus()))
}

func TestBaseMethods(t *testing.T) {
	chaincode := &JincorBaseChaincode{}
	stub, err := tests.GetMockStubWithCidAndCC("msp1", tests.TestUser1ClientCert, chaincode)

	require.NoError(t, err)

	stub.MockTransactionStart("tx1")
	response := chaincode.Init(stub)
	stub.MockTransactionEnd("tx1")

	require.Equal(t, shim.OK, int(response.GetStatus()), response.String())

	stub.MockTransactionStart("tx2")
	stub.SetArgs([][]byte{
		[]byte("InitOwnerAccount"),
		[]byte{},
	})
	response = chaincode.Invoke(stub)
	stub.MockTransactionEnd("tx2")

	require.Equal(t, shim.OK, int(response.GetStatus()), response.String())

	acc := evm.Account{}

	require.NoError(t, json.Unmarshal(response.GetPayload(), &acc))
	require.Equal(t, tests.TestUser1ClientAddr, acc.Address)

	code := tests.LoadBytecodeFromFile("tests/bin/std/MetaCoin.bin")

	stub.MockTransactionStart("tx2")
	stub.SetArgs([][]byte{
		[]byte("DeployContract"),
		code,
		[]byte{},
	})
	response = chaincode.Invoke(stub)
	stub.MockTransactionEnd("tx2")

	require.Equal(t, shim.OK, int(response.GetStatus()), response.String())
	require.Equal(t, 32, len(response.GetPayload()))

	contractAddress := response.GetPayload()

	stub.MockTransactionStart("tx3")
	stub.SetArgs([][]byte{
		[]byte("InvokeContractMethod"),
		contractAddress,
		tests.JoinBytesArgs(tests.GetKeccakHashFromString("sendCoin(address,uint256)")[:4],
			tests.TestUser2ClientAddr.Bytes(), word256.Int64ToWord256(45).Bytes()),
	})
	response = chaincode.Invoke(stub)
	stub.MockTransactionEnd("tx3")

	require.Equal(t, shim.OK, int(response.GetStatus()), response.String())
	require.Equal(t, word256.Int64ToWord256(1).Bytes(), response.GetPayload())

	stub.MockTransactionStart("tx4")
	stub.SetArgs([][]byte{
		[]byte("InvokeContractMethod"),
		contractAddress,
		tests.JoinBytesArgs(tests.GetKeccakHashFromString("getBalance(address)")[:4],
			tests.TestUser1ClientAddr.Bytes()),
	})
	response = chaincode.Invoke(stub)
	stub.MockTransactionEnd("tx4")

	require.Equal(t, shim.OK, int(response.GetStatus()), response.String())
	require.Equal(t, word256.Int64ToWord256(10000-45).Bytes(), response.GetPayload())

	args := tests.JoinBytesArgs(tests.GetKeccakHashFromString("getBalance(address)")[:4],
		tests.TestUser2ClientAddr.Bytes())
	argsInHex := make([]byte, hex.EncodedLen(len(args)))
	hex.Encode(argsInHex, args)

	stub.MockTransactionStart("tx5")
	stub.SetArgs([][]byte{
		[]byte("InvokeContractMethodHexArgs"),
		[]byte(word256.RightPadWord256(contractAddress).Hex()),
		argsInHex,
	})
	response = chaincode.Invoke(stub)
	stub.MockTransactionEnd("tx5")

	require.Equal(t, shim.OK, int(response.GetStatus()), response.String())
	require.Equal(t, word256.Int64ToWord256(45).Bytes(), response.GetPayload())
}
