package utils

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
)

//
// NamespaceChainStubDecorator namespaced state
//
type NamespaceChainStubDecorator struct {
	proxyChainStubDecorator
	keyBasis string
}

func NewNamespaceChainStubDecorator(keyBasis string, originalStub shim.ChaincodeStubInterface) *NamespaceChainStubDecorator {
	stub := &NamespaceChainStubDecorator{
		keyBasis: keyBasis + ":",
	}
	stub.originalStub = originalStub
	return stub
}

type namespaceStateQueryIterator struct {
	shim.StateQueryIteratorInterface
	keyBasis         string
	originalIterator shim.StateQueryIteratorInterface
}

func (t *namespaceStateQueryIterator) HasNext() bool {
	return t.originalIterator.HasNext()
}

func (t *namespaceStateQueryIterator) Close() error {
	return t.originalIterator.Close()
}

func (t *namespaceStateQueryIterator) Next() (*queryresult.KV, error) {
	var err error
	kv, err := t.originalIterator.Next()
	if err != nil {
		return kv, err
	}
	kv.Key = unwrapKey(kv.Key, t.keyBasis)
	return kv, nil
}

func (t *NamespaceChainStubDecorator) wrapKey(key string) string {
	return t.keyBasis + key
}

func unwrapKey(key, keyBasis string) string {
	return string([]rune(key)[len(keyBasis):])
}

func (t *NamespaceChainStubDecorator) isWrapped(key string) bool {
	return string([]rune(key)[:len(t.keyBasis)]) == t.keyBasis
}

func (t *NamespaceChainStubDecorator) GetState(key string) ([]byte, error) {
	return t.proxyChainStubDecorator.GetState(t.wrapKey(key))
}

func (t *NamespaceChainStubDecorator) PutState(key string, value []byte) error {
	return t.proxyChainStubDecorator.PutState(t.wrapKey(key), value)
}

func (t *NamespaceChainStubDecorator) DelState(key string) error {
	return t.proxyChainStubDecorator.DelState(t.wrapKey(key))
}

func (t *NamespaceChainStubDecorator) GetStateByRange(startKey, endKey string) (shim.StateQueryIteratorInterface, error) {
	iter, err := t.proxyChainStubDecorator.GetStateByRange(t.wrapKey(startKey), t.wrapKey(endKey))
	if err != nil {
		return nil, err
	}
	return &namespaceStateQueryIterator{
		keyBasis:         t.keyBasis,
		originalIterator: iter,
	}, nil
}

func (t *NamespaceChainStubDecorator) GetStateByPartialCompositeKey(objectType string, keys []string) (shim.StateQueryIteratorInterface, error) {
	iter, err := t.proxyChainStubDecorator.GetStateByPartialCompositeKey(t.wrapKey(objectType), keys)
	if err != nil {
		return nil, err
	}
	return &namespaceStateQueryIterator{
		keyBasis:         t.keyBasis,
		originalIterator: iter,
	}, nil
}

func (t *NamespaceChainStubDecorator) CreateCompositeKey(objectType string, attributes []string) (string, error) {
	return t.proxyChainStubDecorator.CreateCompositeKey(t.wrapKey(objectType), attributes)
}

func (t *NamespaceChainStubDecorator) SplitCompositeKey(compositeKey string) (string, []string, error) {
	objType, attrs, err := t.proxyChainStubDecorator.SplitCompositeKey(compositeKey)

	if err != nil {
		return objType, attrs, err
	}

	if !t.isWrapped(objType) {
		return objType, attrs, nil
	}

	return unwrapKey(objType, t.keyBasis), attrs, nil
}

func (t *NamespaceChainStubDecorator) GetHistoryForKey(key string) (shim.HistoryQueryIteratorInterface, error) {
	return t.proxyChainStubDecorator.GetHistoryForKey(t.wrapKey(key))
}
