package ethvm

import (
	"encoding/json"
	"fmt"

	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/txs"
	"github.com/JincorTech/hyperledger-fabric-evmcc/common"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	events "github.com/tendermint/go-events"

	evm "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm"
	// . "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm/opcodes"
	// "github.com/golang/protobuf/proto"
	// cb "github.com/hyperledger/fabric/protos/common"
)

type EvmAdapter struct {
	stub          shim.ChaincodeStubInterface
	state         *EvmStateAdapter
	eventSwitch   events.EventSwitch
	gasLimitValue int
}

func NewEvmAdapter(stub shim.ChaincodeStubInterface, evmState *EvmStateAdapter, gasLimit int) (*EvmAdapter, error) {
	if gasLimit == 0 {
		gasLimit = 1000000
	}
	return &EvmAdapter{
		stub:          stub,
		state:         evmState,
		gasLimitValue: gasLimit,
	}, nil
}

func (t *EvmAdapter) newEvm(ownerAccount *evm.Account) (*evm.VM, events.EventSwitch, error) {
	// evm.SetDebug(true) // @TODO: Remove it
	if ownerAccount == nil {
		return nil, nil, fmt.Errorf("Invalid owner account")
	}
	eventSwitch := events.NewEventSwitch()
	_, err := eventSwitch.Start()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to start eventSwitch: %v", err)
	}
	txTime, err := t.stub.GetTxTimestamp()
	if err != nil {
		return nil, nil, fmt.Errorf("Can't get transaction timestamp %s", err)
	}

	// qccResponse := stub.InvokeChaincode("qcc", [][]byte{
	// 	[]byte("GetChainInfo"),
	// 	[]byte("cid"),
	// }, "")
	// blkInfo := cb.BlockchainInfo{}
	// proto.Unmarshal(qccResponse.GetPayload(), &blkInfo)

	vm := evm.NewVM(t.state, evm.DefaultDynamicMemoryProvider, evm.Params{
		BlockTime: txTime.GetSeconds(),
		// BlockHash: blkInfo.GetCurrentBlockHash()[:32], <- should be word256
		// BlockHeight: blkInfo.GetHeight(),
		GasLimit: int64(t.gasLimitValue),
	}, ownerAccount.Address, nil)
	vm.SetFireable(eventSwitch)

	return vm, eventSwitch, nil
}

func (t *EvmAdapter) installContract(account *evm.Account, transferValue int64, code, inputArgs []byte) (*evm.Account, error) {
	contractAccount := t.state.CreateAccount(account)

	code = append(code, inputArgs...) // << pass as args for "constructor"

	contractAccount.Code = code
	result, err := t.callMethod(account, contractAccount, transferValue, inputArgs)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("Empty code result")
	}
	contractAccount.Code = result
	t.state.UpdateAccount(contractAccount)
	t.state.UpdateAccount(account)
	return contractAccount, nil
}

func (t *EvmAdapter) InstallContractByOwner(code, inputArgs []byte) (*evm.Account, error) {
	if len(code) == 0 {
		return nil, fmt.Errorf("Empty code")
	}
	ownerAddress, err := common.GetOwnerEtheriumLikeAddressFromStub(t.stub)
	if err != nil {
		return nil, err
	}
	ownerAccount := t.state.GetAccount(ownerAddress)
	if ownerAccount == nil {
		return nil, fmt.Errorf("Invalid owner account %X", ownerAddress)
	}
	return t.installContract(ownerAccount, 0, code, inputArgs)
}

func (t *EvmAdapter) callMethod(caller, contract *evm.Account, transferValue int64, inputArgs []byte) ([]byte, error) {
	vm, eventSwitch, err := t.newEvm(caller)
	if err != nil {
		return nil, err
	}

	eventLogID := txs.EventStringLogEvent(contract.Address.Postfix(20))
	eventSwitch.AddListenerForEvent("evm", eventLogID, func(event events.EventData) {
		addr := event.(txs.EventDataLog).Address.Hex()[24:]
		payLoad, err := json.Marshal(event.(txs.EventDataLog))
		if err != nil {
			t.stub.SetEvent("EVM:LOG:"+addr+":ERROR", []byte(err.Error()))
		} else {
			t.stub.SetEvent("EVM:LOG:"+addr+":OK", payLoad)
		}
	})

	defer func() {
		eventSwitch.RemoveListener(eventLogID)
		eventSwitch.Stop()
	}()

	endGas := int64(t.gasLimitValue)
	result, err := vm.Call(caller, contract, contract.Code, inputArgs, transferValue, &endGas)
	return result, err
}

func (t *EvmAdapter) CallMethodByOwner(contract *evm.Account, transferValue int64, inputArgs []byte) ([]byte, error) {
	ownerAddress, err := common.GetOwnerEtheriumLikeAddressFromStub(t.stub)
	if err != nil {
		return nil, err
	}
	ownerAccount := t.state.GetAccount(ownerAddress)
	if ownerAccount == nil {
		return nil, fmt.Errorf("Invalid owner account %X", ownerAddress)
	}
	result, err := t.callMethod(ownerAccount, contract, transferValue, inputArgs)
	if err != nil {
		return nil, err
	}
	return result, nil
}
