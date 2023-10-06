// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/weaveworks/eksctl/pkg/utils/kubectl"
)

type FakeKubernetesVersionManager struct {
	ClientVersionStub        func() (string, error)
	clientVersionMutex       sync.RWMutex
	clientVersionArgsForCall []struct {
	}
	clientVersionReturns struct {
		result1 string
		result2 error
	}
	clientVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ServerVersionStub        func([]string, []string) (string, error)
	serverVersionMutex       sync.RWMutex
	serverVersionArgsForCall []struct {
		arg1 []string
		arg2 []string
	}
	serverVersionReturns struct {
		result1 string
		result2 error
	}
	serverVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ValidateVersionStub        func(string, kubectl.VersionType) error
	validateVersionMutex       sync.RWMutex
	validateVersionArgsForCall []struct {
		arg1 string
		arg2 kubectl.VersionType
	}
	validateVersionReturns struct {
		result1 error
	}
	validateVersionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKubernetesVersionManager) ClientVersion() (string, error) {
	fake.clientVersionMutex.Lock()
	ret, specificReturn := fake.clientVersionReturnsOnCall[len(fake.clientVersionArgsForCall)]
	fake.clientVersionArgsForCall = append(fake.clientVersionArgsForCall, struct {
	}{})
	stub := fake.ClientVersionStub
	fakeReturns := fake.clientVersionReturns
	fake.recordInvocation("ClientVersion", []interface{}{})
	fake.clientVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKubernetesVersionManager) ClientVersionCallCount() int {
	fake.clientVersionMutex.RLock()
	defer fake.clientVersionMutex.RUnlock()
	return len(fake.clientVersionArgsForCall)
}

func (fake *FakeKubernetesVersionManager) ClientVersionCalls(stub func() (string, error)) {
	fake.clientVersionMutex.Lock()
	defer fake.clientVersionMutex.Unlock()
	fake.ClientVersionStub = stub
}

func (fake *FakeKubernetesVersionManager) ClientVersionReturns(result1 string, result2 error) {
	fake.clientVersionMutex.Lock()
	defer fake.clientVersionMutex.Unlock()
	fake.ClientVersionStub = nil
	fake.clientVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKubernetesVersionManager) ClientVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.clientVersionMutex.Lock()
	defer fake.clientVersionMutex.Unlock()
	fake.ClientVersionStub = nil
	if fake.clientVersionReturnsOnCall == nil {
		fake.clientVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.clientVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKubernetesVersionManager) ServerVersion(arg1 []string, arg2 []string) (string, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.serverVersionMutex.Lock()
	ret, specificReturn := fake.serverVersionReturnsOnCall[len(fake.serverVersionArgsForCall)]
	fake.serverVersionArgsForCall = append(fake.serverVersionArgsForCall, struct {
		arg1 []string
		arg2 []string
	}{arg1Copy, arg2Copy})
	stub := fake.ServerVersionStub
	fakeReturns := fake.serverVersionReturns
	fake.recordInvocation("ServerVersion", []interface{}{arg1Copy, arg2Copy})
	fake.serverVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKubernetesVersionManager) ServerVersionCallCount() int {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	return len(fake.serverVersionArgsForCall)
}

func (fake *FakeKubernetesVersionManager) ServerVersionCalls(stub func([]string, []string) (string, error)) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = stub
}

func (fake *FakeKubernetesVersionManager) ServerVersionArgsForCall(i int) ([]string, []string) {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	argsForCall := fake.serverVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKubernetesVersionManager) ServerVersionReturns(result1 string, result2 error) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	fake.serverVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKubernetesVersionManager) ServerVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	if fake.serverVersionReturnsOnCall == nil {
		fake.serverVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.serverVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKubernetesVersionManager) ValidateVersion(arg1 string, arg2 kubectl.VersionType) error {
	fake.validateVersionMutex.Lock()
	ret, specificReturn := fake.validateVersionReturnsOnCall[len(fake.validateVersionArgsForCall)]
	fake.validateVersionArgsForCall = append(fake.validateVersionArgsForCall, struct {
		arg1 string
		arg2 kubectl.VersionType
	}{arg1, arg2})
	stub := fake.ValidateVersionStub
	fakeReturns := fake.validateVersionReturns
	fake.recordInvocation("ValidateVersion", []interface{}{arg1, arg2})
	fake.validateVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKubernetesVersionManager) ValidateVersionCallCount() int {
	fake.validateVersionMutex.RLock()
	defer fake.validateVersionMutex.RUnlock()
	return len(fake.validateVersionArgsForCall)
}

func (fake *FakeKubernetesVersionManager) ValidateVersionCalls(stub func(string, kubectl.VersionType) error) {
	fake.validateVersionMutex.Lock()
	defer fake.validateVersionMutex.Unlock()
	fake.ValidateVersionStub = stub
}

func (fake *FakeKubernetesVersionManager) ValidateVersionArgsForCall(i int) (string, kubectl.VersionType) {
	fake.validateVersionMutex.RLock()
	defer fake.validateVersionMutex.RUnlock()
	argsForCall := fake.validateVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKubernetesVersionManager) ValidateVersionReturns(result1 error) {
	fake.validateVersionMutex.Lock()
	defer fake.validateVersionMutex.Unlock()
	fake.ValidateVersionStub = nil
	fake.validateVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKubernetesVersionManager) ValidateVersionReturnsOnCall(i int, result1 error) {
	fake.validateVersionMutex.Lock()
	defer fake.validateVersionMutex.Unlock()
	fake.ValidateVersionStub = nil
	if fake.validateVersionReturnsOnCall == nil {
		fake.validateVersionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateVersionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKubernetesVersionManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clientVersionMutex.RLock()
	defer fake.clientVersionMutex.RUnlock()
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	fake.validateVersionMutex.RLock()
	defer fake.validateVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKubernetesVersionManager) recordInvocation(key string, args []interface{}) {
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

var _ kubectl.KubernetesVersionManager = new(FakeKubernetesVersionManager)
