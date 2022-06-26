package main

import (
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct{
	gorm.Model
	Text       string
	Status     string
}

func dbInit() {
  db,err := gorm.Open("mysql","mysql.test")
	if err != nil{
		panic("データベースが開けません。")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}


func dbInsert(text string,status string){
	db,err := gorm.Open("mysql","mysql.test")
	if err != nil{
		panic("データベースが開けません。")
	}
	db.Create(&Todo{Text: text,Status: status})
	defer db.Close();
}

func dbUpdate(id int,text string,status string){
	db,err := gorm.Open("mysql","mysql.test")
	if err != nil{
		panic("データベースが開けません。")
	}
	var todo Todo
	db.First(&todo,id)
	if text != ""{
		todo.Text = text
	}
	if status != ""{
		todo.Status = status
	}
	db.Save(&todo)
	db.Close()
}
