package utils

import (
	"testing"

	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	"github.com/stretchr/testify/assert"
)

const (
	ENCKEY = "01234567890123456789012345678901" // 256 bit
	IV     = "0123456789012345"                 // 128 bit
)

func TestCipherPutState(t *testing.T) {
	factory.InitFactories(nil)

	entity, err := entities.NewAES256EncrypterEntity("ID", factory.GetDefault(), []byte(ENCKEY), []byte(IV))
	assert.NoError(t, err)

	mock := shim.NewMockStub("hyperledger-fabric-evmcc", &chainCode{})
	stub, err := NewCipherChainStubDecorator(entity, mock)

	assert.NoError(t, err)

	mock.MockTransactionStart("tx1")
	stub.PutState("key1", []byte("01234567890123456789012345678901"))
	mock.MockTransactionEnd("tx1")

	// @TODO: Find solution for remove all not needed data
	assert.Equal(t, []byte(IV), mock.State["key1"][:16])
	assert.Equal(t, []byte{
		0xf7, 0xcb, 0x93, 0xfe, 0x68, 0x01, 0x61, 0x74, 0x7c, 0x01, 0x94, 0xe6, 0xa3, 0x5b, 0xab, 0xea,
		0xc0, 0x94, 0x87, 0x54, 0x7d, 0xde, 0xcc, 0xf7, 0x8c, 0x6a, 0x3f, 0x6b, 0xeb, 0x15, 0x11, 0x32,
	}, mock.State["key1"][16:48])
}
