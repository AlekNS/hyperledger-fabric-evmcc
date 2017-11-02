package ethvm

import (
	"encoding/json"
	"encoding/hex"
	"strings"
	"fmt"

	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
	"github.com/JincorTech/hyperledger-fabric-evmcc/utils"

	evm "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm"
	ptypes "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/permission/types"
	"github.com/JincorTech/hyperledger-fabric-evmcc/common"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type EvmDispatcher struct {
	stub  shim.ChaincodeStubInterface
	state *EvmStateAdapter
}

func NewEvmDispatcher(stub shim.ChaincodeStubInterface) (*EvmDispatcher, error) {
	decoratedStub := utils.NewWriteCachedStubDecorator(stub)
	return &EvmDispatcher{
		stub:  decoratedStub,
		state: NewEvmStateAdapter(decoratedStub),
	}, nil
}

func (t *EvmDispatcher) getEvm() (*EvmAdapter, error) {
	ethvm, err := NewEvmAdapter(t.stub, t.state, 0)
	if err != nil {
		return nil, err
	}

	return ethvm, nil
}

func (t *EvmDispatcher) getOwnerAccount() (*evm.Account, error) {
	ownerAddress, err := common.GetOwnerEtheriumLikeAddressFromStub(t.stub)
	if err != nil {
		return nil, err
	}
	state := NewEvmStateAdapter(t.stub)

	ownerAccount, err := state.TryGetAccount(ownerAddress)
	if ownerAccount == nil {
		t.state.UpdateAccount(&evm.Account{
			Address:     ownerAddress,
			Permissions: ptypes.DefaultAccountPermissions,
		})
		ownerAccount = state.GetAccount(ownerAddress)
	}

	return ownerAccount, nil
}

func (t *EvmDispatcher) Initiate() pb.Response {
	return shim.Success(nil)
}

func (t *EvmDispatcher) execute(method func(*evm.Account, *EvmAdapter) (response pb.Response)) (response pb.Response) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				response = shim.Error(v.Error())
			case string:
				response = shim.Error(v)
			default:
				response = shim.Error(fmt.Sprintf("%v", v))
			}
		}
	}()

	ownerAccount, err := t.getOwnerAccount()
	if err != nil {
		panic(err)
	}
	vm, err := t.getEvm()
	if err != nil {
		panic(err)
	}

	response = method(ownerAccount, vm)
	return response
}

func (t *EvmDispatcher) Invoke() (response pb.Response) {
	args := t.stub.GetArgs()
	if len(args) < 1 {
		return shim.Error("Empty arguments")
	}

	fn := string(args[0])
	args = args[1:]

	hexArgsIndex := strings.Index(fn, "HexArgs")

	if hexArgsIndex > 0 && hexArgsIndex == len(fn) - 7 {
		for inx, arg := range args {
			dst := make([]byte, hex.DecodedLen(len(arg)))
			n, err := hex.Decode(dst, arg)
			if err != nil {
				return shim.Error(fmt.Sprintf("Invalid data in hex format was receieved for method %s, %d / %v", fn, inx, err))
			}
			args[inx] = dst[:n]
		}
		fn = fn[:len(fn) - 7]
	}

	switch fn {
	case "DeployContract":
		if len(args) != 2 {
			return shim.Error("Too less arguments to deploy contract")
		}
		if len(args[0]) == 0 {
			return shim.Error("First argument is code, but it's empty")
		}
		return t.execute(func(owner *evm.Account, vm *EvmAdapter) pb.Response {
			cotractAccount, err := vm.InstallContractByOwner(args[0], args[1])

			if err != nil {
				return shim.Error(fmt.Sprintf("%v", err))
			}

			return shim.Success(cotractAccount.Address.Bytes())
		})
	case "InvokeContractMethod":
		if len(args) != 2 {
			return shim.Error("Too less arguments to invoke method of a contract")
		}
		if len(args[0]) < 20 || len(args[0]) > 32 {
			return shim.Error("Invalid first argument, invalid contract address")
		}
		return t.execute(func(owner *evm.Account, vm *EvmAdapter) pb.Response {
			contractAccount := t.state.GetAccount(word256.LeftPadWord256(args[0]))

			if len(contractAccount.Code) == 0 {
				return shim.Error(fmt.Sprintf("Account hasn't code at addr: %X", args[0]))
			}

			result, err := vm.CallMethodByOwner(contractAccount, 0, args[1])
			if err != nil {
				return shim.Error(fmt.Sprintf("Error was occurred whan invoke method, account with addr: %X / %v", args[0], err))
			}

			return shim.Success(result)
		})
	case "InitOwnerAccount":
		return t.execute(func(owner *evm.Account, vm *EvmAdapter) pb.Response {
			ownerAccount, err := t.getOwnerAccount()

			if err != nil {
				return shim.Error(fmt.Sprintf("Init owner account failed: %v", err))
			}

			data, err := json.Marshal(ownerAccount)
			if err != nil {
				return shim.Error(fmt.Sprintf("Init owner account marshaling failed: %v", err))
			}

			return shim.Success(data)
		})
	default:
		return shim.Error(fmt.Sprintf("Unknown method to invoke %s", string(fn)))
	}
}
