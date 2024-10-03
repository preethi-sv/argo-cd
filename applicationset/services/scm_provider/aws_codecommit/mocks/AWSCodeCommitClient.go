// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	codecommit "github.com/aws/aws-sdk-go/service/codecommit"

	mock "github.com/stretchr/testify/mock"

	request "github.com/aws/aws-sdk-go/aws/request"
)

// AWSCodeCommitClient is an autogenerated mock type for the AWSCodeCommitClient type
type AWSCodeCommitClient struct {
	mock.Mock
}

type AWSCodeCommitClient_Expecter struct {
	mock *mock.Mock
}

func (_m *AWSCodeCommitClient) EXPECT() *AWSCodeCommitClient_Expecter {
	return &AWSCodeCommitClient_Expecter{mock: &_m.Mock}
}

// GetFolderWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *AWSCodeCommitClient) GetFolderWithContext(_a0 context.Context, _a1 *codecommit.GetFolderInput, _a2 ...request.Option) (*codecommit.GetFolderOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *codecommit.GetFolderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codecommit.GetFolderInput, ...request.Option) (*codecommit.GetFolderOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codecommit.GetFolderInput, ...request.Option) *codecommit.GetFolderOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codecommit.GetFolderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codecommit.GetFolderInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AWSCodeCommitClient_GetFolderWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFolderWithContext'
type AWSCodeCommitClient_GetFolderWithContext_Call struct {
	*mock.Call
}

// GetFolderWithContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *codecommit.GetFolderInput
//   - _a2 ...request.Option
func (_e *AWSCodeCommitClient_Expecter) GetFolderWithContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *AWSCodeCommitClient_GetFolderWithContext_Call {
	return &AWSCodeCommitClient_GetFolderWithContext_Call{Call: _e.mock.On("GetFolderWithContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *AWSCodeCommitClient_GetFolderWithContext_Call) Run(run func(_a0 context.Context, _a1 *codecommit.GetFolderInput, _a2 ...request.Option)) *AWSCodeCommitClient_GetFolderWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]request.Option, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(request.Option)
			}
		}
		run(args[0].(context.Context), args[1].(*codecommit.GetFolderInput), variadicArgs...)
	})
	return _c
}

func (_c *AWSCodeCommitClient_GetFolderWithContext_Call) Return(_a0 *codecommit.GetFolderOutput, _a1 error) *AWSCodeCommitClient_GetFolderWithContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AWSCodeCommitClient_GetFolderWithContext_Call) RunAndReturn(run func(context.Context, *codecommit.GetFolderInput, ...request.Option) (*codecommit.GetFolderOutput, error)) *AWSCodeCommitClient_GetFolderWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetRepositoryWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *AWSCodeCommitClient) GetRepositoryWithContext(_a0 context.Context, _a1 *codecommit.GetRepositoryInput, _a2 ...request.Option) (*codecommit.GetRepositoryOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *codecommit.GetRepositoryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codecommit.GetRepositoryInput, ...request.Option) (*codecommit.GetRepositoryOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codecommit.GetRepositoryInput, ...request.Option) *codecommit.GetRepositoryOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codecommit.GetRepositoryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codecommit.GetRepositoryInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AWSCodeCommitClient_GetRepositoryWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRepositoryWithContext'
type AWSCodeCommitClient_GetRepositoryWithContext_Call struct {
	*mock.Call
}

// GetRepositoryWithContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *codecommit.GetRepositoryInput
//   - _a2 ...request.Option
func (_e *AWSCodeCommitClient_Expecter) GetRepositoryWithContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *AWSCodeCommitClient_GetRepositoryWithContext_Call {
	return &AWSCodeCommitClient_GetRepositoryWithContext_Call{Call: _e.mock.On("GetRepositoryWithContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *AWSCodeCommitClient_GetRepositoryWithContext_Call) Run(run func(_a0 context.Context, _a1 *codecommit.GetRepositoryInput, _a2 ...request.Option)) *AWSCodeCommitClient_GetRepositoryWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]request.Option, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(request.Option)
			}
		}
		run(args[0].(context.Context), args[1].(*codecommit.GetRepositoryInput), variadicArgs...)
	})
	return _c
}

func (_c *AWSCodeCommitClient_GetRepositoryWithContext_Call) Return(_a0 *codecommit.GetRepositoryOutput, _a1 error) *AWSCodeCommitClient_GetRepositoryWithContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AWSCodeCommitClient_GetRepositoryWithContext_Call) RunAndReturn(run func(context.Context, *codecommit.GetRepositoryInput, ...request.Option) (*codecommit.GetRepositoryOutput, error)) *AWSCodeCommitClient_GetRepositoryWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// ListBranchesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *AWSCodeCommitClient) ListBranchesWithContext(_a0 context.Context, _a1 *codecommit.ListBranchesInput, _a2 ...request.Option) (*codecommit.ListBranchesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *codecommit.ListBranchesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codecommit.ListBranchesInput, ...request.Option) (*codecommit.ListBranchesOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codecommit.ListBranchesInput, ...request.Option) *codecommit.ListBranchesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codecommit.ListBranchesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codecommit.ListBranchesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AWSCodeCommitClient_ListBranchesWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBranchesWithContext'
type AWSCodeCommitClient_ListBranchesWithContext_Call struct {
	*mock.Call
}

// ListBranchesWithContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *codecommit.ListBranchesInput
//   - _a2 ...request.Option
func (_e *AWSCodeCommitClient_Expecter) ListBranchesWithContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *AWSCodeCommitClient_ListBranchesWithContext_Call {
	return &AWSCodeCommitClient_ListBranchesWithContext_Call{Call: _e.mock.On("ListBranchesWithContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *AWSCodeCommitClient_ListBranchesWithContext_Call) Run(run func(_a0 context.Context, _a1 *codecommit.ListBranchesInput, _a2 ...request.Option)) *AWSCodeCommitClient_ListBranchesWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]request.Option, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(request.Option)
			}
		}
		run(args[0].(context.Context), args[1].(*codecommit.ListBranchesInput), variadicArgs...)
	})
	return _c
}

func (_c *AWSCodeCommitClient_ListBranchesWithContext_Call) Return(_a0 *codecommit.ListBranchesOutput, _a1 error) *AWSCodeCommitClient_ListBranchesWithContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AWSCodeCommitClient_ListBranchesWithContext_Call) RunAndReturn(run func(context.Context, *codecommit.ListBranchesInput, ...request.Option) (*codecommit.ListBranchesOutput, error)) *AWSCodeCommitClient_ListBranchesWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// ListRepositoriesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *AWSCodeCommitClient) ListRepositoriesWithContext(_a0 context.Context, _a1 *codecommit.ListRepositoriesInput, _a2 ...request.Option) (*codecommit.ListRepositoriesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *codecommit.ListRepositoriesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codecommit.ListRepositoriesInput, ...request.Option) (*codecommit.ListRepositoriesOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codecommit.ListRepositoriesInput, ...request.Option) *codecommit.ListRepositoriesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codecommit.ListRepositoriesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codecommit.ListRepositoriesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AWSCodeCommitClient_ListRepositoriesWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListRepositoriesWithContext'
type AWSCodeCommitClient_ListRepositoriesWithContext_Call struct {
	*mock.Call
}

// ListRepositoriesWithContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *codecommit.ListRepositoriesInput
//   - _a2 ...request.Option
func (_e *AWSCodeCommitClient_Expecter) ListRepositoriesWithContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *AWSCodeCommitClient_ListRepositoriesWithContext_Call {
	return &AWSCodeCommitClient_ListRepositoriesWithContext_Call{Call: _e.mock.On("ListRepositoriesWithContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *AWSCodeCommitClient_ListRepositoriesWithContext_Call) Run(run func(_a0 context.Context, _a1 *codecommit.ListRepositoriesInput, _a2 ...request.Option)) *AWSCodeCommitClient_ListRepositoriesWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]request.Option, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(request.Option)
			}
		}
		run(args[0].(context.Context), args[1].(*codecommit.ListRepositoriesInput), variadicArgs...)
	})
	return _c
}

func (_c *AWSCodeCommitClient_ListRepositoriesWithContext_Call) Return(_a0 *codecommit.ListRepositoriesOutput, _a1 error) *AWSCodeCommitClient_ListRepositoriesWithContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AWSCodeCommitClient_ListRepositoriesWithContext_Call) RunAndReturn(run func(context.Context, *codecommit.ListRepositoriesInput, ...request.Option) (*codecommit.ListRepositoriesOutput, error)) *AWSCodeCommitClient_ListRepositoriesWithContext_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAWSCodeCommitClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewAWSCodeCommitClient creates a new instance of AWSCodeCommitClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAWSCodeCommitClient(t mockConstructorTestingTNewAWSCodeCommitClient) *AWSCodeCommitClient {
	mock := &AWSCodeCommitClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
