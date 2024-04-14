package base

import (
	"github.com/gpa-gorm/data"
)

type IBaseCrudController[E GenericModel[ID], ID int] interface {

	// Base Generic CRUD Controller Interface.
	data.IGenericCrudController[E, ID]
}

type BaseCrudController[E GenericModel[ID], ID int] struct {
	IBaseCrudController[E, ID] // implements
}

// Constructor for BaseCrudController.
func NewBaseCrudController[E GenericModel[ID], ID int](service IBaseCrudService[E, ID]) IBaseCrudController[E, ID] {

	// Initialize GenericCrudController with the controller for every model.
	crudService := data.NewController[E, ID](service)

	return &BaseCrudController[E, ID]{IBaseCrudController: crudService}
}