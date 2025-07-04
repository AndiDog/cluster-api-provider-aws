/*
Copyright The Kubernetes Authors.

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

// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services/instancestate (interfaces: EventBridgeAPI)

// Package mock_eventbridgeiface is a generated GoMock package.
package mock_eventbridgeiface

import (
	context "context"
	reflect "reflect"

	eventbridge "github.com/aws/aws-sdk-go-v2/service/eventbridge"
	gomock "github.com/golang/mock/gomock"
)

// MockEventBridgeAPI is a mock of EventBridgeAPI interface.
type MockEventBridgeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockEventBridgeAPIMockRecorder
}

// MockEventBridgeAPIMockRecorder is the mock recorder for MockEventBridgeAPI.
type MockEventBridgeAPIMockRecorder struct {
	mock *MockEventBridgeAPI
}

// NewMockEventBridgeAPI creates a new mock instance.
func NewMockEventBridgeAPI(ctrl *gomock.Controller) *MockEventBridgeAPI {
	mock := &MockEventBridgeAPI{ctrl: ctrl}
	mock.recorder = &MockEventBridgeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventBridgeAPI) EXPECT() *MockEventBridgeAPIMockRecorder {
	return m.recorder
}

// DeleteRule mocks base method.
func (m *MockEventBridgeAPI) DeleteRule(arg0 context.Context, arg1 *eventbridge.DeleteRuleInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DeleteRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRule", varargs...)
	ret0, _ := ret[0].(*eventbridge.DeleteRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRule indicates an expected call of DeleteRule.
func (mr *MockEventBridgeAPIMockRecorder) DeleteRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRule", reflect.TypeOf((*MockEventBridgeAPI)(nil).DeleteRule), varargs...)
}

// DescribeRule mocks base method.
func (m *MockEventBridgeAPI) DescribeRule(arg0 context.Context, arg1 *eventbridge.DescribeRuleInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRule", varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRule indicates an expected call of DescribeRule.
func (mr *MockEventBridgeAPIMockRecorder) DescribeRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRule", reflect.TypeOf((*MockEventBridgeAPI)(nil).DescribeRule), varargs...)
}

// ListTargetsByRule mocks base method.
func (m *MockEventBridgeAPI) ListTargetsByRule(arg0 context.Context, arg1 *eventbridge.ListTargetsByRuleInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListTargetsByRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTargetsByRule", varargs...)
	ret0, _ := ret[0].(*eventbridge.ListTargetsByRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTargetsByRule indicates an expected call of ListTargetsByRule.
func (mr *MockEventBridgeAPIMockRecorder) ListTargetsByRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTargetsByRule", reflect.TypeOf((*MockEventBridgeAPI)(nil).ListTargetsByRule), varargs...)
}

// PutRule mocks base method.
func (m *MockEventBridgeAPI) PutRule(arg0 context.Context, arg1 *eventbridge.PutRuleInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.PutRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutRule", varargs...)
	ret0, _ := ret[0].(*eventbridge.PutRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRule indicates an expected call of PutRule.
func (mr *MockEventBridgeAPIMockRecorder) PutRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRule", reflect.TypeOf((*MockEventBridgeAPI)(nil).PutRule), varargs...)
}

// PutTargets mocks base method.
func (m *MockEventBridgeAPI) PutTargets(arg0 context.Context, arg1 *eventbridge.PutTargetsInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.PutTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutTargets", varargs...)
	ret0, _ := ret[0].(*eventbridge.PutTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutTargets indicates an expected call of PutTargets.
func (mr *MockEventBridgeAPIMockRecorder) PutTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutTargets", reflect.TypeOf((*MockEventBridgeAPI)(nil).PutTargets), varargs...)
}

// RemoveTargets mocks base method.
func (m *MockEventBridgeAPI) RemoveTargets(arg0 context.Context, arg1 *eventbridge.RemoveTargetsInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.RemoveTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveTargets", varargs...)
	ret0, _ := ret[0].(*eventbridge.RemoveTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTargets indicates an expected call of RemoveTargets.
func (mr *MockEventBridgeAPIMockRecorder) RemoveTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTargets", reflect.TypeOf((*MockEventBridgeAPI)(nil).RemoveTargets), varargs...)
}
