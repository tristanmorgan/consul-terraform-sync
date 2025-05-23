// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	api "github.com/hashicorp/consul/api"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ConsulClientInterface is an autogenerated mock type for the ConsulClientInterface type
type ConsulClientInterface struct {
	mock.Mock
}

type ConsulClientInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ConsulClientInterface) EXPECT() *ConsulClientInterface_Expecter {
	return &ConsulClientInterface_Expecter{mock: &_m.Mock}
}

// DeregisterService provides a mock function with given fields: ctx, serviceID, q
func (_m *ConsulClientInterface) DeregisterService(ctx context.Context, serviceID string, q *api.QueryOptions) error {
	ret := _m.Called(ctx, serviceID, q)

	if len(ret) == 0 {
		panic("no return value specified for DeregisterService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *api.QueryOptions) error); ok {
		r0 = rf(ctx, serviceID, q)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsulClientInterface_DeregisterService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeregisterService'
type ConsulClientInterface_DeregisterService_Call struct {
	*mock.Call
}

// DeregisterService is a helper method to define mock.On call
//   - ctx context.Context
//   - serviceID string
//   - q *api.QueryOptions
func (_e *ConsulClientInterface_Expecter) DeregisterService(ctx interface{}, serviceID interface{}, q interface{}) *ConsulClientInterface_DeregisterService_Call {
	return &ConsulClientInterface_DeregisterService_Call{Call: _e.mock.On("DeregisterService", ctx, serviceID, q)}
}

func (_c *ConsulClientInterface_DeregisterService_Call) Run(run func(ctx context.Context, serviceID string, q *api.QueryOptions)) *ConsulClientInterface_DeregisterService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*api.QueryOptions))
	})
	return _c
}

func (_c *ConsulClientInterface_DeregisterService_Call) Return(_a0 error) *ConsulClientInterface_DeregisterService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConsulClientInterface_DeregisterService_Call) RunAndReturn(run func(context.Context, string, *api.QueryOptions) error) *ConsulClientInterface_DeregisterService_Call {
	_c.Call.Return(run)
	return _c
}

// GetHealthChecks provides a mock function with given fields: ctx, serviceName, q
func (_m *ConsulClientInterface) GetHealthChecks(ctx context.Context, serviceName string, q *api.QueryOptions) (api.HealthChecks, error) {
	ret := _m.Called(ctx, serviceName, q)

	if len(ret) == 0 {
		panic("no return value specified for GetHealthChecks")
	}

	var r0 api.HealthChecks
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *api.QueryOptions) (api.HealthChecks, error)); ok {
		return rf(ctx, serviceName, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *api.QueryOptions) api.HealthChecks); ok {
		r0 = rf(ctx, serviceName, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.HealthChecks)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *api.QueryOptions) error); ok {
		r1 = rf(ctx, serviceName, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsulClientInterface_GetHealthChecks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHealthChecks'
type ConsulClientInterface_GetHealthChecks_Call struct {
	*mock.Call
}

// GetHealthChecks is a helper method to define mock.On call
//   - ctx context.Context
//   - serviceName string
//   - q *api.QueryOptions
func (_e *ConsulClientInterface_Expecter) GetHealthChecks(ctx interface{}, serviceName interface{}, q interface{}) *ConsulClientInterface_GetHealthChecks_Call {
	return &ConsulClientInterface_GetHealthChecks_Call{Call: _e.mock.On("GetHealthChecks", ctx, serviceName, q)}
}

func (_c *ConsulClientInterface_GetHealthChecks_Call) Run(run func(ctx context.Context, serviceName string, q *api.QueryOptions)) *ConsulClientInterface_GetHealthChecks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*api.QueryOptions))
	})
	return _c
}

func (_c *ConsulClientInterface_GetHealthChecks_Call) Return(_a0 api.HealthChecks, _a1 error) *ConsulClientInterface_GetHealthChecks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConsulClientInterface_GetHealthChecks_Call) RunAndReturn(run func(context.Context, string, *api.QueryOptions) (api.HealthChecks, error)) *ConsulClientInterface_GetHealthChecks_Call {
	_c.Call.Return(run)
	return _c
}

// GetLicense provides a mock function with given fields: ctx, q
func (_m *ConsulClientInterface) GetLicense(ctx context.Context, q *api.QueryOptions) (string, error) {
	ret := _m.Called(ctx, q)

	if len(ret) == 0 {
		panic("no return value specified for GetLicense")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.QueryOptions) (string, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.QueryOptions) string); ok {
		r0 = rf(ctx, q)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.QueryOptions) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsulClientInterface_GetLicense_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLicense'
type ConsulClientInterface_GetLicense_Call struct {
	*mock.Call
}

// GetLicense is a helper method to define mock.On call
//   - ctx context.Context
//   - q *api.QueryOptions
func (_e *ConsulClientInterface_Expecter) GetLicense(ctx interface{}, q interface{}) *ConsulClientInterface_GetLicense_Call {
	return &ConsulClientInterface_GetLicense_Call{Call: _e.mock.On("GetLicense", ctx, q)}
}

func (_c *ConsulClientInterface_GetLicense_Call) Run(run func(ctx context.Context, q *api.QueryOptions)) *ConsulClientInterface_GetLicense_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*api.QueryOptions))
	})
	return _c
}

func (_c *ConsulClientInterface_GetLicense_Call) Return(_a0 string, _a1 error) *ConsulClientInterface_GetLicense_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConsulClientInterface_GetLicense_Call) RunAndReturn(run func(context.Context, *api.QueryOptions) (string, error)) *ConsulClientInterface_GetLicense_Call {
	_c.Call.Return(run)
	return _c
}

// KVGet provides a mock function with given fields: ctx, key, q
func (_m *ConsulClientInterface) KVGet(ctx context.Context, key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error) {
	ret := _m.Called(ctx, key, q)

	if len(ret) == 0 {
		panic("no return value specified for KVGet")
	}

	var r0 *api.KVPair
	var r1 *api.QueryMeta
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error)); ok {
		return rf(ctx, key, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *api.QueryOptions) *api.KVPair); ok {
		r0 = rf(ctx, key, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.KVPair)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *api.QueryOptions) *api.QueryMeta); ok {
		r1 = rf(ctx, key, q)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*api.QueryMeta)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *api.QueryOptions) error); ok {
		r2 = rf(ctx, key, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ConsulClientInterface_KVGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KVGet'
type ConsulClientInterface_KVGet_Call struct {
	*mock.Call
}

// KVGet is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - q *api.QueryOptions
func (_e *ConsulClientInterface_Expecter) KVGet(ctx interface{}, key interface{}, q interface{}) *ConsulClientInterface_KVGet_Call {
	return &ConsulClientInterface_KVGet_Call{Call: _e.mock.On("KVGet", ctx, key, q)}
}

func (_c *ConsulClientInterface_KVGet_Call) Run(run func(ctx context.Context, key string, q *api.QueryOptions)) *ConsulClientInterface_KVGet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*api.QueryOptions))
	})
	return _c
}

func (_c *ConsulClientInterface_KVGet_Call) Return(_a0 *api.KVPair, _a1 *api.QueryMeta, _a2 error) *ConsulClientInterface_KVGet_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ConsulClientInterface_KVGet_Call) RunAndReturn(run func(context.Context, string, *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error)) *ConsulClientInterface_KVGet_Call {
	_c.Call.Return(run)
	return _c
}

// Lock provides a mock function with given fields: l, stopCh
func (_m *ConsulClientInterface) Lock(l *api.Lock, stopCh <-chan struct{}) (<-chan struct{}, error) {
	ret := _m.Called(l, stopCh)

	if len(ret) == 0 {
		panic("no return value specified for Lock")
	}

	var r0 <-chan struct{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*api.Lock, <-chan struct{}) (<-chan struct{}, error)); ok {
		return rf(l, stopCh)
	}
	if rf, ok := ret.Get(0).(func(*api.Lock, <-chan struct{}) <-chan struct{}); ok {
		r0 = rf(l, stopCh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func(*api.Lock, <-chan struct{}) error); ok {
		r1 = rf(l, stopCh)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsulClientInterface_Lock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Lock'
type ConsulClientInterface_Lock_Call struct {
	*mock.Call
}

// Lock is a helper method to define mock.On call
//   - l *api.Lock
//   - stopCh <-chan struct{}
func (_e *ConsulClientInterface_Expecter) Lock(l interface{}, stopCh interface{}) *ConsulClientInterface_Lock_Call {
	return &ConsulClientInterface_Lock_Call{Call: _e.mock.On("Lock", l, stopCh)}
}

func (_c *ConsulClientInterface_Lock_Call) Run(run func(l *api.Lock, stopCh <-chan struct{})) *ConsulClientInterface_Lock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*api.Lock), args[1].(<-chan struct{}))
	})
	return _c
}

func (_c *ConsulClientInterface_Lock_Call) Return(_a0 <-chan struct{}, _a1 error) *ConsulClientInterface_Lock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConsulClientInterface_Lock_Call) RunAndReturn(run func(*api.Lock, <-chan struct{}) (<-chan struct{}, error)) *ConsulClientInterface_Lock_Call {
	_c.Call.Return(run)
	return _c
}

// LockOpts provides a mock function with given fields: opts
func (_m *ConsulClientInterface) LockOpts(opts *api.LockOptions) (*api.Lock, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for LockOpts")
	}

	var r0 *api.Lock
	var r1 error
	if rf, ok := ret.Get(0).(func(*api.LockOptions) (*api.Lock, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*api.LockOptions) *api.Lock); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Lock)
		}
	}

	if rf, ok := ret.Get(1).(func(*api.LockOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsulClientInterface_LockOpts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LockOpts'
type ConsulClientInterface_LockOpts_Call struct {
	*mock.Call
}

// LockOpts is a helper method to define mock.On call
//   - opts *api.LockOptions
func (_e *ConsulClientInterface_Expecter) LockOpts(opts interface{}) *ConsulClientInterface_LockOpts_Call {
	return &ConsulClientInterface_LockOpts_Call{Call: _e.mock.On("LockOpts", opts)}
}

func (_c *ConsulClientInterface_LockOpts_Call) Run(run func(opts *api.LockOptions)) *ConsulClientInterface_LockOpts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*api.LockOptions))
	})
	return _c
}

func (_c *ConsulClientInterface_LockOpts_Call) Return(_a0 *api.Lock, _a1 error) *ConsulClientInterface_LockOpts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConsulClientInterface_LockOpts_Call) RunAndReturn(run func(*api.LockOptions) (*api.Lock, error)) *ConsulClientInterface_LockOpts_Call {
	_c.Call.Return(run)
	return _c
}

// QueryServices provides a mock function with given fields: ctx, filter, q
func (_m *ConsulClientInterface) QueryServices(ctx context.Context, filter string, q *api.QueryOptions) ([]*api.AgentService, error) {
	ret := _m.Called(ctx, filter, q)

	if len(ret) == 0 {
		panic("no return value specified for QueryServices")
	}

	var r0 []*api.AgentService
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *api.QueryOptions) ([]*api.AgentService, error)); ok {
		return rf(ctx, filter, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *api.QueryOptions) []*api.AgentService); ok {
		r0 = rf(ctx, filter, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.AgentService)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *api.QueryOptions) error); ok {
		r1 = rf(ctx, filter, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsulClientInterface_QueryServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryServices'
type ConsulClientInterface_QueryServices_Call struct {
	*mock.Call
}

// QueryServices is a helper method to define mock.On call
//   - ctx context.Context
//   - filter string
//   - q *api.QueryOptions
func (_e *ConsulClientInterface_Expecter) QueryServices(ctx interface{}, filter interface{}, q interface{}) *ConsulClientInterface_QueryServices_Call {
	return &ConsulClientInterface_QueryServices_Call{Call: _e.mock.On("QueryServices", ctx, filter, q)}
}

func (_c *ConsulClientInterface_QueryServices_Call) Run(run func(ctx context.Context, filter string, q *api.QueryOptions)) *ConsulClientInterface_QueryServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*api.QueryOptions))
	})
	return _c
}

func (_c *ConsulClientInterface_QueryServices_Call) Return(_a0 []*api.AgentService, _a1 error) *ConsulClientInterface_QueryServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConsulClientInterface_QueryServices_Call) RunAndReturn(run func(context.Context, string, *api.QueryOptions) ([]*api.AgentService, error)) *ConsulClientInterface_QueryServices_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterService provides a mock function with given fields: ctx, s
func (_m *ConsulClientInterface) RegisterService(ctx context.Context, s *api.AgentServiceRegistration) error {
	ret := _m.Called(ctx, s)

	if len(ret) == 0 {
		panic("no return value specified for RegisterService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.AgentServiceRegistration) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsulClientInterface_RegisterService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterService'
type ConsulClientInterface_RegisterService_Call struct {
	*mock.Call
}

// RegisterService is a helper method to define mock.On call
//   - ctx context.Context
//   - s *api.AgentServiceRegistration
func (_e *ConsulClientInterface_Expecter) RegisterService(ctx interface{}, s interface{}) *ConsulClientInterface_RegisterService_Call {
	return &ConsulClientInterface_RegisterService_Call{Call: _e.mock.On("RegisterService", ctx, s)}
}

func (_c *ConsulClientInterface_RegisterService_Call) Run(run func(ctx context.Context, s *api.AgentServiceRegistration)) *ConsulClientInterface_RegisterService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*api.AgentServiceRegistration))
	})
	return _c
}

func (_c *ConsulClientInterface_RegisterService_Call) Return(_a0 error) *ConsulClientInterface_RegisterService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConsulClientInterface_RegisterService_Call) RunAndReturn(run func(context.Context, *api.AgentServiceRegistration) error) *ConsulClientInterface_RegisterService_Call {
	_c.Call.Return(run)
	return _c
}

// SessionCreate provides a mock function with given fields: ctx, se, q
func (_m *ConsulClientInterface) SessionCreate(ctx context.Context, se *api.SessionEntry, q *api.WriteOptions) (string, *api.WriteMeta, error) {
	ret := _m.Called(ctx, se, q)

	if len(ret) == 0 {
		panic("no return value specified for SessionCreate")
	}

	var r0 string
	var r1 *api.WriteMeta
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.SessionEntry, *api.WriteOptions) (string, *api.WriteMeta, error)); ok {
		return rf(ctx, se, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.SessionEntry, *api.WriteOptions) string); ok {
		r0 = rf(ctx, se, q)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.SessionEntry, *api.WriteOptions) *api.WriteMeta); ok {
		r1 = rf(ctx, se, q)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*api.WriteMeta)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *api.SessionEntry, *api.WriteOptions) error); ok {
		r2 = rf(ctx, se, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ConsulClientInterface_SessionCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SessionCreate'
type ConsulClientInterface_SessionCreate_Call struct {
	*mock.Call
}

// SessionCreate is a helper method to define mock.On call
//   - ctx context.Context
//   - se *api.SessionEntry
//   - q *api.WriteOptions
func (_e *ConsulClientInterface_Expecter) SessionCreate(ctx interface{}, se interface{}, q interface{}) *ConsulClientInterface_SessionCreate_Call {
	return &ConsulClientInterface_SessionCreate_Call{Call: _e.mock.On("SessionCreate", ctx, se, q)}
}

func (_c *ConsulClientInterface_SessionCreate_Call) Run(run func(ctx context.Context, se *api.SessionEntry, q *api.WriteOptions)) *ConsulClientInterface_SessionCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*api.SessionEntry), args[2].(*api.WriteOptions))
	})
	return _c
}

func (_c *ConsulClientInterface_SessionCreate_Call) Return(_a0 string, _a1 *api.WriteMeta, _a2 error) *ConsulClientInterface_SessionCreate_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ConsulClientInterface_SessionCreate_Call) RunAndReturn(run func(context.Context, *api.SessionEntry, *api.WriteOptions) (string, *api.WriteMeta, error)) *ConsulClientInterface_SessionCreate_Call {
	_c.Call.Return(run)
	return _c
}

// SessionRenewPeriodic provides a mock function with given fields: initialTTL, id, q, doneCh
func (_m *ConsulClientInterface) SessionRenewPeriodic(initialTTL string, id string, q *api.WriteOptions, doneCh <-chan struct{}) error {
	ret := _m.Called(initialTTL, id, q, doneCh)

	if len(ret) == 0 {
		panic("no return value specified for SessionRenewPeriodic")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *api.WriteOptions, <-chan struct{}) error); ok {
		r0 = rf(initialTTL, id, q, doneCh)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsulClientInterface_SessionRenewPeriodic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SessionRenewPeriodic'
type ConsulClientInterface_SessionRenewPeriodic_Call struct {
	*mock.Call
}

// SessionRenewPeriodic is a helper method to define mock.On call
//   - initialTTL string
//   - id string
//   - q *api.WriteOptions
//   - doneCh <-chan struct{}
func (_e *ConsulClientInterface_Expecter) SessionRenewPeriodic(initialTTL interface{}, id interface{}, q interface{}, doneCh interface{}) *ConsulClientInterface_SessionRenewPeriodic_Call {
	return &ConsulClientInterface_SessionRenewPeriodic_Call{Call: _e.mock.On("SessionRenewPeriodic", initialTTL, id, q, doneCh)}
}

func (_c *ConsulClientInterface_SessionRenewPeriodic_Call) Run(run func(initialTTL string, id string, q *api.WriteOptions, doneCh <-chan struct{})) *ConsulClientInterface_SessionRenewPeriodic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(*api.WriteOptions), args[3].(<-chan struct{}))
	})
	return _c
}

func (_c *ConsulClientInterface_SessionRenewPeriodic_Call) Return(_a0 error) *ConsulClientInterface_SessionRenewPeriodic_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConsulClientInterface_SessionRenewPeriodic_Call) RunAndReturn(run func(string, string, *api.WriteOptions, <-chan struct{}) error) *ConsulClientInterface_SessionRenewPeriodic_Call {
	_c.Call.Return(run)
	return _c
}

// Unlock provides a mock function with given fields: l
func (_m *ConsulClientInterface) Unlock(l *api.Lock) error {
	ret := _m.Called(l)

	if len(ret) == 0 {
		panic("no return value specified for Unlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*api.Lock) error); ok {
		r0 = rf(l)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsulClientInterface_Unlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unlock'
type ConsulClientInterface_Unlock_Call struct {
	*mock.Call
}

// Unlock is a helper method to define mock.On call
//   - l *api.Lock
func (_e *ConsulClientInterface_Expecter) Unlock(l interface{}) *ConsulClientInterface_Unlock_Call {
	return &ConsulClientInterface_Unlock_Call{Call: _e.mock.On("Unlock", l)}
}

func (_c *ConsulClientInterface_Unlock_Call) Run(run func(l *api.Lock)) *ConsulClientInterface_Unlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*api.Lock))
	})
	return _c
}

func (_c *ConsulClientInterface_Unlock_Call) Return(_a0 error) *ConsulClientInterface_Unlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConsulClientInterface_Unlock_Call) RunAndReturn(run func(*api.Lock) error) *ConsulClientInterface_Unlock_Call {
	_c.Call.Return(run)
	return _c
}

// NewConsulClientInterface creates a new instance of ConsulClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsulClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConsulClientInterface {
	mock := &ConsulClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
