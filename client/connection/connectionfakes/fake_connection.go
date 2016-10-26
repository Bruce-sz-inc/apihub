// This file was generated by counterfeiter
package connectionfakes

import (
	"sync"

	"github.com/apihub/apihub"
	"github.com/apihub/apihub/client/connection"
)

type FakeConnection struct {
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 error
	}
	AddServiceStub        func(apihub.ServiceSpec) (apihub.ServiceSpec, error)
	addServiceMutex       sync.RWMutex
	addServiceArgsForCall []struct {
		arg1 apihub.ServiceSpec
	}
	addServiceReturns struct {
		result1 apihub.ServiceSpec
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnection) Ping() error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	} else {
		return fake.pingReturns.result1
	}
}

func (fake *FakeConnection) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeConnection) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) AddService(arg1 apihub.ServiceSpec) (apihub.ServiceSpec, error) {
	fake.addServiceMutex.Lock()
	fake.addServiceArgsForCall = append(fake.addServiceArgsForCall, struct {
		arg1 apihub.ServiceSpec
	}{arg1})
	fake.recordInvocation("AddService", []interface{}{arg1})
	fake.addServiceMutex.Unlock()
	if fake.AddServiceStub != nil {
		return fake.AddServiceStub(arg1)
	} else {
		return fake.addServiceReturns.result1, fake.addServiceReturns.result2
	}
}

func (fake *FakeConnection) AddServiceCallCount() int {
	fake.addServiceMutex.RLock()
	defer fake.addServiceMutex.RUnlock()
	return len(fake.addServiceArgsForCall)
}

func (fake *FakeConnection) AddServiceArgsForCall(i int) apihub.ServiceSpec {
	fake.addServiceMutex.RLock()
	defer fake.addServiceMutex.RUnlock()
	return fake.addServiceArgsForCall[i].arg1
}

func (fake *FakeConnection) AddServiceReturns(result1 apihub.ServiceSpec, result2 error) {
	fake.AddServiceStub = nil
	fake.addServiceReturns = struct {
		result1 apihub.ServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.addServiceMutex.RLock()
	defer fake.addServiceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConnection) recordInvocation(key string, args []interface{}) {
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

var _ connection.Connection = new(FakeConnection)