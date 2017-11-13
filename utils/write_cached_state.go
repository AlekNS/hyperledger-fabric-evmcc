package utils

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

//
// Write cached state
//
type WriteCachedStubDecorator struct {
	proxyChainStubDecorator
	cache map[string][]byte
}

func NewWriteCachedStubDecorator(originalStub shim.ChaincodeStubInterface) *WriteCachedStubDecorator {
	return &WriteCachedStubDecorator{
		proxyChainStubDecorator: proxyChainStubDecorator{
			originalStub: originalStub,
		},
		cache: make(map[string][]byte),
	}
}

func (t *WriteCachedStubDecorator) GetState(key string) ([]byte, error) {
	var err error
	value, ok := t.cache[key]
	if !ok {
		value, err = t.proxyChainStubDecorator.GetState(key)
		if err != nil {
			return nil, err
		}
	}
	return value, nil
}

func (t *WriteCachedStubDecorator) PutState(key string, value []byte) error {
	err := t.proxyChainStubDecorator.PutState(key, value)
	if err != nil {
		return err
	}
	t.cache[key] = value
	return nil
}

func (t *WriteCachedStubDecorator) DelState(key string) error {
    delete(t.cache, key)
    return t.proxyChainStubDecorator.DelState(key)
}
