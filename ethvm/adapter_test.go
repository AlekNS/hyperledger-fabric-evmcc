package ethvm

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/require"
)

type ChainCode struct{}

func (t *ChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)

}
func (t *ChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func TestSuccessCreateNewAdapter(t *testing.T) {
	stub := shim.NewMockStub("hyperledger-fabric-evmcc", &ChainCode{})

	state := NewEvmStateAdapter(stub)

	require.NotNil(t, state)

	ev, err := NewEvmAdapter(stub, state, 0)

	require.NoError(t, err)
	require.NotNil(t, ev)
}
