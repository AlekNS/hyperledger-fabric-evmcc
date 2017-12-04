package main

import (
	"github.com/JincorTech/hyperledger-fabric-evmcc/ethvm"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("hyperledger-fabric-evmcc/jincorbase")

type JincorBaseChaincode struct {
	shim.Chaincode
}

func (t *JincorBaseChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Initiate")

	ethvmDispatcher, err := ethvm.NewEvmDispatcher(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	return ethvmDispatcher.Initiate()
}

func (t *JincorBaseChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Invoke")

	ethvmDispatcher, err := ethvm.NewEvmDispatcher(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	return ethvmDispatcher.Invoke()
}

func main() {
	defer func() {
		if reason := recover(); reason != nil {
			logger.Errorf("Error was occurred when running of the hyperledger-fabric-evmcc: %v", reason)
		}
	}()

	shim.SetupChaincodeLogging()

	logger.Info("Start")

	factory.InitFactories(nil)

	err := shim.Start(&JincorBaseChaincode{})
	if err != nil {
		panic(err)
	}
}
