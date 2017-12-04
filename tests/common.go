package tests

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"

	evm "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm"
	ptypes "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/permission/types"
	"github.com/JincorTech/hyperledger-fabric-evmcc/ethvm"
	"github.com/keybase/go-triplesec/sha3"
	// . "github.com/JincorTech/hyperledger-fabric-evmcc/burrow/evm/opcodes"
	pb "github.com/hyperledger/fabric/protos/peer"

	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"

	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
)

type ChainCode struct{}

func (t *ChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)

}
func (t *ChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

type MockStubWithCreator struct {
	shim.MockStub
	cid.ChaincodeStubInterface
	Events  map[string][][]byte
	creator []byte
	args    [][]byte
}

func (s *MockStubWithCreator) GetArgs() [][]byte {
	return s.args
}

func (s *MockStubWithCreator) SetArgs(args [][]byte) {
	s.args = args
}

func (s *MockStubWithCreator) GetCreator() ([]byte, error) {
	return s.creator, nil
}

func (s *MockStubWithCreator) SetEvent(name string, payload []byte) error {
	if _, ok := s.Events[name]; !ok {
		s.Events[name] = [][]byte{}
	}
	s.Events[name] = append(s.Events[name], payload)
	return nil
}

func (s *MockStubWithCreator) SetCreator(mspId, cert string) error {
	creator, err := GetCreatorFromSerializedIdentitiy(mspId, cert)
	if err != nil {
		return err
	}
	s.creator = creator
	return nil
}

var TestUser1ClientAddr = word256.LeftPadWord256([]byte{
	0x0f, 0x5a, 0x50, 0xa0, 0x87, 0xab, 0xd0, 0x82, 0x08, 0x40,
	0xaf, 0x41, 0x8c, 0xbb, 0x01, 0xf2, 0x29, 0xb5, 0x28, 0x7e,
})

const TestUser1ClientCert = `
-----BEGIN CERTIFICATE-----
MIICBzCCAa6gAwIBAgIUR0wk/DLjm2PCGskw7CRue0uhLaQwCgYIKoZIzj0EAwIw
dzELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNh
biBGcmFuY2lzY28xGzAZBgNVBAoTEm5ldHdvcmsuamluY29yLmNvbTEeMBwGA1UE
AxMVY2EubmV0d29yay5qaW5jb3IuY29tMB4XDTE3MTEwMjEwNDEwMFoXDTE4MTEw
MjEwNDEwMFowIzEhMB8GA1UEAwwYVXNlcjFAbmV0d29yay5qaW5jb3IuY29tMFkw
EwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAECbIXQcbZ5U7ru0XeIKDNcJcgRPmY3VrM
bGOXW+Yk0s8oIovbMWtgEZ/pdOZbynIGm8GT7OkQHLL7sxDY1mSqgqNsMGowDgYD
VR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFPgn4/MewRsikCTt
U1cLNQHnz7TqMCsGA1UdIwQkMCKAIG1tY5pkJ0JpckTk7JtUm5pxZkKnS0IQdhIh
0PGVvRD7MAoGCCqGSM49BAMCA0cAMEQCIGc/2GTVOxaBgqQGVw3JZslyh10Ul4eo
poqbrgGwyzyeAiAEy9uFn48L2pe7YIOG8Byg98VVyJZQ0l9pzJOdApT/Rw==
-----END CERTIFICATE-----
`

var TestUser2ClientAddr = word256.LeftPadWord256([]byte{
	0x57, 0xe8, 0xc5, 0xda, 0xca, 0xab, 0x4b, 0x8b, 0x40, 0xe5,
	0xaf, 0xc7, 0xdc, 0x78, 0x8d, 0x4d, 0x78, 0x6a, 0x10, 0x91,
})

const TestUser2ClientCert = `
-----BEGIN CERTIFICATE-----
MIICBzCCAa6gAwIBAgIUNXXAEMhVfsfSebWJBEyYCTx+4x0wCgYIKoZIzj0EAwIw
dzELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNh
biBGcmFuY2lzY28xGzAZBgNVBAoTEm5ldHdvcmsuamluY29yLmNvbTEeMBwGA1UE
AxMVY2EubmV0d29yay5qaW5jb3IuY29tMB4XDTE3MTEwMjEwNDEwMFoXDTE4MTEw
MjEwNDEwMFowIzEhMB8GA1UEAwwYVXNlcjJAbmV0d29yay5qaW5jb3IuY29tMFkw
EwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYhIsUrOBkhBNRibBD7o07bmgoTZ+qaxj
Hboz04b4Rww1O7N/2mGRr2j3JZDjw1foA4K7bQfALbk0eCb5yqSEiaNsMGowDgYD
VR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFB/AfbVOBpfG+Ngl
jW4VZPjG3NybMCsGA1UdIwQkMCKAIG1tY5pkJ0JpckTk7JtUm5pxZkKnS0IQdhIh
0PGVvRD7MAoGCCqGSM49BAMCA0cAMEQCIFHgHYmTC1kooEJmG4HcY4O3VG5NhmRm
KI9YcEoATCW3AiA7CSAHB32EOKD/vbRR6jkOkiKyIg6s59wwGkso02Nnqg==
-----END CERTIFICATE-----
`

var TestUser3ClientAddr = word256.LeftPadWord256([]byte{
	0xa0, 0x2e, 0x67, 0x12, 0x43, 0xd4, 0x2c, 0xd1, 0xb0, 0x63,
	0xd7, 0x1a, 0x68, 0xca, 0xf9, 0xec, 0x41, 0x95, 0xec, 0x82,
})

const TestUser3ClientCert = `
-----BEGIN CERTIFICATE-----
MIICCDCCAa6gAwIBAgIUAIJ/V6kOKgs9UHDo4nx0CInckUkwCgYIKoZIzj0EAwIw
dzELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNh
biBGcmFuY2lzY28xGzAZBgNVBAoTEm5ldHdvcmsuamluY29yLmNvbTEeMBwGA1UE
AxMVY2EubmV0d29yay5qaW5jb3IuY29tMB4XDTE3MTEwMjEwNDEwMFoXDTE4MTEw
MjEwNDEwMFowIzEhMB8GA1UEAwwYVXNlcjNAbmV0d29yay5qaW5jb3IuY29tMFkw
EwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESty/SjrYAlr59+iPvbbtG8NcEQbOR23c
sHn3it2PpHviZB0KM4ekzIKA0OSaWIgCnugxlS8SsfQTjYPF0YuIdKNsMGowDgYD
VR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFELRU3pbHrUkkVni
aobzEf1mz5TzMCsGA1UdIwQkMCKAIG1tY5pkJ0JpckTk7JtUm5pxZkKnS0IQdhIh
0PGVvRD7MAoGCCqGSM49BAMCA0gAMEUCIQDsYF21604g92PGCk4Z3MXU49EDNzaY
vIJCXch0HRZtEAIgA7QOXp0V11LKb4PKP5wraS5QncGMcUGW6S1/VmDjHHw=
-----END CERTIFICATE-----
`

func HexStringToBinary(strCode string) []byte {
	code, _ := hex.DecodeString(strings.Replace(strCode, "\n", "", -1))
	return code
}

func GetCreatorFromSerializedIdentitiy(mspId string, cert string) ([]byte, error) {
	sid := &msp.SerializedIdentity{Mspid: mspId,
		IdBytes: []byte(cert)}
	b, err := proto.Marshal(sid)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetMockStubWithCid(mspId string, cert string) (*MockStubWithCreator, error) {
	stub := &MockStubWithCreator{
		MockStub: *shim.NewMockStub("hyperledger-fabric-evmcc", &ChainCode{}),
		Events:   make(map[string][][]byte),
	}

	return stub, stub.SetCreator(mspId, cert)
}

func GetMockStubWithCidAndCC(mspId string, cert string, cc shim.Chaincode) (*MockStubWithCreator, error) {
	stub := &MockStubWithCreator{
		MockStub: *shim.NewMockStub("hyperledger-fabric-evmcc", cc),
		Events:   make(map[string][][]byte),
		args:     [][]byte{},
	}
	return stub, stub.SetCreator(mspId, cert)
}

func GetEvmAdapterAndMockStub(mspId,
	userCert string,
	userAddr word256.Word256) (stub *MockStubWithCreator, state *ethvm.EvmStateAdapter, ev *ethvm.EvmAdapter) {

	stub, _ = GetMockStubWithCid(mspId, TestUser1ClientCert)

	state = ethvm.NewEvmStateAdapter(stub)
	stub.MockTransactionStart("tx1")

	state.SetGlobalPermissions(ptypes.DefaultAccountPermissions)

	owner := &evm.Account{
		Address: TestUser1ClientAddr,
		Balance: 100,
	}

	state.UpdateAccount(owner)
	state.UpdateAccount(&evm.Account{
		Address: TestUser2ClientAddr,
		Balance: 200,
	})
	state.UpdateAccount(&evm.Account{
		Address: TestUser3ClientAddr,
		Balance: 300,
	})
	stub.MockTransactionEnd("tx1")

	ev, _ = ethvm.NewEvmAdapter(stub, state, 10000)
	return
}

func GetKeccakHashFromString(dataStr string) []byte {
	digest := sha3.NewKeccak256()
	if _, err := digest.Write([]byte(dataStr)); err != nil {
		panic(fmt.Sprintf("Keccak256 can't digest a data %s", err))
	}
	return digest.Sum(nil)
}

func JoinBytesArgs(args ...[]byte) []byte {
	inputArgs := []byte{}
	for _, arg := range args {
		inputArgs = append(inputArgs, arg...)
	}
	return inputArgs
}

func CallEvmMethod(stub *MockStubWithCreator,
	ev *ethvm.EvmAdapter, contractAcc *evm.Account,
	transactionValue int64, methodSignature string, args ...[]byte) ([]byte, error) {
	inputArgs := GetKeccakHashFromString(methodSignature)[:4]
	inputArgs = append(inputArgs, JoinBytesArgs(args...)...)
	stub.MockTransactionStart("tx")
	result, err := ev.CallMethodByOwner(contractAcc, transactionValue, inputArgs)
	stub.MockTransactionEnd("tx")

	return result, err
}

func LoadBytecodeFromFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("can't load file %s", path))
	}
	return HexStringToBinary(string(data))
}
