package utils

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//
// proxyChainStubDecorator proxy all calls
//
type proxyChainStubDecorator struct {
	shim.ChaincodeStubInterface
	originalStub shim.ChaincodeStubInterface
}

func (t *proxyChainStubDecorator) GetArgs() [][]byte {
	return t.originalStub.GetArgs()
}

func (t *proxyChainStubDecorator) GetStringArgs() []string {
	return t.originalStub.GetStringArgs()
}

func (t *proxyChainStubDecorator) GetFunctionAndParameters() (string, []string) {
	return t.originalStub.GetFunctionAndParameters()
}

func (t *proxyChainStubDecorator) GetArgsSlice() ([]byte, error) {
	return t.originalStub.GetArgsSlice()
}

func (t *proxyChainStubDecorator) GetTxID() string {
	return t.originalStub.GetTxID()
}

func (t *proxyChainStubDecorator) InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response {
	return t.originalStub.InvokeChaincode(chaincodeName, args, channel)
}

func (t *proxyChainStubDecorator) GetState(key string) ([]byte, error) {
	return t.originalStub.GetState(key)
}

func (t *proxyChainStubDecorator) PutState(key string, value []byte) error {
	return t.originalStub.PutState(key, value)
}

func (t *proxyChainStubDecorator) DelState(key string) error {
	return t.originalStub.DelState(key)
}

func (t *proxyChainStubDecorator) GetStateByRange(startKey, endKey string) (shim.StateQueryIteratorInterface, error) {
	return t.originalStub.GetStateByRange(startKey, endKey)
}

func (t *proxyChainStubDecorator) GetStateByPartialCompositeKey(objectType string, keys []string) (shim.StateQueryIteratorInterface, error) {
	return t.originalStub.GetStateByPartialCompositeKey(objectType, keys)
}

func (t *proxyChainStubDecorator) CreateCompositeKey(objectType string, attributes []string) (string, error) {
	return t.originalStub.CreateCompositeKey(objectType, attributes)
}

func (t *proxyChainStubDecorator) SplitCompositeKey(compositeKey string) (string, []string, error) {
	return t.originalStub.SplitCompositeKey(compositeKey)
}

func (t *proxyChainStubDecorator) GetQueryResult(query string) (shim.StateQueryIteratorInterface, error) {
	return t.originalStub.GetQueryResult(query)
}

func (t *proxyChainStubDecorator) GetHistoryForKey(key string) (shim.HistoryQueryIteratorInterface, error) {
	return t.originalStub.GetHistoryForKey(key)
}

func (t *proxyChainStubDecorator) GetCreator() ([]byte, error) {
	return t.originalStub.GetCreator()
}

func (t *proxyChainStubDecorator) GetTransient() (map[string][]byte, error) {
	return t.originalStub.GetTransient()
}

func (t *proxyChainStubDecorator) GetBinding() ([]byte, error) {
	return t.originalStub.GetBinding()
}

func (t *proxyChainStubDecorator) GetDecorations() map[string][]byte {
	return t.originalStub.GetDecorations()
}

func (t *proxyChainStubDecorator) GetSignedProposal() (*pb.SignedProposal, error) {
	return t.originalStub.GetSignedProposal()
}

func (t *proxyChainStubDecorator) GetTxTimestamp() (*timestamp.Timestamp, error) {
	return t.originalStub.GetTxTimestamp()
}

func (t *proxyChainStubDecorator) SetEvent(name string, payload []byte) error {
	return t.originalStub.SetEvent(name, payload)
}
