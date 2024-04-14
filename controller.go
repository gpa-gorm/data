package data

import (
	"encoding/json"
	//"fmt"
	"net/http"
	//"reflect"
	//"strconv"
)

// IGenericCrudController is a interface for Generic Crud Controller.
type IGenericCrudController[E GormEntity[ID], ID GormEntityId] interface {
	Save(w http.ResponseWriter, r *http.Request)
	SaveAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	UpdateAll(w http.ResponseWriter, r *http.Request)
	SoftDelete(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Count(w http.ResponseWriter, r *http.Request)
}

// GenericCrudController is a Generic Crud Controller implementation.
type GenericCrudController[E GormEntity[ID], ID GormEntityId] struct {
	S IGenericCrudService[E, ID]
	
}

// NewController creates a new IGenericCrudController.
func NewController[E GormEntity[ID], ID GormEntityId](S IGenericCrudService[E, ID]) IGenericCrudController[E, ID] {
	return &GenericCrudController[E, ID]{
		S: S,
	}
}

// Save provides save entity to database.
func (c *GenericCrudController[E, ID]) Save(w http.ResponseWriter, r *http.Request) {

	var entity E

	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		ErrorJSON(w, err)
		return
	}

	_, err = c.S.Save(&entity, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(w, entity)
}

// SaveAll provides save entities to database.
func (c *GenericCrudController[E, ID]) SaveAll(w http.ResponseWriter, r *http.Request) {

	var entities []E

	err := json.NewDecoder(r.Body).Decode(&entities)
	if err != nil {
		ErrorJSON(w, err)
		return
	}

	_, err = c.S.SaveAll(&entities, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(w, entities)
}

// FindById provides find entity by id.
func (c *GenericCrudController[E, ID]) FindById(w http.ResponseWriter, r *http.Request) {

	id_ := r.URL.Query().Get("id")
	id := ConvertToID(id_).(ID)

	entity, err := c.S.FindById(id, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(w, entity)
}

// FindAll provides find all entities.
func (c *GenericCrudController[E, ID]) FindAll(w http.ResponseWriter, r *http.Request) {

	entities, err := c.S.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(w, entities)
}

// Update provides update entity.
func (c *GenericCrudController[E, ID]) Update(w http.ResponseWriter, r *http.Request) {

	var entity E

	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		ErrorJSON(w, err)
		return
	}

	_, err = c.S.Update(&entity, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(w, entity)
}

// UpdateAll provides update entities.
func (c *GenericCrudController[E, ID]) UpdateAll(w http.ResponseWriter, r *http.Request) {

	var entities []E

	err := json.NewDecoder(r.Body).Decode(&entities)
	if err != nil {
		ErrorJSON(w, err)
		return
	}

	_, err = c.S.UpdateAll(&entities, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(w, entities)
}

// Delete provides delete entity.
func (c *GenericCrudController[E, ID]) SoftDelete(w http.ResponseWriter, r *http.Request) {

	id_ := r.URL.Query().Get("id")
	id := ConvertToID(id_).(ID)

	err := c.S.SoftDelete(id, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(w, nil)
}

// HardDelete provides hard delete entity.
func (c *GenericCrudController[E, ID]) Delete(w http.ResponseWriter, r *http.Request) {

	id_ := r.URL.Query().Get("id")
	id := ConvertToID(id_).(ID)

	err := c.S.Delete(id, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(w, nil)
}

// Count provides count entities.
func (c *GenericCrudController[E, ID]) Count(w http.ResponseWriter, r *http.Request) {

	count := c.S.Count(r.Context())
	SuccessResponse(w, count)
}
