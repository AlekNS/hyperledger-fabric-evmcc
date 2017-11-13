package utils

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
)

func TestWriteCachedStatePutState(t *testing.T) {
	var value = []byte("value")

	mock := shim.NewMockStub("hyperledger-fabric-evmcc", &chainCode{})
	stub := NewWriteCachedStubDecorator(mock)

	mock.MockTransactionStart("tx1")
	stub.PutState("key1", value)
	mock.MockTransactionEnd("tx1")

	valFromState, exists := mock.State["key1"]
	assert.True(t, exists)
	assert.Equal(t, value, valFromState)

	valFromState, exists = stub.cache["key1"]
	assert.True(t, exists)
	assert.Equal(t, value, valFromState)
}

func TestWriteCachedStateGetState(t *testing.T) {
	var value1 = []byte("value1")
	var value2 = []byte("value2")

	mock := shim.NewMockStub("hyperledger-fabric-evmcc", &chainCode{})
	stub := NewWriteCachedStubDecorator(mock)

	stub.cache["key1"] = value1
	mock.State["key2"] = value2

	v, err := stub.GetState("key1")
	assert.NoError(t, err)
	assert.Equal(t, value1, v)

	v, err = stub.GetState("key2")
	assert.NoError(t, err)
	assert.Equal(t, value2, v)
}
