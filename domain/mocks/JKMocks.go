package mocks

import (
	"context"
	"testAPIYTS/domain"

	"github.com/stretchr/testify/mock"
)

type JKMocks struct {
	mock.Mock
}

func (m *JKMocks) GetAll(ctx context.Context) ([]domain.JadwalKajian, error) {
	ret := m.Called(ctx)

	var r0 []domain.JadwalKajian
	if rf, ok := ret.Get(0).(func(context.Context) []domain.JadwalKajian); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.JadwalKajian)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *JKMocks) Store(ctx context.Context, jk *domain.JadwalKajian) error {
	ret := m.Called(ctx, jk)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.JadwalKajian) error); ok {
		r0 = rf(ctx, jk)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
