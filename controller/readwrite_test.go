package controller

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/hashicorp/consul-terraform-sync/config"
	"github.com/hashicorp/consul-terraform-sync/driver"
	"github.com/hashicorp/consul-terraform-sync/logging"
	mocksD "github.com/hashicorp/consul-terraform-sync/mocks/driver"
	mocks "github.com/hashicorp/consul-terraform-sync/mocks/templates"
	mocksTmpl "github.com/hashicorp/consul-terraform-sync/mocks/templates"
	"github.com/hashicorp/consul-terraform-sync/state"
	"github.com/hashicorp/consul-terraform-sync/templates"
	"github.com/hashicorp/consul-terraform-sync/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ReadWrite_Run(t *testing.T) {
	t.Parallel()

	w := new(mocksTmpl.Watcher)
	w.On("Watch", mock.Anything, mock.Anything).Return(nil)

	port := testutils.FreePort(t)

	ctl := ReadWrite{}

	tm := newTestTasksManager()
	tm.watcher = w
	tm.state = state.NewInMemoryStore(&config.Config{
		Port: config.Int(port),
	})
	ctl.tasksManager = &tm

	t.Run("cancel exits successfully", func(t *testing.T) {
		errCh := make(chan error)
		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			if err := ctl.Run(ctx); err != nil {
				errCh <- err
			}
		}()
		cancel()

		select {
		case err := <-errCh:
			// Confirm that exit is due to context cancel
			assert.Equal(t, err, context.Canceled)
		case <-time.After(time.Second * 15):
			t.Fatal("Run did not exit properly from cancelling context")
		}
	})

	t.Run("error exits successfully", func(t *testing.T) {
		errCh := make(chan error)
		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			if err := ctl.Run(ctx); err != nil {
				errCh <- err
			}
		}()

		// Re-run controller to create an "address already in use" error when
		// trying to serve api at same port
		go func() {
			if err := ctl.Run(ctx); err != nil {
				errCh <- err
			}
		}()
		defer cancel()

		select {
		case err := <-errCh:
			// Confirm error was received and successfully exit
			assert.Contains(t, err.Error(), "address already in use")
		case <-time.After(time.Second * 5):
			t.Fatal("Run did not error and exit properly")
		}
	})
}

func Test_ReadWrite_Once_Terraform(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("test error")

	testCases := []struct {
		name           string
		numTasks       int
		setupNewDriver func(*driver.Task) driver.Driver
		expectErr      bool
	}{
		{
			"consecutive one task",
			1,
			func(task *driver.Task) driver.Driver {
				return onceMockDriver(task, nil)
			},
			false,
		},
		{
			"consecutive multiple tasks",
			10,
			func(task *driver.Task) driver.Driver {
				return onceMockDriver(task, nil)
			},
			false,
		},
		{
			"consecutive error",
			5,
			func(task *driver.Task) driver.Driver {
				if task.Name() == "task_03" {
					// Mock an error during apply for a task
					return onceMockDriver(task, expectedErr)
				}
				return onceMockDriver(task, nil)
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			driverConf := &config.DriverConfig{
				Terraform: &config.TerraformConfig{},
			}

			err := testOnce(t, tc.numTasks, driverConf, tc.setupNewDriver)
			if tc.expectErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), expectedErr.Error(),
					"unexpected error in Once")
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_ReadWrite_Once_WatchDep_errors_Terraform(t *testing.T) {
	// Mock the situation where WatchDep WaitCh errors which can cause
	// driver.RenderTemplate() to always returns false. Confirm on WatchDep
	// error, that creating/running tasks does not hang and exits.
	t.Parallel()

	driverConf := &config.DriverConfig{
		Terraform: &config.TerraformConfig{},
	}

	testOnceWatchDepErrors(t, driverConf)
}

func Test_ReadWrite_Once_then_Run_Terraform(t *testing.T) {
	// Tests Run behaves as expected with triggers after once completes
	t.Parallel()

	driverConf := &config.DriverConfig{
		Terraform: &config.TerraformConfig{},
	}

	testOnceThenRun(t, driverConf)
}

func Test_ReadWrite_onceConsecutive_context_canceled(t *testing.T) {
	// - Controller will try to create and run 5 tasks
	// - Mock a task to take 2 seconds to create and run
	// - Cancel context after 1 second. Confirm only 1 task created and run
	t.Parallel()

	conf := multipleTaskConfig(5)
	ss := state.NewInMemoryStore(conf)

	rw := ReadWrite{
		logger: logging.NewNullLogger(),
		state:  ss,
	}

	// Set up tasks manager
	tm := newTestTasksManager()
	tm.state = ss
	rw.tasksManager = &tm

	// Set up baseController
	tm.baseController.initConf = conf
	tm.baseController.newDriver = func(c *config.Config, task *driver.Task, w templates.Watcher) (driver.Driver, error) {
		d := new(mocksD.Driver)
		d.On("Task").Return(task).Times(4)
		d.On("TemplateIDs").Return(nil)
		d.On("RenderTemplate", mock.Anything).Return(true, nil).Once()
		d.On("InitTask", mock.Anything, mock.Anything).Return(nil).Once()
		d.On("ApplyTask", mock.Anything).Return(nil).Once()
		d.On("OverrideNotifier").Return().Once()
		// Last driver call takes 2 seconds
		d.On("SetBufferPeriod").Return().After(2 * time.Second).Once()
		return d, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)
	go func() {
		err := rw.onceConsecutive(ctx)
		if err != nil {
			errCh <- err
		}
	}()
	// Cancel context after 1 second (during first task)
	time.Sleep(time.Second)
	cancel()

	select {
	case err := <-errCh:
		assert.Equal(t, err, context.Canceled)
	case <-time.After(time.Second * 5):
		t.Fatal("onceConsecutive did not exit properly from cancelling context")
	}

	assert.Equal(t, 1, tm.drivers.Len(), "only one driver should have been created")
	for _, d := range tm.drivers.Map() {
		d.(*mocksD.Driver).AssertExpectations(t)
	}
}

func testOnce(t *testing.T, numTasks int, driverConf *config.DriverConfig,
	setupNewDriver func(*driver.Task) driver.Driver) error {

	conf := multipleTaskConfig(numTasks)
	conf.Driver = driverConf
	ss := state.NewInMemoryStore(conf)

	rw := ReadWrite{
		logger: logging.NewNullLogger(),
		state:  ss,
	}

	// Set up tasks manager
	tm := newTestTasksManager()
	tm.state = ss
	tm.deleteCh = make(chan string, 1)
	rw.tasksManager = &tm

	// Mock watcher
	errCh := make(chan error)
	var errChRc <-chan error = errCh
	go func() { errCh <- nil }()
	w := new(mocksTmpl.Watcher)
	w.On("WaitCh", mock.Anything).Return(errChRc)
	w.On("Size").Return(numTasks)
	tm.watcher = w

	// Set up baseController
	tm.baseController.initConf = conf
	tm.baseController.newDriver = func(c *config.Config, task *driver.Task, w templates.Watcher) (driver.Driver, error) {
		return setupNewDriver(task), nil
	}

	err := rw.Once(context.Background())

	w.AssertExpectations(t)
	for _, d := range tm.drivers.Map() {
		d.(*mocksD.Driver).AssertExpectations(t)
	}

	return err
}

func testOnceWatchDepErrors(t *testing.T, driverConf *config.DriverConfig) {
	conf := singleTaskConfig()

	if driverConf != nil {
		conf.Driver = driverConf
	}

	ss := state.NewInMemoryStore(conf)

	rw := ReadWrite{
		logger: logging.NewNullLogger(),
		state:  ss,
	}

	// Set up tasks manager
	tm := newTestTasksManager()
	tm.state = ss
	rw.tasksManager = &tm

	// Mock watcher
	expectedErr := errors.New("error!")
	waitErrCh := make(chan error)
	var waitErrChRc <-chan error = waitErrCh
	go func() { waitErrCh <- expectedErr }()
	w := new(mocksTmpl.Watcher)
	w.On("WaitCh", mock.Anything).Return(waitErrChRc)
	tm.watcher = w

	// Set up baseController
	tm.baseController.initConf = conf
	tm.baseController.newDriver = func(c *config.Config, task *driver.Task, w templates.Watcher) (driver.Driver, error) {
		d := new(mocksD.Driver)
		d.On("InitTask", mock.Anything, mock.Anything).Return(nil)
		// Always return false on render template to mock what happens when
		// WaitCh returns an error
		d.On("RenderTemplate", mock.Anything).Return(false, nil)

		return d, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := make(chan error)
	go func() {
		err := rw.Once(ctx)
		if err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		require.Error(t, err)
		assert.Equal(t, err, expectedErr)
	case <-time.After(time.Second * 5):
		t.Fatal("Once did not exit properly after WatcherDep errored")
	}

	w.AssertExpectations(t)
	for _, d := range tm.drivers.Map() {
		d.(*mocksD.Driver).AssertExpectations(t)
	}
}

func testOnceThenRun(t *testing.T, driverConf *config.DriverConfig) {
	port := testutils.FreePort(t)
	conf := singleTaskConfig()
	conf.Driver = driverConf
	conf.Port = config.Int(port)
	conf.Finalize()

	st := state.NewInMemoryStore(conf)

	rw := ReadWrite{
		logger: logging.NewNullLogger(),
		state:  st,
	}

	// Setup taskmanager
	tm := newTestTasksManager()
	tm.watcherCh = make(chan string, 5)
	tm.state = st
	rw.tasksManager = &tm

	// Mock driver
	tm.baseController.newDriver = func(c *config.Config, task *driver.Task, w templates.Watcher) (driver.Driver, error) {
		d := new(mocksD.Driver)
		d.On("Task").Return(task)
		d.On("TemplateIDs").Return([]string{"{{tmpl}}"})
		d.On("RenderTemplate", mock.Anything).Return(true, nil)
		d.On("InitTask", mock.Anything, mock.Anything).Return(nil).Once()
		d.On("ApplyTask", mock.Anything).Return(nil)
		d.On("OverrideNotifier").Return().Once()
		d.On("SetBufferPeriod").Return().Once()
		return d, nil
	}

	// Setup variables for testing
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	completedTasksCh := tm.EnableTestMode()

	// Mock watcher
	w := new(mocks.Watcher)
	waitErrCh := make(chan error)
	var waitErrChRc <-chan error = waitErrCh
	go func() { errCh <- nil }()
	w.On("WaitCh", mock.Anything).Return(waitErrChRc).Once()
	w.On("Size").Return(5)
	w.On("Watch", ctx, tm.watcherCh).Return(nil)
	tm.watcher = w

	go func() {
		err := rw.Once(ctx)
		if err != nil {
			errCh <- err
			return
		}

		err = rw.Run(ctx)
		if err != nil {
			errCh <- err
		}
	}()

	// Emulate triggers to evaluate task completion
	for i := 0; i < 5; i++ {
		tm.watcherCh <- "{{tmpl}}"
		select {
		case taskName := <-completedTasksCh:
			assert.Equal(t, "task", taskName)

		case err := <-errCh:
			require.NoError(t, err)
		case <-ctx.Done():
			assert.NoError(t, ctx.Err(), "Context should not timeout. Once and Run usage of Watcher does not match the expected triggers.")
		}
	}

	w.AssertExpectations(t)
	for _, d := range tm.drivers.Map() {
		d.(*mocksD.Driver).AssertExpectations(t)
	}
}

// onceMockDriver mocks the driver with the methods needed for once-mode
func onceMockDriver(task *driver.Task, applyTaskErr error) driver.Driver {
	d := new(mocksD.Driver)
	d.On("Task").Return(task).Times(4)
	d.On("TemplateIDs").Return(nil)
	d.On("RenderTemplate", mock.Anything).Return(false, nil).Once()
	d.On("RenderTemplate", mock.Anything).Return(true, nil).Once()
	d.On("InitTask", mock.Anything, mock.Anything).Return(nil).Once()
	d.On("ApplyTask", mock.Anything).Return(applyTaskErr).Once()
	d.On("OverrideNotifier").Return().Once()
	d.On("SetBufferPeriod").Return().Once()
	return d
}
