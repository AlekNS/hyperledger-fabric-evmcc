package utils

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func PutJsonToState(stub shim.ChaincodeStubInterface, key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return stub.PutState(key, bytes)
}

func GetJsonFromState(stub shim.ChaincodeStubInterface, key string, value interface{}) error {
	bytes, err := stub.GetState(key)
	if err != nil {
		return err
	}
	if len(bytes) == 0 {
		return fmt.Errorf("Empty data")
	}

	return json.Unmarshal(bytes, value)
}

type JsonStateQueryIterator struct {
	shim.StateQueryIteratorInterface
	originalIterator shim.StateQueryIteratorInterface
}

func (t *JsonStateQueryIterator) HasNext() bool {
	return t.originalIterator.HasNext()
}

func (t *JsonStateQueryIterator) Close() error {
	return t.originalIterator.Close()
}

func (t *JsonStateQueryIterator) Next(value interface{}) error {
	kv, err := t.originalIterator.Next()
	if err != nil {
		return err
	}
	return json.Unmarshal(kv.GetValue(), value)
}

type JsonHistoryQueryIterator struct {
	shim.HistoryQueryIteratorInterface
	originalIterator shim.HistoryQueryIteratorInterface
}

func (t *JsonHistoryQueryIterator) HasNext() bool {
	return t.originalIterator.HasNext()
}

func (t *JsonHistoryQueryIterator) Close() error {
	return t.originalIterator.Close()
}

func (t *JsonHistoryQueryIterator) Next(value interface{}) error {
	kvm, err := t.originalIterator.Next()
	if err != nil {
		return err
	}
	return json.Unmarshal(kvm.GetValue(), value)
}

func GetJsonStateByRange(stub shim.ChaincodeStubInterface, startKey, endKey string) (*JsonStateQueryIterator, error) {
	iter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	return &JsonStateQueryIterator{
		originalIterator: iter,
	}, nil
}

func GetStateByPartialCompositeKey(stub shim.ChaincodeStubInterface, objectType string, keys []string) (*JsonStateQueryIterator, error) {
	iter, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, err
	}
	return &JsonStateQueryIterator{
		originalIterator: iter,
	}, nil
}

func GetHistoryForKey(stub shim.ChaincodeStubInterface, key string) (*JsonHistoryQueryIterator, error) {
	iter, err := stub.GetHistoryForKey(key)
	if err != nil {
		return nil, err
	}
	return &JsonHistoryQueryIterator{
		originalIterator: iter,
	}, nil
}

func GetCCArgAsJson(value interface{}, argIndex int, args [][]byte) error {
	if argIndex < 0 || argIndex >= len(args) {
		return fmt.Errorf("ArgIndex in GetArgAsJSON violate size of args %d/%d", argIndex, len(args))
	}
	return json.Unmarshal(args[argIndex], value)
}

func ccResponseAsJson(isSuccess bool, value interface{}) pb.Response {
	if value == nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprint("Can't response with json passed as nil value"),
		}
	}
	payload, err := json.Marshal(value)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("Error was occurred when marshal value to json %s", err),
		}
	}
	if isSuccess {
		return pb.Response{Status: shim.OK, Payload: payload}
	}
	return pb.Response{Status: shim.ERROR, Payload: payload}
}

func CcSuccessResponseAsJson(value interface{}) pb.Response {
	return ccResponseAsJson(true, value)
}

func CcErrorResponseAsJson(value interface{}) pb.Response {
	return ccResponseAsJson(false, value)
}
