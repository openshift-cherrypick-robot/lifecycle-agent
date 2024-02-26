// Code generated by MockGen. DO NOT EDIT.
// Source: ../common.go
//
// Generated by this command:
//
//	mockgen -source ../common.go -destination mock_backuprestore.go -write_generate_directive
//
// Package mock_backuprestore is a generated GoMock package.
package mock_backuprestore

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/openshift-kni/lifecycle-agent/api/v1alpha1"
	backuprestore "github.com/openshift-kni/lifecycle-agent/internal/backuprestore"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -source ../common.go -destination mock_backuprestore.go -write_generate_directive

// MockBackuperRestorer is a mock of BackuperRestorer interface.
type MockBackuperRestorer struct {
	ctrl     *gomock.Controller
	recorder *MockBackuperRestorerMockRecorder
}

// MockBackuperRestorerMockRecorder is the mock recorder for MockBackuperRestorer.
type MockBackuperRestorerMockRecorder struct {
	mock *MockBackuperRestorer
}

// NewMockBackuperRestorer creates a new mock instance.
func NewMockBackuperRestorer(ctrl *gomock.Controller) *MockBackuperRestorer {
	mock := &MockBackuperRestorer{ctrl: ctrl}
	mock.recorder = &MockBackuperRestorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackuperRestorer) EXPECT() *MockBackuperRestorerMockRecorder {
	return m.recorder
}

// CheckOadpOperatorAvailability mocks base method.
func (m *MockBackuperRestorer) CheckOadpOperatorAvailability(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOadpOperatorAvailability", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckOadpOperatorAvailability indicates an expected call of CheckOadpOperatorAvailability.
func (mr *MockBackuperRestorerMockRecorder) CheckOadpOperatorAvailability(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOadpOperatorAvailability", reflect.TypeOf((*MockBackuperRestorer)(nil).CheckOadpOperatorAvailability), ctx)
}

// CleanupBackups mocks base method.
func (m *MockBackuperRestorer) CleanupBackups(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupBackups", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CleanupBackups indicates an expected call of CleanupBackups.
func (mr *MockBackuperRestorerMockRecorder) CleanupBackups(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupBackups", reflect.TypeOf((*MockBackuperRestorer)(nil).CleanupBackups), ctx)
}

// EnsureOadpConfiguration mocks base method.
func (m *MockBackuperRestorer) EnsureOadpConfiguration(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureOadpConfiguration", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureOadpConfiguration indicates an expected call of EnsureOadpConfiguration.
func (mr *MockBackuperRestorerMockRecorder) EnsureOadpConfiguration(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureOadpConfiguration", reflect.TypeOf((*MockBackuperRestorer)(nil).EnsureOadpConfiguration), ctx)
}

// ExportOadpConfigurationToDir mocks base method.
func (m *MockBackuperRestorer) ExportOadpConfigurationToDir(ctx context.Context, toDir, oadpNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportOadpConfigurationToDir", ctx, toDir, oadpNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExportOadpConfigurationToDir indicates an expected call of ExportOadpConfigurationToDir.
func (mr *MockBackuperRestorerMockRecorder) ExportOadpConfigurationToDir(ctx, toDir, oadpNamespace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportOadpConfigurationToDir", reflect.TypeOf((*MockBackuperRestorer)(nil).ExportOadpConfigurationToDir), ctx, toDir, oadpNamespace)
}

// ExportRestoresToDir mocks base method.
func (m *MockBackuperRestorer) ExportRestoresToDir(ctx context.Context, configMaps []v1alpha1.ConfigMapRef, toDir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportRestoresToDir", ctx, configMaps, toDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExportRestoresToDir indicates an expected call of ExportRestoresToDir.
func (mr *MockBackuperRestorerMockRecorder) ExportRestoresToDir(ctx, configMaps, toDir any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportRestoresToDir", reflect.TypeOf((*MockBackuperRestorer)(nil).ExportRestoresToDir), ctx, configMaps, toDir)
}

// GetSortedBackupsFromConfigmap mocks base method.
func (m *MockBackuperRestorer) GetSortedBackupsFromConfigmap(ctx context.Context, content []v1alpha1.ConfigMapRef) ([][]*v1.Backup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSortedBackupsFromConfigmap", ctx, content)
	ret0, _ := ret[0].([][]*v1.Backup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSortedBackupsFromConfigmap indicates an expected call of GetSortedBackupsFromConfigmap.
func (mr *MockBackuperRestorerMockRecorder) GetSortedBackupsFromConfigmap(ctx, content any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSortedBackupsFromConfigmap", reflect.TypeOf((*MockBackuperRestorer)(nil).GetSortedBackupsFromConfigmap), ctx, content)
}

// LoadRestoresFromOadpRestorePath mocks base method.
func (m *MockBackuperRestorer) LoadRestoresFromOadpRestorePath() ([][]*v1.Restore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadRestoresFromOadpRestorePath")
	ret0, _ := ret[0].([][]*v1.Restore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadRestoresFromOadpRestorePath indicates an expected call of LoadRestoresFromOadpRestorePath.
func (mr *MockBackuperRestorerMockRecorder) LoadRestoresFromOadpRestorePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadRestoresFromOadpRestorePath", reflect.TypeOf((*MockBackuperRestorer)(nil).LoadRestoresFromOadpRestorePath))
}

// StartOrTrackBackup mocks base method.
func (m *MockBackuperRestorer) StartOrTrackBackup(ctx context.Context, backups []*v1.Backup) (*backuprestore.BackupTracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartOrTrackBackup", ctx, backups)
	ret0, _ := ret[0].(*backuprestore.BackupTracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartOrTrackBackup indicates an expected call of StartOrTrackBackup.
func (mr *MockBackuperRestorerMockRecorder) StartOrTrackBackup(ctx, backups any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartOrTrackBackup", reflect.TypeOf((*MockBackuperRestorer)(nil).StartOrTrackBackup), ctx, backups)
}

// StartOrTrackRestore mocks base method.
func (m *MockBackuperRestorer) StartOrTrackRestore(ctx context.Context, restores []*v1.Restore) (*backuprestore.RestoreTracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartOrTrackRestore", ctx, restores)
	ret0, _ := ret[0].(*backuprestore.RestoreTracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartOrTrackRestore indicates an expected call of StartOrTrackRestore.
func (mr *MockBackuperRestorerMockRecorder) StartOrTrackRestore(ctx, restores any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartOrTrackRestore", reflect.TypeOf((*MockBackuperRestorer)(nil).StartOrTrackRestore), ctx, restores)
}

// ValidateOadpConfigmap mocks base method.
func (m *MockBackuperRestorer) ValidateOadpConfigmap(ctx context.Context, content []v1alpha1.ConfigMapRef) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateOadpConfigmap", ctx, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateOadpConfigmap indicates an expected call of ValidateOadpConfigmap.
func (mr *MockBackuperRestorerMockRecorder) ValidateOadpConfigmap(ctx, content any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateOadpConfigmap", reflect.TypeOf((*MockBackuperRestorer)(nil).ValidateOadpConfigmap), ctx, content)
}
