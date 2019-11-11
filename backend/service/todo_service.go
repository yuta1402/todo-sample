package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yuta1402/todo-sample/backend/db"
	"github.com/yuta1402/todo-sample/backend/entity"
)

type TodoService struct{}

type Todo entity.Todo

func (s TodoService) GetAll() ([]Todo, error) {
	db := db.GetDB()
	var t []Todo

	if err := db.Find(&t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (s TodoService) CreateModel(c *gin.Context) (Todo, error) {
	db := db.GetDB()
	var t Todo

	if err := c.BindJSON(&t); err != nil {
		return t, err
	}

	if err := db.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (s TodoService) GetByID(id string) (Todo, error) {
	db := db.GetDB()
	var t Todo

	if err := db.Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (s TodoService) UpdateByID(id string, c *gin.Context) (Todo, error) {
	db := db.GetDB()
	var t Todo

	if err := db.Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}

	if err := c.BindJSON(&t); err != nil {
		return t, err
	}

	db.Save(&t)

	return t, nil
}

func (s TodoService) DeleteByID(id string) error {
	db := db.GetDB()
	var t Todo

	if err := db.Where("id = ?", id).Delete(&t).Error; err != nil {
		return err
	}

	return nil
}

func (s TodoService) ToggleByID(id string) (Todo, error) {
	db := db.GetDB()
	var t Todo

	if err := db.Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}

	t.Completed = !t.Completed
	db.Save(&t)

	return t, nil
}
