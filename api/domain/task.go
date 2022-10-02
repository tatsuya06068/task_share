package domain

type TTask struct {
    ID      int    `json:":task_id"`
    Name   string `json:"name"`
}

type TaskInteractor interface {
    ListTask() ([]TTask, error)
    DetailTask(int) (TTask, error)
}