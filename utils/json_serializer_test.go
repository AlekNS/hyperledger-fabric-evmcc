package utils

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
)

type sampleObject struct {
	IntVal int    `json:"int_val"`
	StrVal string `json:"str_val"`
}

var sample = sampleObject{
	IntVal: 123,
	StrVal: "456",
}

const sampleAsJSON = "{\"int_val\":123,\"str_val\":\"456\"}"

func TestGetJsonStateSerializer(t *testing.T) {
	var joDst = sampleObject{}

	mock := shim.NewMockStub("hyperledger-fabric-evmcc", &chainCode{})
	mock.State["key1"] = []byte(sampleAsJSON)

	assert.NoError(t, GetJsonFromState(mock, "key1", &joDst))
	assert.Equal(t, 123, joDst.IntVal)
	assert.Equal(t, "456", joDst.StrVal)
}

func TestPutJsonStateSerializer(t *testing.T) {
	mock := shim.NewMockStub("hyperledger-fabric-evmcc", &chainCode{})

	mock.MockTransactionStart("tx1")
	assert.NoError(t, PutJsonToState(mock, "key1", &sample))
	assert.Equal(t, []byte(sampleAsJSON), mock.State["key1"])
	mock.MockTransactionEnd("tx1")
}

func TestGetCCValidArgAsJson(t *testing.T) {
	var joDst = sampleObject{}
	err := GetCCArgAsJson(&joDst, 0, [][]byte{
		[]byte(sampleAsJSON),
	})
	assert.NoError(t, err)
	assert.Equal(t, sample, joDst)
}

func TestGetCCInvalidArgAsJson(t *testing.T) {
	var joDst = sampleObject{}
	err := GetCCArgAsJson(&joDst, 0, [][]byte{
		[]byte("{["),
	})
	assert.EqualError(t, err, "invalid character '[' looking for beginning of object key string")
}

func TestGetCCInvalidArgIndexAsJson(t *testing.T) {
	var joDst = sampleObject{}
	err := GetCCArgAsJson(&joDst, 10, [][]byte{
		[]byte(sampleAsJSON),
	})
	assert.EqualError(t, err, "ArgIndex in GetArgAsJSON violate size of args 10/1")
}

func TestSuccessResponseAsJson(t *testing.T) {
	response := CcSuccessResponseAsJson(&sample)
	assert.Equal(t, shim.OK, int(response.GetStatus()))
	assert.Equal(t, response.Payload, response.GetPayload())
}

func TestSuccessResponseAsJsonWithInvalidArgument(t *testing.T) {
	response := CcSuccessResponseAsJson(nil)
	assert.Equal(t, shim.ERROR, int(response.GetStatus()))
	assert.Equal(t, "Can't response with json passed as nil value", response.GetMessage())
}
