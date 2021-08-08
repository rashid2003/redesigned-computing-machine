package main

import (
	"GoDo/models"
	"fmt"
	"gorm.io/gorm"
)

func allTasks(db gorm.DB)  {
	var tasks []models.Task
	db.Find(&tasks)
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i].Title)
	}
}

func addTask(db gorm.DB, title string)  {
	db.Create(&models.Task{Title: title})
	//TODO return a colorized success message
}
