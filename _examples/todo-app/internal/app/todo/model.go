package todo

import (
	"gpadata/todoapp/internal/app/base"
	"gorm.io/gorm"
)

// TodoModel is the struct for TodoModel.
type TodoModel[ID int] struct {

	// gorm.Model is a struct that contains the following fields: ID, CreatedAt, UpdatedAt, DeletedAt, and Deleted.
	gorm.Model
	base.BaseModel[int] // implements
	Name     string  `gorm:"size:50"`
	Price    float64 `gorm:"type:decimal(10,2)"`
	Quantity int     `gorm:"type:int"`
}

// GetId returns the ID of the TodoModel.
func (t TodoModel[ID]) GetId() ID {
	return ID(t.BaseModel.ID)
}

// TableName sets the table name for the TodoModel.
func (t TodoModel[ID]) TableName() string {
	return "todos"
}