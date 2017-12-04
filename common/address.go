package common

import (
	"fmt"

	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/keybase/go-triplesec/sha3"
)

// GetEtheriumLikeAddressAsBytes method to get etherium like address.
func GetEtheriumLikeAddressAsBytes(srcData []byte) ([]byte, error) {
	if srcData == nil {
		return nil, fmt.Errorf("Data can't be null for address transformation")
	}

	digest := sha3.NewKeccak256()

	if _, err := digest.Write(srcData); err != nil {
		return nil, fmt.Errorf("Keccak256 can't digest a data %s", err)
	}

	return digest.Sum(nil)[12:], nil
}

// GetEtheriumLikeAddress get etherium like address by string args
func GetEtheriumLikeAddress(srcData string) (word256.Word256, error) {
	var address []byte
	var err error
	if address, err = GetEtheriumLikeAddressAsBytes([]byte(srcData)); err != nil {
		return word256.Zero256, err
	}
	return word256.LeftPadWord256(address), nil
}

func GetOwnerEtheriumLikeAddressAsStrFromCid(clientIdent cid.ClientIdentity) (word256.Word256, error) {
	clientID, err := clientIdent.GetID()
	if err != nil {
		return word256.Zero256, err
	}
	return GetEtheriumLikeAddress(clientID)
}

// GetOwnerEtheriumLikeAddressFromStub from stub shim
func GetOwnerEtheriumLikeAddressFromStub(stub cid.ChaincodeStubInterface) (word256.Word256, error) {
	clientID, err := cid.GetID(stub)
	if err != nil {
		return word256.Zero256, err
	}
	return GetEtheriumLikeAddress(clientID)
}

// IsOwnerAddressByClientId match owner address
func IsOwnerAddressByClientId(accountAddress word256.Word256, clientIdent cid.ClientIdentity) error {
	currentOwnerAddress, err := GetOwnerEtheriumLikeAddressAsStrFromCid(clientIdent)
	if err != nil {
		return fmt.Errorf("Can't receive current owner address")
	}
	if currentOwnerAddress != accountAddress {
		return fmt.Errorf("It's not owner account")
	}

	return nil
}

// IsOwnerAddressByStub match owner address
func IsOwnerAddressByStub(accountAddress word256.Word256, stub shim.ChaincodeStubInterface) error {
	clientIdent, err := cid.New(stub)
	if err != nil {
		return err
	}

	return IsOwnerAddressByClientId(accountAddress, clientIdent)
}
