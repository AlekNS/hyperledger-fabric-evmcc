package utils

import (
	"fmt"

	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
)

const CIPHER_KEY = "CIPHER_KEY"
const CIPHER_IV = "CIPHER_IV"

//
// CipherChainStubDecorator cipher for state values
//
type CipherChainStubDecorator struct {
	proxyChainStubDecorator
	cipher entities.EncrypterEntity
}

func NewCipherChainStubDecorator(cipher entities.EncrypterEntity, originalStub shim.ChaincodeStubInterface) (*CipherChainStubDecorator, error) {
	return &CipherChainStubDecorator{
		proxyChainStubDecorator: proxyChainStubDecorator{
			originalStub: originalStub,
		},
		cipher: cipher,
	}, nil
}

func NewCipherChainStubDecoratorFromTransientMap(keyPrivateKey, keyIv string, originalStub shim.ChaincodeStubInterface) (*CipherChainStubDecorator, error) {
	var key, iv []byte
	var exists bool
	var cipher entities.EncrypterEntity

	transientMap, err := originalStub.GetTransient()
	if err != nil {
		return nil, err
	}
	if key, exists = transientMap[keyPrivateKey]; !exists {
		return nil, fmt.Errorf("No %s Key in transient map", keyPrivateKey)
	}
	if iv, exists = transientMap[keyIv]; !exists {
		return nil, fmt.Errorf("No %s IV in transient map", keyIv)
	}
	if cipher, err = entities.NewAES256EncrypterEntity("ID", factory.GetDefault(), key, iv); err != nil {
		return nil, err
	}

	return NewCipherChainStubDecorator(cipher, originalStub)
}

func NewCipherChainStubDecoratorFromDefaultTransientMap(originalStub shim.ChaincodeStubInterface) (*CipherChainStubDecorator, error) {
	return NewCipherChainStubDecoratorFromTransientMap(CIPHER_KEY, CIPHER_IV, originalStub)
}

type cipherStateQueryIterator struct {
	shim.StateQueryIteratorInterface
	cipher           entities.EncrypterEntity
	originalIterator shim.StateQueryIteratorInterface
}

func (t *cipherStateQueryIterator) HasNext() bool {
	return t.originalIterator.HasNext()
}

func (t *cipherStateQueryIterator) Close() error {
	return t.originalIterator.Close()
}

func (t *cipherStateQueryIterator) Next() (*queryresult.KV, error) {
	var err error
	kv, err := t.originalIterator.Next()
	if err != nil {
		return kv, err
	}
	kv.Value, err = t.cipher.Decrypt(kv.GetValue())
	if err != nil {
		return kv, err
	}
	return kv, nil
}

type cipherHistoryQueryIterator struct {
	shim.HistoryQueryIteratorInterface
	originalIterator shim.HistoryQueryIteratorInterface
	cipher           entities.EncrypterEntity
}

func (t *cipherHistoryQueryIterator) HasNext() bool {
	return t.originalIterator.HasNext()
}

func (t *cipherHistoryQueryIterator) Close() error {
	return t.originalIterator.Close()
}

func (t *cipherHistoryQueryIterator) Next() (*queryresult.KeyModification, error) {
	var err error
	kvm, err := t.originalIterator.Next()
	if err != nil {
		return kvm, err
	}
	kvm.Value, err = t.cipher.Decrypt(kvm.GetValue())
	if err != nil {
		return kvm, err
	}
	return kvm, nil
}

func (t *CipherChainStubDecorator) GetState(key string) ([]byte, error) {
	encryptedValue, err := t.proxyChainStubDecorator.GetState(key)
	if err != nil {
		return nil, err
	}
	value, err := t.cipher.Decrypt(encryptedValue)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (t *CipherChainStubDecorator) PutState(key string, value []byte) error {
	encryptedValue, err := t.cipher.Encrypt(value)
	if err != nil {
		return err
	}
	return t.proxyChainStubDecorator.PutState(key, encryptedValue)
}

func (t *CipherChainStubDecorator) GetStateByRange(startKey, endKey string) (shim.StateQueryIteratorInterface, error) {
	iter, err := t.proxyChainStubDecorator.GetStateByRange(startKey, endKey)
	if err != nil {
		return iter, err
	}
	return &cipherStateQueryIterator{
		cipher:           t.cipher,
		originalIterator: iter,
	}, nil
}

func (t *CipherChainStubDecorator) GetStateByPartialCompositeKey(objectType string, keys []string) (shim.StateQueryIteratorInterface, error) {
	iter, err := t.proxyChainStubDecorator.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return iter, err
	}
	return &cipherStateQueryIterator{
		cipher:           t.cipher,
		originalIterator: iter,
	}, nil
}

func (t *CipherChainStubDecorator) GetHistoryForKey(key string) (shim.HistoryQueryIteratorInterface, error) {
	iter, err := t.proxyChainStubDecorator.GetHistoryForKey(key)
	if err != nil {
		return iter, err
	}
	return &cipherHistoryQueryIterator{
		cipher:           t.cipher,
		originalIterator: iter,
	}, nil
}
