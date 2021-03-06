// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-travel-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	travel "github.com/ozonva/ova-travel-api/internal/travel"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddEntities mocks base method.
func (m *MockRepo) AddEntities(arg0 []travel.Trip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEntities", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEntities indicates an expected call of AddEntities.
func (mr *MockRepoMockRecorder) AddEntities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntities", reflect.TypeOf((*MockRepo)(nil).AddEntities), arg0)
}

// DescribeEntity mocks base method.
func (m *MockRepo) DescribeEntity(arg0 uint64) (*travel.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEntity", arg0)
	ret0, _ := ret[0].(*travel.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntity indicates an expected call of DescribeEntity.
func (mr *MockRepoMockRecorder) DescribeEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntity", reflect.TypeOf((*MockRepo)(nil).DescribeEntity), arg0)
}

// ListEntities mocks base method.
func (m *MockRepo) ListEntities(arg0, arg1 uint64) ([]travel.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntities", arg0, arg1)
	ret0, _ := ret[0].([]travel.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntities indicates an expected call of ListEntities.
func (mr *MockRepoMockRecorder) ListEntities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntities", reflect.TypeOf((*MockRepo)(nil).ListEntities), arg0, arg1)
}
