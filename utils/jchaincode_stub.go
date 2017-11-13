package utils

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//
// JcChaincodeStubInterface proxy all calls
//
type JcChaincodeStubInterface struct {
	shim.ChaincodeStubInterface
	originalStub shim.ChaincodeStubInterface
}

func (t *JcChaincodeStubInterface) GetArgs() [][]byte {
	return t.originalStub.GetArgs()
}

func (t *JcChaincodeStubInterface) GetStringArgs() []string {
	return t.originalStub.GetStringArgs()
}

func (t *JcChaincodeStubInterface) GetFunctionAndParameters() (string, []string) {
	return t.originalStub.GetFunctionAndParameters()
}

func (t *JcChaincodeStubInterface) GetArgsSlice() ([]byte, error) {
	return t.originalStub.GetArgsSlice()
}

func (t *JcChaincodeStubInterface) GetTxID() string {
	return t.originalStub.GetTxID()
}

func (t *JcChaincodeStubInterface) InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response {
	return t.originalStub.InvokeChaincode(chaincodeName, args, channel)
}

func (t *JcChaincodeStubInterface) GetState(key string) ([]byte, error) {
	return t.originalStub.GetState(key)
}

func (t *JcChaincodeStubInterface) PutState(key string, value []byte) error {
	return t.originalStub.PutState(key, value)
}

func (t *JcChaincodeStubInterface) DelState(key string) error {
	return t.originalStub.DelState(key)
}

func (t *JcChaincodeStubInterface) GetStateByRange(startKey, endKey string) (shim.StateQueryIteratorInterface, error) {
	return t.originalStub.GetStateByRange(startKey, endKey)
}

func (t *JcChaincodeStubInterface) GetStateByPartialCompositeKey(objectType string, keys []string) (shim.StateQueryIteratorInterface, error) {
	return t.originalStub.GetStateByPartialCompositeKey(objectType, keys)
}

func (t *JcChaincodeStubInterface) CreateCompositeKey(objectType string, attributes []string) (string, error) {
	return t.originalStub.CreateCompositeKey(objectType, attributes)
}

func (t *JcChaincodeStubInterface) SplitCompositeKey(compositeKey string) (string, []string, error) {
	return t.originalStub.SplitCompositeKey(compositeKey)
}

func (t *JcChaincodeStubInterface) GetQueryResult(query string) (shim.StateQueryIteratorInterface, error) {
	return t.originalStub.GetQueryResult(query)
}

func (t *JcChaincodeStubInterface) GetHistoryForKey(key string) (shim.HistoryQueryIteratorInterface, error) {
	return t.originalStub.GetHistoryForKey(key)
}

func (t *JcChaincodeStubInterface) GetCreator() ([]byte, error) {
	return t.originalStub.GetCreator()
}

func (t *JcChaincodeStubInterface) GetTransient() (map[string][]byte, error) {
	return t.originalStub.GetTransient()
}

func (t *JcChaincodeStubInterface) GetBinding() ([]byte, error) {
	return t.originalStub.GetBinding()
}

func (t *JcChaincodeStubInterface) GetDecorations() map[string][]byte {
	return t.originalStub.GetDecorations()
}

func (t *JcChaincodeStubInterface) GetSignedProposal() (*pb.SignedProposal, error) {
	return t.originalStub.GetSignedProposal()
}

func (t *JcChaincodeStubInterface) GetTxTimestamp() (*timestamp.Timestamp, error) {
	ts, err := t.originalStub.GetTxTimestamp()
	return &timestamp.Timestamp{
		Seconds: ts.Seconds,
		Nanos:   ts.Nanos,
	}, err
}

func (t *JcChaincodeStubInterface) SetEvent(name string, payload []byte) error {
	return t.originalStub.SetEvent(name, payload)
}
