package mocks

import "github.com/gambol99/go-marathon"
import "github.com/stretchr/testify/mock"

import "net/url"

import "time"

type Marathon struct {
	mock.Mock
}

// ListApplications provides a mock function with given fields: _a0
func (_m *Marathon) ListApplications(_a0 url.Values) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(url.Values) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(url.Values) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplicationVersions provides a mock function with given fields: name
func (_m *Marathon) ApplicationVersions(name string) (*marathon.ApplicationVersions, error) {
	ret := _m.Called(name)

	var r0 *marathon.ApplicationVersions
	if rf, ok := ret.Get(0).(func(string) *marathon.ApplicationVersions); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.ApplicationVersions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasApplicationVersion provides a mock function with given fields: name, version
func (_m *Marathon) HasApplicationVersion(name string, version string) (bool, error) {
	ret := _m.Called(name, version)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(name, version)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetApplicationVersion provides a mock function with given fields: name, version
func (_m *Marathon) SetApplicationVersion(name string, version *marathon.ApplicationVersion) (*marathon.DeploymentID, error) {
	ret := _m.Called(name, version)

	var r0 *marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(string, *marathon.ApplicationVersion) *marathon.DeploymentID); ok {
		r0 = rf(name, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *marathon.ApplicationVersion) error); ok {
		r1 = rf(name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplicationOK provides a mock function with given fields: name
func (_m *Marathon) ApplicationOK(name string) (bool, error) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateApplication provides a mock function with given fields: application
func (_m *Marathon) CreateApplication(application *marathon.Application) (*marathon.Application, error) {
	ret := _m.Called(application)

	var r0 *marathon.Application
	if rf, ok := ret.Get(0).(func(*marathon.Application) *marathon.Application); ok {
		r0 = rf(application)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*marathon.Application) error); ok {
		r1 = rf(application)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteApplication provides a mock function with given fields: name
func (_m *Marathon) DeleteApplication(name string) (*marathon.DeploymentID, error) {
	ret := _m.Called(name)

	var r0 *marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(string) *marathon.DeploymentID); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplication provides a mock function with given fields: application, force
func (_m *Marathon) UpdateApplication(application *marathon.Application, force bool) (*marathon.DeploymentID, error) {
	ret := _m.Called(application, force)

	var r0 *marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(*marathon.Application, bool) *marathon.DeploymentID); ok {
		r0 = rf(application, force)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*marathon.Application, bool) error); ok {
		r1 = rf(application, force)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplicationDeployments provides a mock function with given fields: name
func (_m *Marathon) ApplicationDeployments(name string) ([]*marathon.DeploymentID, error) {
	ret := _m.Called(name)

	var r0 []*marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(string) []*marathon.DeploymentID); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScaleApplicationInstances provides a mock function with given fields: name, instances, force
func (_m *Marathon) ScaleApplicationInstances(name string, instances int, force bool) (*marathon.DeploymentID, error) {
	ret := _m.Called(name, instances, force)

	var r0 *marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(string, int, bool) *marathon.DeploymentID); ok {
		r0 = rf(name, instances, force)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, bool) error); ok {
		r1 = rf(name, instances, force)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestartApplication provides a mock function with given fields: name, force
func (_m *Marathon) RestartApplication(name string, force bool) (*marathon.DeploymentID, error) {
	ret := _m.Called(name, force)

	var r0 *marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(string, bool) *marathon.DeploymentID); ok {
		r0 = rf(name, force)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(name, force)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Applications provides a mock function with given fields: _a0
func (_m *Marathon) Applications(_a0 url.Values) (*marathon.Applications, error) {
	ret := _m.Called(_a0)

	var r0 *marathon.Applications
	if rf, ok := ret.Get(0).(func(url.Values) *marathon.Applications); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Applications)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(url.Values) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Application provides a mock function with given fields: name
func (_m *Marathon) Application(name string) (*marathon.Application, error) {
	ret := _m.Called(name)

	var r0 *marathon.Application
	if rf, ok := ret.Get(0).(func(string) *marathon.Application); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitOnApplication provides a mock function with given fields: name, timeout
func (_m *Marathon) WaitOnApplication(name string, timeout time.Duration) error {
	ret := _m.Called(name, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Duration) error); ok {
		r0 = rf(name, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tasks provides a mock function with given fields: application
func (_m *Marathon) Tasks(application string) (*marathon.Tasks, error) {
	ret := _m.Called(application)

	var r0 *marathon.Tasks
	if rf, ok := ret.Get(0).(func(string) *marathon.Tasks); ok {
		r0 = rf(application)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Tasks)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(application)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllTasks provides a mock function with given fields: opts
func (_m *Marathon) AllTasks(opts *marathon.AllTasksOpts) (*marathon.Tasks, error) {
	ret := _m.Called(opts)

	var r0 *marathon.Tasks
	if rf, ok := ret.Get(0).(func(*marathon.AllTasksOpts) *marathon.Tasks); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Tasks)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*marathon.AllTasksOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskEndpoints provides a mock function with given fields: name, port, healthCheck
func (_m *Marathon) TaskEndpoints(name string, port int, healthCheck bool) ([]string, error) {
	ret := _m.Called(name, port, healthCheck)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, int, bool) []string); ok {
		r0 = rf(name, port, healthCheck)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, bool) error); ok {
		r1 = rf(name, port, healthCheck)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KillApplicationTasks provides a mock function with given fields: applicationID, opts
func (_m *Marathon) KillApplicationTasks(applicationID string, opts *marathon.KillApplicationTasksOpts) (*marathon.Tasks, error) {
	ret := _m.Called(applicationID, opts)

	var r0 *marathon.Tasks
	if rf, ok := ret.Get(0).(func(string, *marathon.KillApplicationTasksOpts) *marathon.Tasks); ok {
		r0 = rf(applicationID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Tasks)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *marathon.KillApplicationTasksOpts) error); ok {
		r1 = rf(applicationID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KillTask provides a mock function with given fields: taskID, opts
func (_m *Marathon) KillTask(taskID string, opts *marathon.KillTaskOpts) (*marathon.Task, error) {
	ret := _m.Called(taskID, opts)

	var r0 *marathon.Task
	if rf, ok := ret.Get(0).(func(string, *marathon.KillTaskOpts) *marathon.Task); ok {
		r0 = rf(taskID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *marathon.KillTaskOpts) error); ok {
		r1 = rf(taskID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KillTasks provides a mock function with given fields: taskIDs, opts
func (_m *Marathon) KillTasks(taskIDs []string, opts *marathon.KillTaskOpts) error {
	ret := _m.Called(taskIDs, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, *marathon.KillTaskOpts) error); ok {
		r0 = rf(taskIDs, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Groups provides a mock function with given fields:
func (_m *Marathon) Groups() (*marathon.Groups, error) {
	ret := _m.Called()

	var r0 *marathon.Groups
	if rf, ok := ret.Get(0).(func() *marathon.Groups); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Groups)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Group provides a mock function with given fields: name
func (_m *Marathon) Group(name string) (*marathon.Group, error) {
	ret := _m.Called(name)

	var r0 *marathon.Group
	if rf, ok := ret.Get(0).(func(string) *marathon.Group); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroup provides a mock function with given fields: group
func (_m *Marathon) CreateGroup(group *marathon.Group) error {
	ret := _m.Called(group)

	var r0 error
	if rf, ok := ret.Get(0).(func(*marathon.Group) error); ok {
		r0 = rf(group)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGroup provides a mock function with given fields: name
func (_m *Marathon) DeleteGroup(name string) (*marathon.DeploymentID, error) {
	ret := _m.Called(name)

	var r0 *marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(string) *marathon.DeploymentID); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGroup provides a mock function with given fields: id, group
func (_m *Marathon) UpdateGroup(id string, group *marathon.Group) (*marathon.DeploymentID, error) {
	ret := _m.Called(id, group)

	var r0 *marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(string, *marathon.Group) *marathon.DeploymentID); ok {
		r0 = rf(id, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *marathon.Group) error); ok {
		r1 = rf(id, group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasGroup provides a mock function with given fields: name
func (_m *Marathon) HasGroup(name string) (bool, error) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitOnGroup provides a mock function with given fields: name, timeout
func (_m *Marathon) WaitOnGroup(name string, timeout time.Duration) error {
	ret := _m.Called(name, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Duration) error); ok {
		r0 = rf(name, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deployments provides a mock function with given fields:
func (_m *Marathon) Deployments() ([]*marathon.Deployment, error) {
	ret := _m.Called()

	var r0 []*marathon.Deployment
	if rf, ok := ret.Get(0).(func() []*marathon.Deployment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*marathon.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDeployment provides a mock function with given fields: id, force
func (_m *Marathon) DeleteDeployment(id string, force bool) (*marathon.DeploymentID, error) {
	ret := _m.Called(id, force)

	var r0 *marathon.DeploymentID
	if rf, ok := ret.Get(0).(func(string, bool) *marathon.DeploymentID); ok {
		r0 = rf(id, force)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.DeploymentID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(id, force)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasDeployment provides a mock function with given fields: id
func (_m *Marathon) HasDeployment(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitOnDeployment provides a mock function with given fields: id, timeout
func (_m *Marathon) WaitOnDeployment(id string, timeout time.Duration) error {
	ret := _m.Called(id, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Duration) error); ok {
		r0 = rf(id, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscriptions provides a mock function with given fields:
func (_m *Marathon) Subscriptions() (*marathon.Subscriptions, error) {
	ret := _m.Called()

	var r0 *marathon.Subscriptions
	if rf, ok := ret.Get(0).(func() *marathon.Subscriptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Subscriptions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddEventsListener provides a mock function with given fields: channel, filter
func (_m *Marathon) AddEventsListener(channel marathon.EventsChannel, filter int) error {
	ret := _m.Called(channel, filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(marathon.EventsChannel, int) error); ok {
		r0 = rf(channel, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveEventsListener provides a mock function with given fields: channel
func (_m *Marathon) RemoveEventsListener(channel marathon.EventsChannel) {
	_m.Called(channel)
}

// Unsubscribe provides a mock function with given fields: _a0
func (_m *Marathon) Unsubscribe(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMarathonURL provides a mock function with given fields:
func (_m *Marathon) GetMarathonURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ping provides a mock function with given fields:
func (_m *Marathon) Ping() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Info provides a mock function with given fields:
func (_m *Marathon) Info() (*marathon.Info, error) {
	ret := _m.Called()

	var r0 *marathon.Info
	if rf, ok := ret.Get(0).(func() *marathon.Info); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marathon.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Leader provides a mock function with given fields:
func (_m *Marathon) Leader() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbdicateLeader provides a mock function with given fields:
func (_m *Marathon) AbdicateLeader() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
