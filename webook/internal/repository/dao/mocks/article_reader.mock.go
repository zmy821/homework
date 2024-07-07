// Code generated by MockGen. DO NOT EDIT.
// Source: ./webook/internal/repository/dao/article_reader.go
//
// Generated by this command:
//
//	mockgen -source=./webook/internal/repository/dao/article_reader.go -package=daomocks -destination=./webook/internal/repository/dao/mocks/article_reader.mock.go
//
// Package daomocks is a generated GoMock package.
package daomocks

import (
	context "context"
	reflect "reflect"

	dao "gitee.com/geekbang/basic-go/webook/internal/repository/dao"
	gomock "go.uber.org/mock/gomock"
)

// MockArticleReaderDAO is a mock of ArticleReaderDAO interface.
type MockArticleReaderDAO struct {
	ctrl     *gomock.Controller
	recorder *MockArticleReaderDAOMockRecorder
}

// MockArticleReaderDAOMockRecorder is the mock recorder for MockArticleReaderDAO.
type MockArticleReaderDAOMockRecorder struct {
	mock *MockArticleReaderDAO
}

// NewMockArticleReaderDAO creates a new mock instance.
func NewMockArticleReaderDAO(ctrl *gomock.Controller) *MockArticleReaderDAO {
	mock := &MockArticleReaderDAO{ctrl: ctrl}
	mock.recorder = &MockArticleReaderDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleReaderDAO) EXPECT() *MockArticleReaderDAOMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockArticleReaderDAO) Upsert(ctx context.Context, art dao.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, art)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockArticleReaderDAOMockRecorder) Upsert(ctx, art any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockArticleReaderDAO)(nil).Upsert), ctx, art)
}

// UpsertV2 mocks base method.
func (m *MockArticleReaderDAO) UpsertV2(ctx context.Context, art dao.PublishedArticle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertV2", ctx, art)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertV2 indicates an expected call of UpsertV2.
func (mr *MockArticleReaderDAOMockRecorder) UpsertV2(ctx, art any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertV2", reflect.TypeOf((*MockArticleReaderDAO)(nil).UpsertV2), ctx, art)
}
