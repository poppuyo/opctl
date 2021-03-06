// This file was generated by counterfeiter
package fake

import (
	"sync"

	"github.com/opspec-io/opctl/pkg/containerengine"
	"github.com/opspec-io/opctl/util/eventbus"
	"github.com/opspec-io/sdk-golang/pkg/model"
)

type ContainerEngine struct {
	InspectContainerIfExistsStub        func(containerId string) (container *model.DcgContainerCall, err error)
	inspectContainerIfExistsMutex       sync.RWMutex
	inspectContainerIfExistsArgsForCall []struct {
		containerId string
	}
	inspectContainerIfExistsReturns struct {
		result1 *model.DcgContainerCall
		result2 error
	}
	DeleteContainerIfExistsStub        func(containerId string)
	deleteContainerIfExistsMutex       sync.RWMutex
	deleteContainerIfExistsArgsForCall []struct {
		containerId string
	}
	StartContainerStub        func(req *containerengine.StartContainerReq, eventPublisher eventbus.EventPublisher) (err error)
	startContainerMutex       sync.RWMutex
	startContainerArgsForCall []struct {
		req            *containerengine.StartContainerReq
		eventPublisher eventbus.EventPublisher
	}
	startContainerReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ContainerEngine) InspectContainerIfExists(containerId string) (container *model.DcgContainerCall, err error) {
	fake.inspectContainerIfExistsMutex.Lock()
	fake.inspectContainerIfExistsArgsForCall = append(fake.inspectContainerIfExistsArgsForCall, struct {
		containerId string
	}{containerId})
	fake.recordInvocation("InspectContainerIfExists", []interface{}{containerId})
	fake.inspectContainerIfExistsMutex.Unlock()
	if fake.InspectContainerIfExistsStub != nil {
		return fake.InspectContainerIfExistsStub(containerId)
	} else {
		return fake.inspectContainerIfExistsReturns.result1, fake.inspectContainerIfExistsReturns.result2
	}
}

func (fake *ContainerEngine) InspectContainerIfExistsCallCount() int {
	fake.inspectContainerIfExistsMutex.RLock()
	defer fake.inspectContainerIfExistsMutex.RUnlock()
	return len(fake.inspectContainerIfExistsArgsForCall)
}

func (fake *ContainerEngine) InspectContainerIfExistsArgsForCall(i int) string {
	fake.inspectContainerIfExistsMutex.RLock()
	defer fake.inspectContainerIfExistsMutex.RUnlock()
	return fake.inspectContainerIfExistsArgsForCall[i].containerId
}

func (fake *ContainerEngine) InspectContainerIfExistsReturns(result1 *model.DcgContainerCall, result2 error) {
	fake.InspectContainerIfExistsStub = nil
	fake.inspectContainerIfExistsReturns = struct {
		result1 *model.DcgContainerCall
		result2 error
	}{result1, result2}
}

func (fake *ContainerEngine) DeleteContainerIfExists(containerId string) {
	fake.deleteContainerIfExistsMutex.Lock()
	fake.deleteContainerIfExistsArgsForCall = append(fake.deleteContainerIfExistsArgsForCall, struct {
		containerId string
	}{containerId})
	fake.recordInvocation("DeleteContainerIfExists", []interface{}{containerId})
	fake.deleteContainerIfExistsMutex.Unlock()
	if fake.DeleteContainerIfExistsStub != nil {
		fake.DeleteContainerIfExistsStub(containerId)
	}
}

func (fake *ContainerEngine) DeleteContainerIfExistsCallCount() int {
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	return len(fake.deleteContainerIfExistsArgsForCall)
}

func (fake *ContainerEngine) DeleteContainerIfExistsArgsForCall(i int) string {
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	return fake.deleteContainerIfExistsArgsForCall[i].containerId
}

func (fake *ContainerEngine) StartContainer(req *containerengine.StartContainerReq, eventPublisher eventbus.EventPublisher) (err error) {
	fake.startContainerMutex.Lock()
	fake.startContainerArgsForCall = append(fake.startContainerArgsForCall, struct {
		req            *containerengine.StartContainerReq
		eventPublisher eventbus.EventPublisher
	}{req, eventPublisher})
	fake.recordInvocation("StartContainer", []interface{}{req, eventPublisher})
	fake.startContainerMutex.Unlock()
	if fake.StartContainerStub != nil {
		return fake.StartContainerStub(req, eventPublisher)
	} else {
		return fake.startContainerReturns.result1
	}
}

func (fake *ContainerEngine) StartContainerCallCount() int {
	fake.startContainerMutex.RLock()
	defer fake.startContainerMutex.RUnlock()
	return len(fake.startContainerArgsForCall)
}

func (fake *ContainerEngine) StartContainerArgsForCall(i int) (*containerengine.StartContainerReq, eventbus.EventPublisher) {
	fake.startContainerMutex.RLock()
	defer fake.startContainerMutex.RUnlock()
	return fake.startContainerArgsForCall[i].req, fake.startContainerArgsForCall[i].eventPublisher
}

func (fake *ContainerEngine) StartContainerReturns(result1 error) {
	fake.StartContainerStub = nil
	fake.startContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *ContainerEngine) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.inspectContainerIfExistsMutex.RLock()
	defer fake.inspectContainerIfExistsMutex.RUnlock()
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	fake.startContainerMutex.RLock()
	defer fake.startContainerMutex.RUnlock()
	return fake.invocations
}

func (fake *ContainerEngine) recordInvocation(key string, args []interface{}) {
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

var _ containerengine.ContainerEngine = new(ContainerEngine)
