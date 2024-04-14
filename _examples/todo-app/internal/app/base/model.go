package base

import (
	//"github.com/gpa-gorm/data"
)

// GenericModel is a type for passing Generic Model on base module.
type GenericModel[ID int] any


// BaseModel is a struct that contains the following fields: ID, CreatedAt, UpdatedAt, DeletedAt, and Deleted.
// ID is generic type for ID. We assign int type for ID for this example.
type BaseModel[ID int] struct {

	// Fields that are common to all tables
	ID        int    	`gorm:"primaryKey;autoIncrement;not null"`
	IsActive  bool   	`gorm:"default:true"`
	DeletedAt int    	`gorm:"index"`
	Creator   string 	`gorm:"size:50"`
	Updater   string 	`gorm:"size:50"`
	isDeleted bool   	`gorm:"-"`
}

// GetId is a method that returns the ID of the BaseModel.
func (m BaseModel[ID]) GetId() int {
	return int(m.ID)
}

// TableName is a method that sets the table name for the BaseModel.
//func (m BaseModel[ID]) TableName() string {
	//panic("BaseModel should not be used directly, it should be embedded in other models.")
//}