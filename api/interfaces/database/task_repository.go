package database

import "task.com/task/domain"

// task repositoryはSqlHandlerを所持することを宣言
type TaskRepository struct {
    SqlHandler
}

// 全件取得
func (repo *TaskRepository) All() (task []domain.TTask, err error) {
	rows, err := repo.Query(
		"select * from tasks",
	)
	var tasks []domain.TTask

	defer rows.Close()

	for  rows.Next(){
		var task domain.TTask
		err := rows.Scan(
			&task.ID, &task.Name,
		)
		if err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, task)
	} 
	return
}

// IDを指定して取得
func (repo *TaskRepository) FindById(id int) (task domain.TTask, err error) {
	rows, err :=repo.Query(
		"select * from tasks where id = ?", id,
	)
	defer rows .Close()

	rows.Next()
	err = rows.Scan(
		&task.ID, &task.Name,
	)
	return
}