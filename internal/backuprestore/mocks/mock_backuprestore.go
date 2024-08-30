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

	v1 "github.com/openshift-kni/lifecycle-agent/api/imagebasedupgrade/v1"
	backuprestore "github.com/openshift-kni/lifecycle-agent/internal/backuprestore"
	v10 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	gomock "go.uber.org/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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

// CheckOadpMinimumVersion mocks base method.
func (m *MockBackuperRestorer) CheckOadpMinimumVersion(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOadpMinimumVersion", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOadpMinimumVersion indicates an expected call of CheckOadpMinimumVersion.
func (mr *MockBackuperRestorerMockRecorder) CheckOadpMinimumVersion(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOadpMinimumVersion", reflect.TypeOf((*MockBackuperRestorer)(nil).CheckOadpMinimumVersion), ctx)
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
func (m *MockBackuperRestorer) CleanupBackups(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupBackups", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupBackups indicates an expected call of CleanupBackups.
func (mr *MockBackuperRestorerMockRecorder) CleanupBackups(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupBackups", reflect.TypeOf((*MockBackuperRestorer)(nil).CleanupBackups), ctx)
}

// CleanupDeleteBackupRequests mocks base method.
func (m *MockBackuperRestorer) CleanupDeleteBackupRequests(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupDeleteBackupRequests", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupDeleteBackupRequests indicates an expected call of CleanupDeleteBackupRequests.
func (mr *MockBackuperRestorerMockRecorder) CleanupDeleteBackupRequests(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupDeleteBackupRequests", reflect.TypeOf((*MockBackuperRestorer)(nil).CleanupDeleteBackupRequests), ctx)
}

// CleanupStaleBackups mocks base method.
func (m *MockBackuperRestorer) CleanupStaleBackups(ctx context.Context, backups []*v10.Backup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupStaleBackups", ctx, backups)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupStaleBackups indicates an expected call of CleanupStaleBackups.
func (mr *MockBackuperRestorerMockRecorder) CleanupStaleBackups(ctx, backups any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupStaleBackups", reflect.TypeOf((*MockBackuperRestorer)(nil).CleanupStaleBackups), ctx, backups)
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
func (m *MockBackuperRestorer) ExportRestoresToDir(ctx context.Context, configMaps []v1.ConfigMapRef, toDir string) error {
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

// GetDataProtectionApplicationList mocks base method.
func (m *MockBackuperRestorer) GetDataProtectionApplicationList(ctx context.Context) (*unstructured.UnstructuredList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataProtectionApplicationList", ctx)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataProtectionApplicationList indicates an expected call of GetDataProtectionApplicationList.
func (mr *MockBackuperRestorerMockRecorder) GetDataProtectionApplicationList(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataProtectionApplicationList", reflect.TypeOf((*MockBackuperRestorer)(nil).GetDataProtectionApplicationList), ctx)
}

// GetSortedBackupsFromConfigmap mocks base method.
func (m *MockBackuperRestorer) GetSortedBackupsFromConfigmap(ctx context.Context, content []v1.ConfigMapRef) ([][]*v10.Backup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSortedBackupsFromConfigmap", ctx, content)
	ret0, _ := ret[0].([][]*v10.Backup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSortedBackupsFromConfigmap indicates an expected call of GetSortedBackupsFromConfigmap.
func (mr *MockBackuperRestorerMockRecorder) GetSortedBackupsFromConfigmap(ctx, content any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSortedBackupsFromConfigmap", reflect.TypeOf((*MockBackuperRestorer)(nil).GetSortedBackupsFromConfigmap), ctx, content)
}

// IsOadpInstalled mocks base method.
func (m *MockBackuperRestorer) IsOadpInstalled(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOadpInstalled", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOadpInstalled indicates an expected call of IsOadpInstalled.
func (mr *MockBackuperRestorerMockRecorder) IsOadpInstalled(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOadpInstalled", reflect.TypeOf((*MockBackuperRestorer)(nil).IsOadpInstalled), ctx)
}

// LoadRestoresFromOadpRestorePath mocks base method.
func (m *MockBackuperRestorer) LoadRestoresFromOadpRestorePath() ([][]*v10.Restore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadRestoresFromOadpRestorePath")
	ret0, _ := ret[0].([][]*v10.Restore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadRestoresFromOadpRestorePath indicates an expected call of LoadRestoresFromOadpRestorePath.
func (mr *MockBackuperRestorerMockRecorder) LoadRestoresFromOadpRestorePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadRestoresFromOadpRestorePath", reflect.TypeOf((*MockBackuperRestorer)(nil).LoadRestoresFromOadpRestorePath))
}

// PatchPVsReclaimPolicy mocks base method.
func (m *MockBackuperRestorer) PatchPVsReclaimPolicy(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchPVsReclaimPolicy", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchPVsReclaimPolicy indicates an expected call of PatchPVsReclaimPolicy.
func (mr *MockBackuperRestorerMockRecorder) PatchPVsReclaimPolicy(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchPVsReclaimPolicy", reflect.TypeOf((*MockBackuperRestorer)(nil).PatchPVsReclaimPolicy), ctx)
}

// RestorePVsReclaimPolicy mocks base method.
func (m *MockBackuperRestorer) RestorePVsReclaimPolicy(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestorePVsReclaimPolicy", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestorePVsReclaimPolicy indicates an expected call of RestorePVsReclaimPolicy.
func (mr *MockBackuperRestorerMockRecorder) RestorePVsReclaimPolicy(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestorePVsReclaimPolicy", reflect.TypeOf((*MockBackuperRestorer)(nil).RestorePVsReclaimPolicy), ctx)
}

// StartOrTrackBackup mocks base method.
func (m *MockBackuperRestorer) StartOrTrackBackup(ctx context.Context, backups []*v10.Backup) (*backuprestore.BackupTracker, error) {
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
func (m *MockBackuperRestorer) StartOrTrackRestore(ctx context.Context, restores []*v10.Restore) (*backuprestore.RestoreTracker, error) {
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

// ValidateOadpConfigmaps mocks base method.
func (m *MockBackuperRestorer) ValidateOadpConfigmaps(ctx context.Context, content []v1.ConfigMapRef) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateOadpConfigmaps", ctx, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateOadpConfigmaps indicates an expected call of ValidateOadpConfigmaps.
func (mr *MockBackuperRestorerMockRecorder) ValidateOadpConfigmaps(ctx, content any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateOadpConfigmaps", reflect.TypeOf((*MockBackuperRestorer)(nil).ValidateOadpConfigmaps), ctx, content)
}
