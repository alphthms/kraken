// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/kraken/tracker/announceclient (interfaces: Client)
//
// Generated by this command:
//
//	mockgen -typed -package mockannounceclient . Client
//

// Package mockannounceclient is a generated GoMock package.
package mockannounceclient

import (
	gomock "github.com/golang/mock/gomock"
	core "github.com/uber/kraken/core"
	reflect "reflect"
	time "time"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Announce mocks base method.
func (m *MockClient) Announce(d core.Digest, h core.InfoHash, complete bool, version int) ([]*core.PeerInfo, time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Announce", d, h, complete, version)
	ret0, _ := ret[0].([]*core.PeerInfo)
	ret1, _ := ret[1].(time.Duration)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Announce indicates an expected call of Announce.
func (mr *MockClientMockRecorder) Announce(d, h, complete, version interface{}) *MockClientAnnounceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Announce", reflect.TypeOf((*MockClient)(nil).Announce), d, h, complete, version)
	return &MockClientAnnounceCall{Call: call}
}

// MockClientAnnounceCall wrap *gomock.Call
type MockClientAnnounceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientAnnounceCall) Return(arg0 []*core.PeerInfo, arg1 time.Duration, arg2 error) *MockClientAnnounceCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientAnnounceCall) Do(f func(core.Digest, core.InfoHash, bool, int) ([]*core.PeerInfo, time.Duration, error)) *MockClientAnnounceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientAnnounceCall) DoAndReturn(f func(core.Digest, core.InfoHash, bool, int) ([]*core.PeerInfo, time.Duration, error)) *MockClientAnnounceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CheckHealth mocks base method.
func (m *MockClient) CheckHealth() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealth")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckHealth indicates an expected call of CheckHealth.
func (mr *MockClientMockRecorder) CheckHealth() *MockClientCheckHealthCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockClient)(nil).CheckHealth))
	return &MockClientCheckHealthCall{Call: call}
}

// MockClientCheckHealthCall wrap *gomock.Call
type MockClientCheckHealthCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientCheckHealthCall) Return(arg0 error) *MockClientCheckHealthCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientCheckHealthCall) Do(f func() error) *MockClientCheckHealthCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientCheckHealthCall) DoAndReturn(f func() error) *MockClientCheckHealthCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
