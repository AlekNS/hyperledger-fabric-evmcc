package ethvm

import (
	evm "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm"
	ptypes "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/permission/types"
)

//go:generate msgp

type MsgPackEvmAccount struct {
	Address [20]byte    `msg:"address"`
	Balance int64       `msg:"balance"`
	Code    []byte      `msg:"code"`
	Nonce   int64       `msg:"nonce"`
	Other   interface{} `msg:"other"`

	Permissions MsgPackAccountPermissions `msg:"account_perms"`
}

type MsgPackAccountPermissions struct {
	Base  MsgPackBasePermissions `msg:"base"`
	Roles []string               `msg:"roles"`
}

type MsgPackBasePermissions struct {
	Perms  uint64 `msg:"perms"`
	SetBit uint64 `msg:"setbit"`
}

func translateAccountToMsgPack(account *evm.Account) MsgPackEvmAccount {
	var address [20]byte
	copy(address[:], account.Address[12:])
	return MsgPackEvmAccount{
		Address:     address,
		Balance:     account.Balance,
		Code:        account.Code,
		Nonce:       account.Nonce,
		Other:       account.Other,
		Permissions: translateAccountPermissionsToMsgPack(account.Permissions),
	}
}

func translateAccountPermissionsToMsgPack(permissions ptypes.AccountPermissions) MsgPackAccountPermissions {
	return MsgPackAccountPermissions{
		Base:  translateBasePermissionsToMsgPack(permissions.Base),
		Roles: permissions.Roles,
	}
}

func translateBasePermissionsToMsgPack(permissions ptypes.BasePermissions) MsgPackBasePermissions {
	return MsgPackBasePermissions{
		Perms:  uint64(permissions.Perms),
		SetBit: uint64(permissions.SetBit),
	}
}

func translateAccountToEvm(account MsgPackEvmAccount) *evm.Account {
	var address [32]byte
	copy(address[12:], account.Address[:])
	return &evm.Account{
		Address:     address,
		Balance:     account.Balance,
		Code:        account.Code,
		Nonce:       account.Nonce,
		Other:       account.Other,
		Permissions: translateAccountPermissionsToEvm(account.Permissions),
	}
}

func translateAccountPermissionsToEvm(permissions MsgPackAccountPermissions) ptypes.AccountPermissions {
	return ptypes.AccountPermissions{
		Base:  translateBasePermissionsToEvm(permissions.Base),
		Roles: permissions.Roles,
	}
}

func translateBasePermissionsToEvm(permissions MsgPackBasePermissions) ptypes.BasePermissions {
	return ptypes.BasePermissions{
		Perms:  ptypes.PermFlag(permissions.Perms),
		SetBit: ptypes.PermFlag(permissions.SetBit),
	}
}

func MarshalEvmAccountToMsgPack(account *evm.Account) ([]byte, error) {
	accountMsgPack := translateAccountToMsgPack(account)
	var result = []byte{}
	return accountMsgPack.MarshalMsg(result)
}

func UnmarshalEvmAccountFromMsgPack(data []byte) (*evm.Account, error) {
	account := MsgPackEvmAccount{}
	_, err := account.UnmarshalMsg(data)
	if err != nil {
		return nil, err
	}
	return translateAccountToEvm(account), nil
}
