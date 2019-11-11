package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuta1402/todo-sample/backend/service"
)

type TodoController struct{}

// Index action: GET /todos
func (tc TodoController) Index(c *gin.Context) {
	var s service.TodoService
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /todos
func (tc TodoController) Create(c *gin.Context) {
	var s service.TodoService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /todos/:id
func (tc TodoController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.TodoService
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /todos/:id
func (tc TodoController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.TodoService
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /todos/:id
func (tc TodoController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.TodoService

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

// Toggle action: POST /todos/:id/toggle
func (tc TodoController) Toggle(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.TodoService

	t, err := s.ToggleByID(id)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, t)
	}
}
