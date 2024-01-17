// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-mattermusic/server/trim (interfaces: Transcoder)

// Package mock_trim is a generated GoMock package.
package mock_trim

import (
	reflect "reflect"

	elastictranscoder "github.com/aws/aws-sdk-go/service/elastictranscoder"
	gomock "github.com/golang/mock/gomock"
)

// MockTranscoder is a mock of Transcoder interface.
type MockTranscoder struct {
	ctrl     *gomock.Controller
	recorder *MockTranscoderMockRecorder
}

// MockTranscoderMockRecorder is the mock recorder for MockTranscoder.
type MockTranscoderMockRecorder struct {
	mock *MockTranscoder
}

// NewMockTranscoder creates a new mock instance.
func NewMockTranscoder(ctrl *gomock.Controller) *MockTranscoder {
	mock := &MockTranscoder{ctrl: ctrl}
	mock.recorder = &MockTranscoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranscoder) EXPECT() *MockTranscoderMockRecorder {
	return m.recorder
}

// CreateJob mocks base method.
func (m *MockTranscoder) CreateJob(arg0 *elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", arg0)
	ret0, _ := ret[0].(*elastictranscoder.CreateJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob.
func (mr *MockTranscoderMockRecorder) CreateJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockTranscoder)(nil).CreateJob), arg0)
}

// CreatePipeline mocks base method.
func (m *MockTranscoder) CreatePipeline(arg0 *elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", arg0)
	ret0, _ := ret[0].(*elastictranscoder.CreatePipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePipeline indicates an expected call of CreatePipeline.
func (mr *MockTranscoderMockRecorder) CreatePipeline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockTranscoder)(nil).CreatePipeline), arg0)
}

// ListPipelines mocks base method.
func (m *MockTranscoder) ListPipelines(arg0 *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipelines", arg0)
	ret0, _ := ret[0].(*elastictranscoder.ListPipelinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPipelines indicates an expected call of ListPipelines.
func (mr *MockTranscoderMockRecorder) ListPipelines(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelines", reflect.TypeOf((*MockTranscoder)(nil).ListPipelines), arg0)
}

// ReadJob mocks base method.
func (m *MockTranscoder) ReadJob(arg0 *elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadJob", arg0)
	ret0, _ := ret[0].(*elastictranscoder.ReadJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadJob indicates an expected call of ReadJob.
func (mr *MockTranscoderMockRecorder) ReadJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadJob", reflect.TypeOf((*MockTranscoder)(nil).ReadJob), arg0)
}