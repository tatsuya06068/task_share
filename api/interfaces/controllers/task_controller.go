package controllers

import (
	"task.com/task/domain"
	"task.com/task/interfaces/database"
	"task.com/task/usecases"
	"encoding/json"
	"net/http"
	"strconv"
)

type TaskController struct {
    interactor domain.TaskInteractor
}

func NewTaskController(sqlHandler database.SqlHandler) *TaskController {
    return &TaskController {
        interactor: usecase.NewTaskInteractor{
	    	&database.TaskRepository{
	        	SqlHandler: sqlHandler
	    	}
		}
    }
}

func (controller *TaskController) TaskListView(w http.ResponseWriter, r *http.Request) (ret error) {
	tasks, err := controller.interactor.ListTask()
	ret = response(w, err, map[string]interface{}{"data": blogs})
	return ret
}

func (controller *TaskController) Task