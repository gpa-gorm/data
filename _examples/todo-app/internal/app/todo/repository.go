package todo

import (
	"gpadata/todoapp/internal/app/base"

	"gorm.io/gorm"
)

// ITodoRepository is the interface for TodoRepository.
type ITodoRepository[E TodoModel[ID], ID int] interface {
	//base.IBaseCrudRepository[E, ID]
	GetInstance() *TodoRepository[E, ID]
}

// TodoRepository is the struct for TodoRepository.
type TodoRepository[E TodoModel[ID], ID int] struct {
	base.IBaseCrudRepository[E, ID]
}


// Don't Construct model like this. Because compiler can infer it:
func NewTodoRepository[E TodoModel[ID], ID int](db *gorm.DB) ITodoRepository[E, ID] {
// Instead, construct the model like at below:
// Constructor for TodoRepository.
//func NewTodoRepository(db *gorm.DB) ITodoRepository[ITodoModel[base.BaseId], base.BaseId] {

	// Initialize GenericCrudRepository with the model for every model w/type inference.
	//crudRepository := base.NewBaseCrudRepository[ITodoModel[base.BaseId], base.BaseId](db)
	crudRepository := base.NewBaseCrudRepository[E, ID](db)

	return &TodoRepository[E, ID]{IBaseCrudRepository: crudRepository,}

}

// GetInstance returns the instance of TodoRepository.
func (r *TodoRepository[E, ID]) GetInstance() *TodoRepository[E, ID] {
	return r
}


