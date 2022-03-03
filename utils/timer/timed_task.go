package timer

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Timer interface {
	AddTaskByFunc(taskName string, spec string, task func()) (cron.EntryID, error)
	AddTaskByJob(taskName string, spec string, job interface{ Run() }) (cron.EntryID, error)
	FindCron(taskName string) (*cron.Cron, bool)
	StartTask(taskName string)
}

// timer 定时任务管理
type timer struct {
	taskList map[string]*cron.Cron
	sync.Mutex
}

//@function: AddTaskByFunc
//@description: 通过函数的方法添加任务
//@param: taskName string
//@param: spec string
//@param: task func()
//@return cron.EntryID, error

func (t *timer) AddTaskByFunc(taskName string, spec string, task func()) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New()
	}
	id, err := t.taskList[taskName].AddFunc(spec, task)
	t.taskList[taskName].Start()
	return id, err
}

//@function: AddTaskByJob
//@description: 通过接口的方法添加任务
//@param: taskName string
//@param: spec string
//@param: job interface{ Run() }
//@return cron.EntryID, error

func (t *timer) AddTaskByJob(taskName string, spec string, job interface{ Run() }) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New()
	}
	id, err := t.taskList[taskName].AddJob(spec, job)
	t.taskList[taskName].Start()
	return id, err
}

//@function: FindCron
//@description: 获取对应taskName的cron 可能会为空
//@param: taskName string
//@return *cron.Cron, bool

func (t *timer) FindCron(taskName string) (*cron.Cron, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.taskList[taskName]
	return v, ok
}

//@function: StartTask
//@description: 开始任务
//@param: taskName string

func (t *timer) StartTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Start()
	}
}

func NewTimerTask() Timer {
	return &timer{
		taskList: make(map[string]*cron.Cron),
	}
}
