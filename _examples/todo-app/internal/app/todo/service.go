package todo

import (
	"gpadata/todoapp/internal/app/base"
)

// ITodoService is the interface for TodoService.
type ITodoService[E TodoModel[ID], ID int] interface {
	//base.IBaseCrudService[E, ID]
	GetInstance() *TodoService[E, ID]
}

// TodoService is the struct for TodoService.
type TodoService[E TodoModel[ID], ID int] struct {
	base.IBaseCrudService[E, ID]
}

// Constructor for TodoService.
//[ITodoModel[base.BaseId], base.BaseId]
func NewTodoService[E TodoModel[ID], ID int](repo *TodoRepository[E, ID]) ITodoService[E, ID] {

	//infer := repo.(*base.IBaseCrudController[E, ID])
//func NewTodoService(baseService ITodoService[ITodoModel[base.BaseId], base.BaseId]) ITodoService[ITodoModel[base.BaseId], base.BaseId] {

	// Initialize GenericCrudService with the service for every model.
	crudService := base.NewBaseCrudService[E, ID](repo)

	
	return &TodoService[E, ID]{
	//return &TodoService[ITodoModel[base.BaseId], base.BaseId]{
		IBaseCrudService: crudService,
	}
}

func (s *TodoService[E, ID]) GetInstance() *TodoService[E, ID] {
	return s
}


