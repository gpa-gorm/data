package db

import (
	"gpadata/todoapp/internal/app/todo"

	"gorm.io/gorm"
)


// [E ITodoModel[ID], ID base.BaseId]

/* InitialMigration creates the tables in the database */
func  InitialMigration(db *gorm.DB) {

	err := db.AutoMigrate(&todo.TodoModel[int]{})

	if err != nil {
		panic(err)
	}
}
