// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/ethereumapis/eth/v1alpha1 (interfaces: BeaconNodeValidatorClient,BeaconNodeValidator_WaitForSyncedClient,BeaconNodeValidator_WaitForChainStartClient,BeaconNodeValidator_WaitForActivationClient)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	eth "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockBeaconNodeValidatorClient is a mock of BeaconNodeValidatorClient interface
type MockBeaconNodeValidatorClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeValidatorClientMockRecorder
}

// MockBeaconNodeValidatorClientMockRecorder is the mock recorder for MockBeaconNodeValidatorClient
type MockBeaconNodeValidatorClientMockRecorder struct {
	mock *MockBeaconNodeValidatorClient
}

// NewMockBeaconNodeValidatorClient creates a new mock instance
func NewMockBeaconNodeValidatorClient(ctrl *gomock.Controller) *MockBeaconNodeValidatorClient {
	mock := &MockBeaconNodeValidatorClient{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeValidatorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconNodeValidatorClient) EXPECT() *MockBeaconNodeValidatorClientMockRecorder {
	return m.recorder
}

// DomainData mocks base method
func (m *MockBeaconNodeValidatorClient) DomainData(arg0 context.Context, arg1 *eth.DomainRequest, arg2 ...grpc.CallOption) (*eth.DomainResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DomainData", varargs...)
	ret0, _ := ret[0].(*eth.DomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainData indicates an expected call of DomainData
func (mr *MockBeaconNodeValidatorClientMockRecorder) DomainData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainData", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).DomainData), varargs...)
}

// GetAttestationData mocks base method
func (m *MockBeaconNodeValidatorClient) GetAttestationData(arg0 context.Context, arg1 *eth.AttestationDataRequest, arg2 ...grpc.CallOption) (*eth.AttestationData, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAttestationData", varargs...)
	ret0, _ := ret[0].(*eth.AttestationData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttestationData indicates an expected call of GetAttestationData
func (mr *MockBeaconNodeValidatorClientMockRecorder) GetAttestationData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttestationData", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).GetAttestationData), varargs...)
}

// GetBlock mocks base method
func (m *MockBeaconNodeValidatorClient) GetBlock(arg0 context.Context, arg1 *eth.BlockRequest, arg2 ...grpc.CallOption) (*eth.BeaconBlock, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlock", varargs...)
	ret0, _ := ret[0].(*eth.BeaconBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockBeaconNodeValidatorClientMockRecorder) GetBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).GetBlock), varargs...)
}

// GetDuties mocks base method
func (m *MockBeaconNodeValidatorClient) GetDuties(arg0 context.Context, arg1 *eth.DutiesRequest, arg2 ...grpc.CallOption) (*eth.DutiesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDuties", varargs...)
	ret0, _ := ret[0].(*eth.DutiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDuties indicates an expected call of GetDuties
func (mr *MockBeaconNodeValidatorClientMockRecorder) GetDuties(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDuties", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).GetDuties), varargs...)
}

// ProposeAttestation mocks base method
func (m *MockBeaconNodeValidatorClient) ProposeAttestation(arg0 context.Context, arg1 *eth.Attestation, arg2 ...grpc.CallOption) (*eth.AttestResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProposeAttestation", varargs...)
	ret0, _ := ret[0].(*eth.AttestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeAttestation indicates an expected call of ProposeAttestation
func (mr *MockBeaconNodeValidatorClientMockRecorder) ProposeAttestation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeAttestation", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).ProposeAttestation), varargs...)
}

// ProposeBlock mocks base method
func (m *MockBeaconNodeValidatorClient) ProposeBlock(arg0 context.Context, arg1 *eth.SignedBeaconBlock, arg2 ...grpc.CallOption) (*eth.ProposeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProposeBlock", varargs...)
	ret0, _ := ret[0].(*eth.ProposeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeBlock indicates an expected call of ProposeBlock
func (mr *MockBeaconNodeValidatorClientMockRecorder) ProposeBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeBlock", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).ProposeBlock), varargs...)
}

// ProposeExit mocks base method
func (m *MockBeaconNodeValidatorClient) ProposeExit(arg0 context.Context, arg1 *eth.SignedVoluntaryExit, arg2 ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProposeExit", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeExit indicates an expected call of ProposeExit
func (mr *MockBeaconNodeValidatorClientMockRecorder) ProposeExit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeExit", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).ProposeExit), varargs...)
}

// StreamDuties mocks base method
func (m *MockBeaconNodeValidatorClient) StreamDuties(arg0 context.Context, arg1 ...grpc.CallOption) (eth.BeaconNodeValidator_StreamDutiesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamDuties", varargs...)
	ret0, _ := ret[0].(eth.BeaconNodeValidator_StreamDutiesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamDuties indicates an expected call of StreamDuties
func (mr *MockBeaconNodeValidatorClientMockRecorder) StreamDuties(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamDuties", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).StreamDuties), varargs...)
}

// SubmitAggregateSelectionProof mocks base method
func (m *MockBeaconNodeValidatorClient) SubmitAggregateSelectionProof(arg0 context.Context, arg1 *eth.AggregateSelectionRequest, arg2 ...grpc.CallOption) (*eth.AggregateSelectionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitAggregateSelectionProof", varargs...)
	ret0, _ := ret[0].(*eth.AggregateSelectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitAggregateSelectionProof indicates an expected call of SubmitAggregateSelectionProof
func (mr *MockBeaconNodeValidatorClientMockRecorder) SubmitAggregateSelectionProof(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitAggregateSelectionProof", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).SubmitAggregateSelectionProof), varargs...)
}

// SubmitSignedAggregateSelectionProof mocks base method
func (m *MockBeaconNodeValidatorClient) SubmitSignedAggregateSelectionProof(arg0 context.Context, arg1 *eth.SignedAggregateSubmitRequest, arg2 ...grpc.CallOption) (*eth.SignedAggregateSubmitResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitSignedAggregateSelectionProof", varargs...)
	ret0, _ := ret[0].(*eth.SignedAggregateSubmitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitSignedAggregateSelectionProof indicates an expected call of SubmitSignedAggregateSelectionProof
func (mr *MockBeaconNodeValidatorClientMockRecorder) SubmitSignedAggregateSelectionProof(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitSignedAggregateSelectionProof", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).SubmitSignedAggregateSelectionProof), varargs...)
}

// SubscribeCommitteeSubnets mocks base method
func (m *MockBeaconNodeValidatorClient) SubscribeCommitteeSubnets(arg0 context.Context, arg1 *eth.CommitteeSubnetsSubscribeRequest, arg2 ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeCommitteeSubnets", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeCommitteeSubnets indicates an expected call of SubscribeCommitteeSubnets
func (mr *MockBeaconNodeValidatorClientMockRecorder) SubscribeCommitteeSubnets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeCommitteeSubnets", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).SubscribeCommitteeSubnets), varargs...)
}

// ValidatorIndex mocks base method
func (m *MockBeaconNodeValidatorClient) ValidatorIndex(arg0 context.Context, arg1 *eth.ValidatorIndexRequest, arg2 ...grpc.CallOption) (*eth.ValidatorIndexResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidatorIndex", varargs...)
	ret0, _ := ret[0].(*eth.ValidatorIndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorIndex indicates an expected call of ValidatorIndex
func (mr *MockBeaconNodeValidatorClientMockRecorder) ValidatorIndex(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorIndex", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).ValidatorIndex), varargs...)
}

// ValidatorStatus mocks base method
func (m *MockBeaconNodeValidatorClient) ValidatorStatus(arg0 context.Context, arg1 *eth.ValidatorStatusRequest, arg2 ...grpc.CallOption) (*eth.ValidatorStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidatorStatus", varargs...)
	ret0, _ := ret[0].(*eth.ValidatorStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorStatus indicates an expected call of ValidatorStatus
func (mr *MockBeaconNodeValidatorClientMockRecorder) ValidatorStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorStatus", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).ValidatorStatus), varargs...)
}

// WaitForActivation mocks base method
func (m *MockBeaconNodeValidatorClient) WaitForActivation(arg0 context.Context, arg1 *eth.ValidatorActivationRequest, arg2 ...grpc.CallOption) (eth.BeaconNodeValidator_WaitForActivationClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForActivation", varargs...)
	ret0, _ := ret[0].(eth.BeaconNodeValidator_WaitForActivationClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForActivation indicates an expected call of WaitForActivation
func (mr *MockBeaconNodeValidatorClientMockRecorder) WaitForActivation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForActivation", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).WaitForActivation), varargs...)
}

// WaitForChainStart mocks base method
func (m *MockBeaconNodeValidatorClient) WaitForChainStart(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (eth.BeaconNodeValidator_WaitForChainStartClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForChainStart", varargs...)
	ret0, _ := ret[0].(eth.BeaconNodeValidator_WaitForChainStartClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForChainStart indicates an expected call of WaitForChainStart
func (mr *MockBeaconNodeValidatorClientMockRecorder) WaitForChainStart(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForChainStart", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).WaitForChainStart), varargs...)
}

// WaitForSynced mocks base method
func (m *MockBeaconNodeValidatorClient) WaitForSynced(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (eth.BeaconNodeValidator_WaitForSyncedClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForSynced", varargs...)
	ret0, _ := ret[0].(eth.BeaconNodeValidator_WaitForSyncedClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForSynced indicates an expected call of WaitForSynced
func (mr *MockBeaconNodeValidatorClientMockRecorder) WaitForSynced(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForSynced", reflect.TypeOf((*MockBeaconNodeValidatorClient)(nil).WaitForSynced), varargs...)
}

// MockBeaconNodeValidator_WaitForSyncedClient is a mock of BeaconNodeValidator_WaitForSyncedClient interface
type MockBeaconNodeValidator_WaitForSyncedClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder
}

// MockBeaconNodeValidator_WaitForSyncedClientMockRecorder is the mock recorder for MockBeaconNodeValidator_WaitForSyncedClient
type MockBeaconNodeValidator_WaitForSyncedClientMockRecorder struct {
	mock *MockBeaconNodeValidator_WaitForSyncedClient
}

// NewMockBeaconNodeValidator_WaitForSyncedClient creates a new mock instance
func NewMockBeaconNodeValidator_WaitForSyncedClient(ctrl *gomock.Controller) *MockBeaconNodeValidator_WaitForSyncedClient {
	mock := &MockBeaconNodeValidator_WaitForSyncedClient{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeValidator_WaitForSyncedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconNodeValidator_WaitForSyncedClient) EXPECT() *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockBeaconNodeValidator_WaitForSyncedClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBeaconNodeValidator_WaitForSyncedClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockBeaconNodeValidator_WaitForSyncedClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconNodeValidator_WaitForSyncedClient)(nil).Context))
}

// Header mocks base method
func (m *MockBeaconNodeValidator_WaitForSyncedClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBeaconNodeValidator_WaitForSyncedClient)(nil).Header))
}

// Recv mocks base method
func (m *MockBeaconNodeValidator_WaitForSyncedClient) Recv() (*eth.SyncedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*eth.SyncedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBeaconNodeValidator_WaitForSyncedClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockBeaconNodeValidator_WaitForSyncedClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForSyncedClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconNodeValidator_WaitForSyncedClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForSyncedClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockBeaconNodeValidator_WaitForSyncedClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockBeaconNodeValidator_WaitForSyncedClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBeaconNodeValidator_WaitForSyncedClient)(nil).Trailer))
}

// MockBeaconNodeValidator_WaitForChainStartClient is a mock of BeaconNodeValidator_WaitForChainStartClient interface
type MockBeaconNodeValidator_WaitForChainStartClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder
}

// MockBeaconNodeValidator_WaitForChainStartClientMockRecorder is the mock recorder for MockBeaconNodeValidator_WaitForChainStartClient
type MockBeaconNodeValidator_WaitForChainStartClientMockRecorder struct {
	mock *MockBeaconNodeValidator_WaitForChainStartClient
}

// NewMockBeaconNodeValidator_WaitForChainStartClient creates a new mock instance
func NewMockBeaconNodeValidator_WaitForChainStartClient(ctrl *gomock.Controller) *MockBeaconNodeValidator_WaitForChainStartClient {
	mock := &MockBeaconNodeValidator_WaitForChainStartClient{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeValidator_WaitForChainStartClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconNodeValidator_WaitForChainStartClient) EXPECT() *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockBeaconNodeValidator_WaitForChainStartClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockBeaconNodeValidator_WaitForChainStartClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartClient)(nil).Context))
}

// Header mocks base method
func (m *MockBeaconNodeValidator_WaitForChainStartClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartClient)(nil).Header))
}

// Recv mocks base method
func (m *MockBeaconNodeValidator_WaitForChainStartClient) Recv() (*eth.ChainStartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*eth.ChainStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockBeaconNodeValidator_WaitForChainStartClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconNodeValidator_WaitForChainStartClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockBeaconNodeValidator_WaitForChainStartClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockBeaconNodeValidator_WaitForChainStartClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartClient)(nil).Trailer))
}

// MockBeaconNodeValidator_WaitForActivationClient is a mock of BeaconNodeValidator_WaitForActivationClient interface
type MockBeaconNodeValidator_WaitForActivationClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeValidator_WaitForActivationClientMockRecorder
}

// MockBeaconNodeValidator_WaitForActivationClientMockRecorder is the mock recorder for MockBeaconNodeValidator_WaitForActivationClient
type MockBeaconNodeValidator_WaitForActivationClientMockRecorder struct {
	mock *MockBeaconNodeValidator_WaitForActivationClient
}

// NewMockBeaconNodeValidator_WaitForActivationClient creates a new mock instance
func NewMockBeaconNodeValidator_WaitForActivationClient(ctrl *gomock.Controller) *MockBeaconNodeValidator_WaitForActivationClient {
	mock := &MockBeaconNodeValidator_WaitForActivationClient{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeValidator_WaitForActivationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconNodeValidator_WaitForActivationClient) EXPECT() *MockBeaconNodeValidator_WaitForActivationClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockBeaconNodeValidator_WaitForActivationClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockBeaconNodeValidator_WaitForActivationClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockBeaconNodeValidator_WaitForActivationClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconNodeValidator_WaitForActivationClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationClient)(nil).Context))
}

// Header mocks base method
func (m *MockBeaconNodeValidator_WaitForActivationClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockBeaconNodeValidator_WaitForActivationClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationClient)(nil).Header))
}

// Recv mocks base method
func (m *MockBeaconNodeValidator_WaitForActivationClient) Recv() (*eth.ValidatorActivationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*eth.ValidatorActivationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockBeaconNodeValidator_WaitForActivationClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockBeaconNodeValidator_WaitForActivationClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconNodeValidator_WaitForActivationClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconNodeValidator_WaitForActivationClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconNodeValidator_WaitForActivationClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockBeaconNodeValidator_WaitForActivationClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockBeaconNodeValidator_WaitForActivationClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationClient)(nil).Trailer))
}
