package base

import (
	"github.com/go-chi/chi/v5"
)

// Generic Crud Router implementation with go-chi.
type IGenericCrudRoutes[E GenericModel[ID], ID int] interface {
	SetSave() IGenericCrudRoutes[E, ID]
	SetSaveAll() IGenericCrudRoutes[E, ID]
	SetFindById() IGenericCrudRoutes[E, ID]
	SetFindAll() IGenericCrudRoutes[E, ID]
	SetUpdate() IGenericCrudRoutes[E, ID]
	SetUpdateAll() IGenericCrudRoutes[E, ID]
	SetDelete() IGenericCrudRoutes[E, ID]
	SetSoftDelete() IGenericCrudRoutes[E, ID]
	SetCount() IGenericCrudRoutes[E, ID]
	GetRouter() *chi.Mux
}

type GenericCrudRoutes[E GenericModel[ID], ID int] struct {
	C     	IBaseCrudController[E, ID]
	R 		*chi.Mux
}

func NewRouterBuilder[E GenericModel[ID], ID int](c IBaseCrudController[E, ID]) IGenericCrudRoutes[E, ID] {
	return &GenericCrudRoutes[E, ID]{
		C:      c,
		R: chi.NewRouter(),
	}
}

// Create new entity.
func (r *GenericCrudRoutes[E, ID]) SetSave() IGenericCrudRoutes[E, ID] {
	r.R.Post("/create", r.C.Save)
	return r

}

// Create new entities.
func (r *GenericCrudRoutes[E, ID]) SetSaveAll() IGenericCrudRoutes[E, ID] {
	r.R.Post("/create/all", r.C.SaveAll)
	return r
}

// Get entity by id.
func (r *GenericCrudRoutes[E, ID]) SetFindById() IGenericCrudRoutes[E, ID] {
	r.R.Get("/get/{id}", r.C.FindById)
	return r
}

func (r *GenericCrudRoutes[E, ID]) SetFindAll() IGenericCrudRoutes[E, ID] {
	r.R.Get("/get/all", r.C.FindAll)
	return r
}

func (r *GenericCrudRoutes[E, ID]) SetUpdate() IGenericCrudRoutes[E, ID] {
	r.R.Put("/update/{id}", r.C.Update)
	return r
}

func (r *GenericCrudRoutes[E, ID]) SetUpdateAll() IGenericCrudRoutes[E, ID] {
	r.R.Put("/update/all", r.C.UpdateAll)
	return r
}

func (r *GenericCrudRoutes[E, ID]) SetDelete() IGenericCrudRoutes[E, ID] {
	r.R.Delete("/delete/{id}", r.C.Delete)
	return r
}

func (r *GenericCrudRoutes[E, ID]) SetSoftDelete() IGenericCrudRoutes[E, ID] {

	r.R.Delete("/hard-delete/{id}", r.C.SoftDelete)
	return r
}

func (r *GenericCrudRoutes[E, ID]) SetCount() IGenericCrudRoutes[E, ID] {
	r.R.Get("/count", r.C.Count)
	return r
}

func (r *GenericCrudRoutes[E, ID]) GetRouter() *chi.Mux {
	return r.R
}
