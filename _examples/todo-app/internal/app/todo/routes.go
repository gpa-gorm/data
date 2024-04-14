package todo

import (
	"gpadata/todoapp/internal/app/base"
)

type ITodoRoutes[E TodoModel[ID], ID int] interface {
	//base.IGenericCrudRoutes[E, ID]
	base.IGenericCrudRoutes[E, ID]
	GetInstance() *TodoRoutes[E, ID]
}

type TodoRoutes[E TodoModel[ID], ID int] struct {
	C ITodoController[E, ID]
	base.IGenericCrudRoutes[E, ID]
}

func NewTodoRoutes[E TodoModel[ID], ID int](c *TodoController[E, ID]) ITodoRoutes[E, ID] {

	crudRoutes := base.NewRouterBuilder[E, ID](c).
		SetCount().
		SetSave().
		SetSaveAll().
		SetDelete().
		SetFindAll().
		SetFindById().
		SetDelete().
		SetUpdate().
		SetUpdateAll()

	return &TodoRoutes[E, ID]{
		IGenericCrudRoutes: crudRoutes,
		C:    c,
	//	todoController:    c,
	}
}

//func TodoRoutes(todoController ITodoController[TodoEntity]) *chi.Mux {
//	r := chi.NewRouter()
//	r.Post("/todo/create", todoController.Create)
//	return r
//}

func (r *TodoRoutes[E, ID]) GetInstance() *TodoRoutes[E, ID] {
	return r
}
