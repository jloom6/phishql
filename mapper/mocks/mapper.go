// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jloom6/phishql/mapper (interfaces: Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	phishql "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	structs "github.com/jloom6/phishql/structs"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// ArtistsToProto mocks base method
func (m *MockInterface) ArtistsToProto(arg0 []structs.Artist) []*phishql.Artist {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArtistsToProto", arg0)
	ret0, _ := ret[0].([]*phishql.Artist)
	return ret0
}

// ArtistsToProto indicates an expected call of ArtistsToProto
func (mr *MockInterfaceMockRecorder) ArtistsToProto(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArtistsToProto", reflect.TypeOf((*MockInterface)(nil).ArtistsToProto), arg0)
}

// ProtoToGetArtistsRequest mocks base method
func (m *MockInterface) ProtoToGetArtistsRequest(arg0 *phishql.GetArtistsRequest) structs.GetArtistsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtoToGetArtistsRequest", arg0)
	ret0, _ := ret[0].(structs.GetArtistsRequest)
	return ret0
}

// ProtoToGetArtistsRequest indicates an expected call of ProtoToGetArtistsRequest
func (mr *MockInterfaceMockRecorder) ProtoToGetArtistsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoToGetArtistsRequest", reflect.TypeOf((*MockInterface)(nil).ProtoToGetArtistsRequest), arg0)
}

// ProtoToGetShowsRequest mocks base method
func (m *MockInterface) ProtoToGetShowsRequest(arg0 *phishql.GetShowsRequest) structs.GetShowsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtoToGetShowsRequest", arg0)
	ret0, _ := ret[0].(structs.GetShowsRequest)
	return ret0
}

// ProtoToGetShowsRequest indicates an expected call of ProtoToGetShowsRequest
func (mr *MockInterfaceMockRecorder) ProtoToGetShowsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoToGetShowsRequest", reflect.TypeOf((*MockInterface)(nil).ProtoToGetShowsRequest), arg0)
}

// ProtoToGetSongsRequest mocks base method
func (m *MockInterface) ProtoToGetSongsRequest(arg0 *phishql.GetSongsRequest) structs.GetSongsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtoToGetSongsRequest", arg0)
	ret0, _ := ret[0].(structs.GetSongsRequest)
	return ret0
}

// ProtoToGetSongsRequest indicates an expected call of ProtoToGetSongsRequest
func (mr *MockInterfaceMockRecorder) ProtoToGetSongsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoToGetSongsRequest", reflect.TypeOf((*MockInterface)(nil).ProtoToGetSongsRequest), arg0)
}

// ProtoToGetTagsRequest mocks base method
func (m *MockInterface) ProtoToGetTagsRequest(arg0 *phishql.GetTagsRequest) structs.GetTagsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtoToGetTagsRequest", arg0)
	ret0, _ := ret[0].(structs.GetTagsRequest)
	return ret0
}

// ProtoToGetTagsRequest indicates an expected call of ProtoToGetTagsRequest
func (mr *MockInterfaceMockRecorder) ProtoToGetTagsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoToGetTagsRequest", reflect.TypeOf((*MockInterface)(nil).ProtoToGetTagsRequest), arg0)
}

// ShowsToProto mocks base method
func (m *MockInterface) ShowsToProto(arg0 []structs.Show) ([]*phishql.Show, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowsToProto", arg0)
	ret0, _ := ret[0].([]*phishql.Show)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowsToProto indicates an expected call of ShowsToProto
func (mr *MockInterfaceMockRecorder) ShowsToProto(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowsToProto", reflect.TypeOf((*MockInterface)(nil).ShowsToProto), arg0)
}

// SongsToProto mocks base method
func (m *MockInterface) SongsToProto(arg0 []structs.Song) []*phishql.Song {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SongsToProto", arg0)
	ret0, _ := ret[0].([]*phishql.Song)
	return ret0
}

// SongsToProto indicates an expected call of SongsToProto
func (mr *MockInterfaceMockRecorder) SongsToProto(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SongsToProto", reflect.TypeOf((*MockInterface)(nil).SongsToProto), arg0)
}

// TagsToProto mocks base method
func (m *MockInterface) TagsToProto(arg0 []structs.Tag) []*phishql.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagsToProto", arg0)
	ret0, _ := ret[0].([]*phishql.Tag)
	return ret0
}

// TagsToProto indicates an expected call of TagsToProto
func (mr *MockInterfaceMockRecorder) TagsToProto(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagsToProto", reflect.TypeOf((*MockInterface)(nil).TagsToProto), arg0)
}
