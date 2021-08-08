package main

import (
	"flag"
)

func main() {

	db := setupDB()
	taskTitle := setupFlag()

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "show" {
			allTasks(*db)
		}
		if f.Name == "add" {
			addTask(*db, taskTitle)
		}
	})

}




