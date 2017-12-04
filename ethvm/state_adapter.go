package ethvm

import (
	"encoding/base64"
	"fmt"

	evm "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm"
	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm/sha3"
	ptypes "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/permission/types"
	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
	"github.com/JincorTech/hyperledger-fabric-evmcc/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type EvmStateAdapter struct {
	evm.AppState
	stub shim.ChaincodeStubInterface
}

func NewEvmStateAdapter(stub shim.ChaincodeStubInterface) *EvmStateAdapter {
	return &EvmStateAdapter{
		stub: utils.NewNamespaceChainStubDecorator("evm", stub),
	}
}

func (t *EvmStateAdapter) SetGlobalPermissions(perms ptypes.AccountPermissions) {
	t.UpdateAccount(&evm.Account{
		Address:     word256.Zero256,
		Balance:     0,
		Code:        nil,
		Permissions: perms,
	})
}

func (t *EvmStateAdapter) CreateAccount(creator *evm.Account) *evm.Account {
	addr := createNewEvmAddress(creator)
	value, err := t.stub.GetState(base64.StdEncoding.EncodeToString(addr.Bytes()[12:]))
	if err == nil && len(value) == 0 {
		return &evm.Account{
			Address: addr,
			Balance: 0,
			Code:    nil,
			Nonce:   creator.Nonce,
		}
	}
	panic(fmt.Sprintf("Create account with addr: %X, error %v", addr, err))
}

func (t *EvmStateAdapter) TryGetAccount(addr word256.Word256) (*evm.Account, error) {
	value, err := t.stub.GetState(base64.StdEncoding.EncodeToString(addr.Bytes()[12:]))
	if err != nil || len(value) == 0 {
		return nil, fmt.Errorf("Get account by addr: %X, with error %v", addr, err)
	}

	account, err := UnmarshalEvmAccountFromMsgPack(value)
	if err != nil {
		return nil, fmt.Errorf("Get account by addr: %X, unmarshal error %v", addr, err)
	}
	return account, nil
}

func (t *EvmStateAdapter) GetAccount(addr word256.Word256) *evm.Account {
	account, err := t.TryGetAccount(addr)
	if err != nil {
		panic(err.Error())
	}
	return account
}

func (t *EvmStateAdapter) UpdateAccount(account *evm.Account) {
	data, err := MarshalEvmAccountToMsgPack(account)
	if err != nil {
		panic(fmt.Sprintf("Update account with addr: %X, marshal error %v", account.Address, err))
	}
	err = t.stub.PutState(base64.StdEncoding.EncodeToString(account.Address.Bytes()[12:]), data)
	if err != nil {
		panic(fmt.Sprintf("Update account with addr: %X, save state error %v", account.Address, err))
	}
}

func (t *EvmStateAdapter) RemoveAccount(account *evm.Account) {
	value, err := t.stub.GetState(base64.StdEncoding.EncodeToString(account.Address.Bytes()[12:]))
	if err != nil || len(value) == 0 {
		panic(fmt.Sprintf("Remove account with addr: %X, error %v", account.Address, err))
	}
	err = t.stub.DelState(base64.StdEncoding.EncodeToString(account.Address.Bytes()[12:]))
	if err != nil {
		panic(fmt.Sprintf("Remove account with addr: %X, delete state error %v", account.Address, err))
	}
}

func (t *EvmStateAdapter) GetStorage(addr word256.Word256, key word256.Word256) word256.Word256 {
	// account, err := t.stub.GetState(base64.StdEncoding.EncodeToString(addr.Bytes()[12:]))
	// if err != nil || len(account) == 0 {
	// 	panic(fmt.Sprintf("Get storage for account addr: %X, key: %X, error %v", addr, key, err))
	// }
	value, err := t.stub.GetState(base64.StdEncoding.EncodeToString(addr.Bytes()[12:]) + ":s:" +
		base64.StdEncoding.EncodeToString(key.Bytes()))
	if err == nil {
		return word256.LeftPadWord256(value)
	}
	return word256.Zero256
}

func (t *EvmStateAdapter) SetStorage(addr word256.Word256, key word256.Word256, value word256.Word256) {
	// account, err := t.stub.GetState(base64.StdEncoding.EncodeToString(addr.Bytes()[12:]))
	// if err != nil || len(account) == 0 {
	// 	panic(fmt.Sprintf("Set storage for account addr: %X, key: %X, error %v", addr, key, err))
	// }
	err := t.stub.PutState(base64.StdEncoding.EncodeToString(addr.Bytes()[12:])+":s:"+
		base64.StdEncoding.EncodeToString(key.Bytes()), value.Bytes())
	if err != nil {
		panic(fmt.Sprintf("Set storage for account addr: %X, key: %X, save state error %v", addr, key, err))
	}
}

func createNewEvmAddress(creator *evm.Account) word256.Word256 {
	nonce := creator.Nonce
	creator.Nonce++
	temp := make([]byte, 32+8)
	copy(temp, creator.Address[:])
	word256.PutInt64BE(temp[32:], nonce)
	return word256.LeftPadWord256(sha3.Sha3(temp)[:20])
}
