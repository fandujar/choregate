package entities

type TaskRunStatus interface {
	Status() string
}

// TaskRunPending is the status of a task run that is pending.
type TaskRunPending struct{}

func (TaskRunPending) Status() string {
	return "pending"
}

// TaskRunRunning is the status of a task run that is running.
type TaskRunRunning struct{}

func (TaskRunRunning) Status() string {
	return "running"
}

// TaskRunSucceeded is the status of a task run that has succeeded.
type TaskRunSucceeded struct{}

func (TaskRunSucceeded) Status() string {
	return "succeeded"
}

// TaskRunFailed is the status of a task run that has failed.
type TaskRunFailed struct{}

func (TaskRunFailed) Status() string {
	return "failed"
}
