package main

import (
	"github.com/yuta1402/todo-sample/backend/db"
	"github.com/yuta1402/todo-sample/backend/server"
)

func main() {
	db.Init()
	server.Init()
}
