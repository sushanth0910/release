/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"
)

type FakeBranchCheckerImpl struct {
	LSRemoteExecStub        func(string, ...string) (string, error)
	lSRemoteExecMutex       sync.RWMutex
	lSRemoteExecArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	lSRemoteExecReturns struct {
		result1 string
		result2 error
	}
	lSRemoteExecReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBranchCheckerImpl) LSRemoteExec(arg1 string, arg2 ...string) (string, error) {
	fake.lSRemoteExecMutex.Lock()
	ret, specificReturn := fake.lSRemoteExecReturnsOnCall[len(fake.lSRemoteExecArgsForCall)]
	fake.lSRemoteExecArgsForCall = append(fake.lSRemoteExecArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	stub := fake.LSRemoteExecStub
	fakeReturns := fake.lSRemoteExecReturns
	fake.recordInvocation("LSRemoteExec", []interface{}{arg1, arg2})
	fake.lSRemoteExecMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBranchCheckerImpl) LSRemoteExecCallCount() int {
	fake.lSRemoteExecMutex.RLock()
	defer fake.lSRemoteExecMutex.RUnlock()
	return len(fake.lSRemoteExecArgsForCall)
}

func (fake *FakeBranchCheckerImpl) LSRemoteExecCalls(stub func(string, ...string) (string, error)) {
	fake.lSRemoteExecMutex.Lock()
	defer fake.lSRemoteExecMutex.Unlock()
	fake.LSRemoteExecStub = stub
}

func (fake *FakeBranchCheckerImpl) LSRemoteExecArgsForCall(i int) (string, []string) {
	fake.lSRemoteExecMutex.RLock()
	defer fake.lSRemoteExecMutex.RUnlock()
	argsForCall := fake.lSRemoteExecArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBranchCheckerImpl) LSRemoteExecReturns(result1 string, result2 error) {
	fake.lSRemoteExecMutex.Lock()
	defer fake.lSRemoteExecMutex.Unlock()
	fake.LSRemoteExecStub = nil
	fake.lSRemoteExecReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBranchCheckerImpl) LSRemoteExecReturnsOnCall(i int, result1 string, result2 error) {
	fake.lSRemoteExecMutex.Lock()
	defer fake.lSRemoteExecMutex.Unlock()
	fake.LSRemoteExecStub = nil
	if fake.lSRemoteExecReturnsOnCall == nil {
		fake.lSRemoteExecReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.lSRemoteExecReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBranchCheckerImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lSRemoteExecMutex.RLock()
	defer fake.lSRemoteExecMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBranchCheckerImpl) recordInvocation(key string, args []interface{}) {
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
