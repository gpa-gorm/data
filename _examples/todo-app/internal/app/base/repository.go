package base

import (
	"github.com/gpa-gorm/data"

	"gorm.io/gorm"
)


type IBaseCrudRepository[E GenericModel[ID], ID int] interface {

	// Base Generic CRUD Repository Interface.
	data.IGpaGormRepository[E, ID]

}

type BaseCrudRepository[E GenericModel[ID], ID int] struct {
	IBaseCrudRepository[E, ID] // implements
}


// Constructor for BaseCrudRepository.
func NewBaseCrudRepository[E GenericModel[ID], ID int](db *gorm.DB) IBaseCrudRepository[E, ID] {

	// Initialize the GpaGormRepository with the Gorm DB for every model.
	gpaGormRepo := data.NewGpaGormRepository[E, ID](db)

	return &BaseCrudRepository[E, ID]{
		IBaseCrudRepository: gpaGormRepo,
	}
}



