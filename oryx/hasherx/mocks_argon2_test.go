// Copyright © 2025 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/x/hasherx (interfaces: Argon2Configurator)
//
// Generated by this command:
//
//	mockgen -package hasherx_test -destination hasherx/mocks_argon2_test.go github.com/ory/x/hasherx Argon2Configurator
//

// Package hasherx_test is a generated GoMock package.
package hasherx_test

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	hasherx "github.com/ory/x/hasherx"
)

// MockArgon2Configurator is a mock of Argon2Configurator interface.
type MockArgon2Configurator struct {
	ctrl     *gomock.Controller
	recorder *MockArgon2ConfiguratorMockRecorder
	isgomock struct{}
}

// MockArgon2ConfiguratorMockRecorder is the mock recorder for MockArgon2Configurator.
type MockArgon2ConfiguratorMockRecorder struct {
	mock *MockArgon2Configurator
}

// NewMockArgon2Configurator creates a new mock instance.
func NewMockArgon2Configurator(ctrl *gomock.Controller) *MockArgon2Configurator {
	mock := &MockArgon2Configurator{ctrl: ctrl}
	mock.recorder = &MockArgon2ConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArgon2Configurator) EXPECT() *MockArgon2ConfiguratorMockRecorder {
	return m.recorder
}

// HasherArgon2Config mocks base method.
func (m *MockArgon2Configurator) HasherArgon2Config(ctx context.Context) *hasherx.Argon2Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasherArgon2Config", ctx)
	ret0, _ := ret[0].(*hasherx.Argon2Config)
	return ret0
}

// HasherArgon2Config indicates an expected call of HasherArgon2Config.
func (mr *MockArgon2ConfiguratorMockRecorder) HasherArgon2Config(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasherArgon2Config", reflect.TypeOf((*MockArgon2Configurator)(nil).HasherArgon2Config), ctx)
}
