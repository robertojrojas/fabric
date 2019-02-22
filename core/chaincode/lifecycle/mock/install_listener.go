// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/chaincode/lifecycle"
	"github.com/hyperledger/fabric/core/chaincode/persistence"
)

type InstallListener struct {
	HandleChaincodeInstalledStub        func(*persistence.ChaincodePackageMetadata, []byte)
	handleChaincodeInstalledMutex       sync.RWMutex
	handleChaincodeInstalledArgsForCall []struct {
		arg1 *persistence.ChaincodePackageMetadata
		arg2 []byte
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *InstallListener) HandleChaincodeInstalled(arg1 *persistence.ChaincodePackageMetadata, arg2 []byte) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.handleChaincodeInstalledMutex.Lock()
	fake.handleChaincodeInstalledArgsForCall = append(fake.handleChaincodeInstalledArgsForCall, struct {
		arg1 *persistence.ChaincodePackageMetadata
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("HandleChaincodeInstalled", []interface{}{arg1, arg2Copy})
	fake.handleChaincodeInstalledMutex.Unlock()
	if fake.HandleChaincodeInstalledStub != nil {
		fake.HandleChaincodeInstalledStub(arg1, arg2)
	}
}

func (fake *InstallListener) HandleChaincodeInstalledCallCount() int {
	fake.handleChaincodeInstalledMutex.RLock()
	defer fake.handleChaincodeInstalledMutex.RUnlock()
	return len(fake.handleChaincodeInstalledArgsForCall)
}

func (fake *InstallListener) HandleChaincodeInstalledCalls(stub func(*persistence.ChaincodePackageMetadata, []byte)) {
	fake.handleChaincodeInstalledMutex.Lock()
	defer fake.handleChaincodeInstalledMutex.Unlock()
	fake.HandleChaincodeInstalledStub = stub
}

func (fake *InstallListener) HandleChaincodeInstalledArgsForCall(i int) (*persistence.ChaincodePackageMetadata, []byte) {
	fake.handleChaincodeInstalledMutex.RLock()
	defer fake.handleChaincodeInstalledMutex.RUnlock()
	argsForCall := fake.handleChaincodeInstalledArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *InstallListener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleChaincodeInstalledMutex.RLock()
	defer fake.handleChaincodeInstalledMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *InstallListener) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ lifecycle.InstallListener = new(InstallListener)
