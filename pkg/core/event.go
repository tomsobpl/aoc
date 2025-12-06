package core

type AocEvent interface {
	AddTask(task AocTask)
	Tasks() []AocTask
}

type aocEvent struct {
	name  string
	tasks []AocTask
}

func (a *aocEvent) AddTask(task AocTask) {
	a.tasks = append(a.tasks, task)
}

func (a *aocEvent) Tasks() []AocTask {
	return a.tasks
}

func NewAocEvent(name string) AocEvent {
	return &aocEvent{name: name}
}
