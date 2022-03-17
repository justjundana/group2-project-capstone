// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_attendanceRepository is a generated GoMock package.
package mock_attendanceRepository

import (
	reflect "reflect"
	attendanceEntities "sirclo/project-capstone/entities/attendanceEntities"

	gomock "github.com/golang/mock/gomock"
)

// MockAttendanceRepoInterface is a mock of AttendanceRepoInterface interface.
type MockAttendanceRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAttendanceRepoInterfaceMockRecorder
}

// MockAttendanceRepoInterfaceMockRecorder is the mock recorder for MockAttendanceRepoInterface.
type MockAttendanceRepoInterfaceMockRecorder struct {
	mock *MockAttendanceRepoInterface
}

// NewMockAttendanceRepoInterface creates a new mock instance.
func NewMockAttendanceRepoInterface(ctrl *gomock.Controller) *MockAttendanceRepoInterface {
	mock := &MockAttendanceRepoInterface{ctrl: ctrl}
	mock.recorder = &MockAttendanceRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAttendanceRepoInterface) EXPECT() *MockAttendanceRepoInterfaceMockRecorder {
	return m.recorder
}

// CreateAttendance mocks base method.
func (m *MockAttendanceRepoInterface) CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttendance", att)
	ret0, _ := ret[0].(attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttendance indicates an expected call of CreateAttendance.
func (mr *MockAttendanceRepoInterfaceMockRecorder) CreateAttendance(att interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttendance", reflect.TypeOf((*MockAttendanceRepoInterface)(nil).CreateAttendance), att)
}

// GetAttendancesById mocks base method.
func (m *MockAttendanceRepoInterface) GetAttendancesById(attID string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendancesById", attID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAttendancesById indicates an expected call of GetAttendancesById.
func (mr *MockAttendanceRepoInterfaceMockRecorder) GetAttendancesById(attID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendancesById", reflect.TypeOf((*MockAttendanceRepoInterface)(nil).GetAttendancesById), attID)
}

// GetAttendancesCurrentUser mocks base method.
func (m *MockAttendanceRepoInterface) GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendancesCurrentUser", userId, status, order)
	ret0, _ := ret[0].([]attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttendancesCurrentUser indicates an expected call of GetAttendancesCurrentUser.
func (mr *MockAttendanceRepoInterfaceMockRecorder) GetAttendancesCurrentUser(userId, status, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendancesCurrentUser", reflect.TypeOf((*MockAttendanceRepoInterface)(nil).GetAttendancesCurrentUser), userId, status, order)
}

// GetAttendancesRangeDate mocks base method.
func (m *MockAttendanceRepoInterface) GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order string) ([]attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendancesRangeDate", employeeEmail, dateStart, dateEnd, status, officeId, order)
	ret0, _ := ret[0].([]attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttendancesRangeDate indicates an expected call of GetAttendancesRangeDate.
func (mr *MockAttendanceRepoInterfaceMockRecorder) GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendancesRangeDate", reflect.TypeOf((*MockAttendanceRepoInterface)(nil).GetAttendancesRangeDate), employeeEmail, dateStart, dateEnd, status, officeId, order)
}

// IsCheckins mocks base method.
func (m *MockAttendanceRepoInterface) IsCheckins() ([]attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCheckins")
	ret0, _ := ret[0].([]attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCheckins indicates an expected call of IsCheckins.
func (mr *MockAttendanceRepoInterfaceMockRecorder) IsCheckins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCheckins", reflect.TypeOf((*MockAttendanceRepoInterface)(nil).IsCheckins))
}

// UpdateAttendance mocks base method.
func (m *MockAttendanceRepoInterface) UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAttendance", att)
	ret0, _ := ret[0].(attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAttendance indicates an expected call of UpdateAttendance.
func (mr *MockAttendanceRepoInterfaceMockRecorder) UpdateAttendance(att interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttendance", reflect.TypeOf((*MockAttendanceRepoInterface)(nil).UpdateAttendance), att)
}
