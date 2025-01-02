package task

/*
* @author: Chen Chiheng
* @date: 2025/1/2 20:54:23
* @description:
**/

type Interface interface {
	Run() error
	RunAndWait() error
	Stop()
	Pipeline(Interface) Interface
}

type Task struct {
}

type JobInterface interface {
}

func (task *Task) RegisterJob() {

}
