package main

import (
	"fmt"
	"go-todo-app/db"
	"go-todo-app/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migration.")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}