package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/yuta1402/todo-sample/backend/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run(":3000")
}

func router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile("public", false)))

	api := r.Group("/api/v1")

	{
		ctrl := controller.TodoController{}
		t := api.Group("/todos")
		t.GET("", ctrl.Index)
		t.GET("/:id", ctrl.Show)
		t.POST("", ctrl.Create)
		t.PUT("/:id", ctrl.Update)
		t.DELETE("/:id", ctrl.Delete)
		t.POST("/:id/toggle", ctrl.Toggle)
	}

	return r
}
