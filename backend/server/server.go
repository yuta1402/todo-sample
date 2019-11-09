package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yuta1402/todo-sample/backend/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	u := r.Group("/users")
	{
		ctrl := controller.UserController{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	api := r.Group("/api/v1")

	{
		ctrl := controller.TodoController{}
		t := api.Group("/todos")
		t.GET("", ctrl.Index)
		t.GET("/:id", ctrl.Show)
		t.POST("", ctrl.Create)
		t.PUT("/:id", ctrl.Update)
		t.DELETE("/:id", ctrl.Delete)
	}

	return r
}
