package todo

import (
	"gpadata/todoapp/internal/app/base"
)

type ITodoController[E TodoModel[ID], ID int] interface {
	//base.IBaseCrudController[E, ID]
	GetInstance() *TodoController[E, ID]
}

type TodoController[E TodoModel[ID], ID int] struct {
	base.IBaseCrudController[E, ID]
}

// Constructor for TodoController.
func NewTodoController[E TodoModel[ID], ID int](service *TodoService[E, ID]) ITodoController[E, ID] {

	// Initialize GenericCrudController with the controller for every model.
	crudController := base.NewBaseCrudController[E, ID](service)

	return &TodoController[E, ID]{IBaseCrudController: crudController}

}

func (c *TodoController[E, ID]) GetInstance() *TodoController[E, ID] {
	return c
}