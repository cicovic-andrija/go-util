package util

import (
	"time"
)

type Task struct {
	manager  *TaskManager
	name     string
	periodic bool
	period   time.Duration
	fn       func(context *Task) error
	lgpref   string
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Log(sev Severity, format string, v ...interface{}) {
	t.manager.log(sev, t.lgpref+format, v...)
}

type TaskManager struct {
	registered map[string]*Task
	logger     *Logger
}

func NewTaskManager(logger *Logger) *TaskManager {
	return &TaskManager{
		registered: make(map[string]*Task),
		logger:     logger,
	}
}

func (tm *TaskManager) RunAsync() {
	// TODO
}

func (tm *TaskManager) Stop() {
	// TODO
}

func (tm *TaskManager) StartTask(name string) error {
	return nil
}

func (tm *TaskManager) Register(name string, fn func(*Task) error) {
	tm.registered[name] = &Task{
		manager:  tm,
		name:     name,
		periodic: false,
		fn:       fn,
		lgpref:   "t. " + name + ": ",
	}
}

func (tm *TaskManager) RegisterPeriodic(name string, fn func(*Task) error, period time.Duration) {
	tm.registered[name] = &Task{
		manager:  tm,
		name:     name,
		periodic: true,
		period:   period,
		fn:       fn,
		lgpref:   "t. " + name + ": ",
	}
}

func (tm *TaskManager) log(sev Severity, format string, v ...interface{}) {
	if tm.logger != nil {
		tm.logger.Output(sev, 3, format, v...)
	}
}
