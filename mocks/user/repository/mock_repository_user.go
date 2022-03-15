// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_userRepository is a generated GoMock package.
package mock_userRepository

import (
	reflect "reflect"
	userEntities "sirclo/project-capstone/entities/userEntities"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepoInterface is a mock of UserRepoInterface interface.
type MockUserRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoInterfaceMockRecorder
}

// MockUserRepoInterfaceMockRecorder is the mock recorder for MockUserRepoInterface.
type MockUserRepoInterfaceMockRecorder struct {
	mock *MockUserRepoInterface
}

// NewMockUserRepoInterface creates a new mock instance.
func NewMockUserRepoInterface(ctrl *gomock.Controller) *MockUserRepoInterface {
	mock := &MockUserRepoInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepoInterface) EXPECT() *MockUserRepoInterfaceMockRecorder {
	return m.recorder
}

// CheckEmail mocks base method.
func (m *MockUserRepoInterface) CheckEmail(userChecked userEntities.User) (userEntities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmail", userChecked)
	ret0, _ := ret[0].(userEntities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmail indicates an expected call of CheckEmail.
func (mr *MockUserRepoInterfaceMockRecorder) CheckEmail(userChecked interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmail", reflect.TypeOf((*MockUserRepoInterface)(nil).CheckEmail), userChecked)
}

// CreateUser mocks base method.
func (m *MockUserRepoInterface) CreateUser(user userEntities.User) (userEntities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(userEntities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoInterfaceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepoInterface)(nil).CreateUser), user)
}

// GetUser mocks base method.
func (m *MockUserRepoInterface) GetUser(id string) (userEntities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(userEntities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepoInterfaceMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepoInterface)(nil).GetUser), id)
}

// Login mocks base method.
func (m *MockUserRepoInterface) Login(identity string) (userEntities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", identity)
	ret0, _ := ret[0].(userEntities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserRepoInterfaceMockRecorder) Login(identity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserRepoInterface)(nil).Login), identity)
}

// UpdateUser mocks base method.
func (m *MockUserRepoInterface) UpdateUser(user userEntities.User) (userEntities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", user)
	ret0, _ := ret[0].(userEntities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepoInterfaceMockRecorder) UpdateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepoInterface)(nil).UpdateUser), user)
}

// UploadAvatarUser mocks base method.
func (m *MockUserRepoInterface) UploadAvatarUser(user userEntities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatarUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAvatarUser indicates an expected call of UploadAvatarUser.
func (mr *MockUserRepoInterfaceMockRecorder) UploadAvatarUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatarUser", reflect.TypeOf((*MockUserRepoInterface)(nil).UploadAvatarUser), user)
}