// This file was generated by counterfeiter
package validate

import (
	"sync"

	"github.com/opspec-io/sdk-golang/pkg/model"
)

type Fake struct {
	ParamStub        func(arg *model.Data, param *model.Param) (errors []error)
	paramMutex       sync.RWMutex
	paramArgsForCall []struct {
		arg   *model.Data
		param *model.Param
	}
	paramReturns struct {
		result1 []error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Param(arg *model.Data, param *model.Param) (errors []error) {
	fake.paramMutex.Lock()
	fake.paramArgsForCall = append(fake.paramArgsForCall, struct {
		arg   *model.Data
		param *model.Param
	}{arg, param})
	fake.recordInvocation("Param", []interface{}{arg, param})
	fake.paramMutex.Unlock()
	if fake.ParamStub != nil {
		return fake.ParamStub(arg, param)
	} else {
		return fake.paramReturns.result1
	}
}

func (fake *Fake) ParamCallCount() int {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	return len(fake.paramArgsForCall)
}

func (fake *Fake) ParamArgsForCall(i int) (*model.Data, *model.Param) {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	return fake.paramArgsForCall[i].arg, fake.paramArgsForCall[i].param
}

func (fake *Fake) ParamReturns(result1 []error) {
	fake.ParamStub = nil
	fake.paramReturns = struct {
		result1 []error
	}{result1}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	return fake.invocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ Validate = new(Fake)
