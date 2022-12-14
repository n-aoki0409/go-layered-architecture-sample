package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/n-aoki0409/go-layered-architecture-sample/domain/model"
	"github.com/n-aoki0409/go-layered-architecture-sample/domain/repository"
)

type TaskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

func (tr *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) FindByID(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	if err := tr.Conn.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Model(&task).Update(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Delete(task *model.Task) error {
	if err := tr.Conn.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
