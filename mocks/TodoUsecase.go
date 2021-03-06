// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/shandysiswandi/echo-service/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// TodoUsecase is an autogenerated mock type for the TodoUsecase type
type TodoUsecase struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: _a0, _a1
func (_m *TodoUsecase) CreateTodo(_a0 context.Context, _a1 domain.TodoCreatePayload) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.TodoCreatePayload) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTodoByID provides a mock function with given fields: _a0, _a1
func (_m *TodoUsecase) DeleteTodoByID(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchTodos provides a mock function with given fields: _a0
func (_m *TodoUsecase) FetchTodos(_a0 context.Context) ([]*domain.Todo, error) {
	ret := _m.Called(_a0)

	var r0 []*domain.Todo
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Todo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoByID provides a mock function with given fields: _a0, _a1
func (_m *TodoUsecase) GetTodoByID(_a0 context.Context, _a1 string) (*domain.Todo, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Todo
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Todo); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceTodoByID provides a mock function with given fields: _a0, _a1
func (_m *TodoUsecase) ReplaceTodoByID(_a0 context.Context, _a1 domain.TodoReplacePayload) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.TodoReplacePayload) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTodoByID provides a mock function with given fields: _a0, _a1
func (_m *TodoUsecase) UpdateTodoByID(_a0 context.Context, _a1 domain.TodoUpdatePayload) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.TodoUpdatePayload) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
