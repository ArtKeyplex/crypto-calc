// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "exchange-rate-calculator/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// CurrencyRepository is an autogenerated mock type for the CurrencyRepository type
type CurrencyRepository struct {
	mock.Mock
}

// GetAllCurrencies provides a mock function with given fields: ctx
func (_m *CurrencyRepository) GetAllCurrencies(ctx context.Context) ([]*entity.Currency, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCurrencies")
	}

	var r0 []*entity.Currency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.Currency, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Currency); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Currency)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrencyByCode provides a mock function with given fields: ctx, code
func (_m *CurrencyRepository) GetCurrencyByCode(ctx context.Context, code string) (*entity.Currency, error) {
	ret := _m.Called(ctx, code)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrencyByCode")
	}

	var r0 *entity.Currency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Currency, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Currency); ok {
		r0 = rf(ctx, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Currency)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRates provides a mock function with given fields: ctx, currencies
func (_m *CurrencyRepository) UpdateRates(ctx context.Context, currencies []*entity.Currency) error {
	ret := _m.Called(ctx, currencies)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRates")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*entity.Currency) error); ok {
		r0 = rf(ctx, currencies)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCurrencyRepository creates a new instance of CurrencyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCurrencyRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CurrencyRepository {
	mock := &CurrencyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
