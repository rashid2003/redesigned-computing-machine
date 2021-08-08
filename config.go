package main

import (
	"GoDo/models"
	"flag"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"os/user"
)



func initSqliteDB() string {
	detectOs := detectOs()
	var dbFile string
	if detectOs == "windows" {
		currentUser, err := user.Current()
		if err != nil {
			panic(err)
		}
		dbFile := currentUser.HomeDir + "\\AppData\\Local\\GoDo\\main.db"
		dbFolder := currentUser.HomeDir + "\\AppData\\Local\\GoDo\\"
		if _, err := os.Stat(dbFile); os.IsNotExist(err) {
			err = os.Mkdir(dbFolder, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	//TODO Add func for Linux and Mac
	return dbFile
}

func setupDB() *gorm.DB {
	dbFile := initSqliteDB()
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Task{}, &models.Category{})
	if err != nil {
		panic(err)
	}

	return db
}

func setupFlag() string {
	taskTitle := flag.String("add", "New Task", "Task Title")
	_ = flag.String("show", "Show Tasks", "Show all tasks")

	flag.Parse()

	return *taskTitle
}
