package usecase

import "task.com/task/domain"

type TaskInteractor struct {
	repository TaskRepository
}

func NewTaskInteractor(task TaskRepository) domain.TaskInteractor {
	return &TaskInteractor{
		repository: task,
	}
}

/************************
        repository
************************/

type TaskRepository interface {
	All() ([]domain.TTask, error)
	FindById(int) (domain.TTask, error)
}

/**********************
   interactor methods
***********************/

func (interactor *TaskInteractor) ListTask() (tasks []domain.TTask, err error) {
	tasks, err = interactor.repository.All()
	return
}

func (interactor *TaskInteractor) DetailTask(id int) (task domain.TTask, err error) {
        task, err = interactor.repository.FindById(id)
	return
}
