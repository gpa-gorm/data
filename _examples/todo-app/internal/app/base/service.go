package base

import (
	"github.com/gpa-gorm/data"
)

type IBaseCrudService[E GenericModel[ID], ID int] interface {

	// Base Generic CRUD Service Interface.
	data.IGenericCrudService[E, ID]
}

type BaseCrudService[E GenericModel[ID], ID int] struct {
	data.IGenericCrudService[E, ID] // implements
}

// Constructor for BaseCrudRepository.
func NewBaseCrudService[E GenericModel[ID], ID int](r IBaseCrudRepository[E,ID]) IBaseCrudService[E, ID] {

	// Initialize the generic CRUD service with the service for every model.
	crudService := data.NewService[E, ID](r)

	return &BaseCrudService[E, ID]{
		IGenericCrudService: crudService,
	}
}