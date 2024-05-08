// Code generated by MockGen. DO NOT EDIT.
// Source: ../extramanifest.go
//
// Generated by this command:
//
//	mockgen -source ../extramanifest.go -destination mock_extramanifest.go -write_generate_directive
//
// Package mock_extramanifest is a generated GoMock package.
package mock_extramanifest

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/openshift-kni/lifecycle-agent/api/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

//go:generate mockgen -source ../extramanifest.go -destination mock_extramanifest.go -write_generate_directive

// MockEManifestHandler is a mock of EManifestHandler interface.
type MockEManifestHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEManifestHandlerMockRecorder
}

// MockEManifestHandlerMockRecorder is the mock recorder for MockEManifestHandler.
type MockEManifestHandlerMockRecorder struct {
	mock *MockEManifestHandler
}

// NewMockEManifestHandler creates a new mock instance.
func NewMockEManifestHandler(ctrl *gomock.Controller) *MockEManifestHandler {
	mock := &MockEManifestHandler{ctrl: ctrl}
	mock.recorder = &MockEManifestHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEManifestHandler) EXPECT() *MockEManifestHandlerMockRecorder {
	return m.recorder
}

// ApplyExtraManifests mocks base method.
func (m *MockEManifestHandler) ApplyExtraManifests(ctx context.Context, fromDir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyExtraManifests", ctx, fromDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyExtraManifests indicates an expected call of ApplyExtraManifests.
func (mr *MockEManifestHandlerMockRecorder) ApplyExtraManifests(ctx, fromDir any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyExtraManifests", reflect.TypeOf((*MockEManifestHandler)(nil).ApplyExtraManifests), ctx, fromDir)
}

// ExportExtraManifestToDir mocks base method.
func (m *MockEManifestHandler) ExportExtraManifestToDir(ctx context.Context, extraManifestCMs []v1alpha1.ConfigMapRef, toDir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportExtraManifestToDir", ctx, extraManifestCMs, toDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExportExtraManifestToDir indicates an expected call of ExportExtraManifestToDir.
func (mr *MockEManifestHandlerMockRecorder) ExportExtraManifestToDir(ctx, extraManifestCMs, toDir any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportExtraManifestToDir", reflect.TypeOf((*MockEManifestHandler)(nil).ExportExtraManifestToDir), ctx, extraManifestCMs, toDir)
}

// ExtractAndExportManifestFromPoliciesToDir mocks base method.
func (m *MockEManifestHandler) ExtractAndExportManifestFromPoliciesToDir(ctx context.Context, policyLabels, objectLabels, validationAnns map[string]string, toDir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractAndExportManifestFromPoliciesToDir", ctx, policyLabels, objectLabels, validationAnns, toDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExtractAndExportManifestFromPoliciesToDir indicates an expected call of ExtractAndExportManifestFromPoliciesToDir.
func (mr *MockEManifestHandlerMockRecorder) ExtractAndExportManifestFromPoliciesToDir(ctx, policyLabels, objectLabels, validationAnns, toDir any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractAndExportManifestFromPoliciesToDir", reflect.TypeOf((*MockEManifestHandler)(nil).ExtractAndExportManifestFromPoliciesToDir), ctx, policyLabels, objectLabels, validationAnns, toDir)
}

// ValidateAndExtractManifestFromPolicies mocks base method.
func (m *MockEManifestHandler) ValidateAndExtractManifestFromPolicies(ctx context.Context, policyLabels, objectLabels, validationAnns map[string]string) ([][]*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAndExtractManifestFromPolicies", ctx, policyLabels, objectLabels, validationAnns)
	ret0, _ := ret[0].([][]*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateAndExtractManifestFromPolicies indicates an expected call of ValidateAndExtractManifestFromPolicies.
func (mr *MockEManifestHandlerMockRecorder) ValidateAndExtractManifestFromPolicies(ctx, policyLabels, objectLabels, validationAnns any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAndExtractManifestFromPolicies", reflect.TypeOf((*MockEManifestHandler)(nil).ValidateAndExtractManifestFromPolicies), ctx, policyLabels, objectLabels, validationAnns)
}

// ValidateExtraManifestConfigmaps mocks base method.
func (m *MockEManifestHandler) ValidateExtraManifestConfigmaps(ctx context.Context, extraManifestCMs []v1alpha1.ConfigMapRef, ibu *v1alpha1.ImageBasedUpgrade) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateExtraManifestConfigmaps", ctx, extraManifestCMs, ibu)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateExtraManifestConfigmaps indicates an expected call of ValidateExtraManifestConfigmaps.
func (mr *MockEManifestHandlerMockRecorder) ValidateExtraManifestConfigmaps(ctx, extraManifestCMs, ibu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateExtraManifestConfigmaps", reflect.TypeOf((*MockEManifestHandler)(nil).ValidateExtraManifestConfigmaps), ctx, extraManifestCMs, ibu)
}

// ValidateExtraManifestConfigmaps mocks base method.
func (m *MockEManifestHandler) ValidateExtraManifestConfigmaps(ctx context.Context, extraManifestCMs []v1alpha1.ConfigMapRef) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateExtraManifestConfigmaps", ctx, extraManifestCMs)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateExtraManifestConfigmaps indicates an expected call of ValidateExtraManifestConfigmaps.
func (mr *MockEManifestHandlerMockRecorder) ValidateExtraManifestConfigmaps(ctx, extraManifestCMs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateExtraManifestConfigmaps", reflect.TypeOf((*MockEManifestHandler)(nil).ValidateExtraManifestConfigmaps), ctx, extraManifestCMs)
}
