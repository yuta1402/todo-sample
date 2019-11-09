package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuta1402/todo-sample/backend/service"
)

// Controller is user controlller
type UserController struct{}

// Index action: GET /users
func (uc UserController) Index(c *gin.Context) {
	var s service.UserService
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (uc UserController) Create(c *gin.Context) {
	var s service.UserService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /users/:id
func (uc UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (uc UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (uc UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
