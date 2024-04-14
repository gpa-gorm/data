package data

import (
	"context"
)

// IGenericCrudService is a interface for Generic Crud Service.
type IGenericCrudService[E GormEntity[ID], ID GormEntityId] interface {
	Save(entity *E, ctx context.Context) (*E, error)
	SaveAll(entity *[]E, ctx context.Context) (*[]E, error)
	FindById(id ID, ctx context.Context) (*E, error)
	FindAll(ctx context.Context) (*[]E, error)
	Update(entity *E, ctx context.Context) (*E, error)
	UpdateAll(entities *[]E, ctx context.Context) (*[]E, error)
	SoftDelete(id ID, ctx context.Context) error
	Delete(id ID, ctx context.Context) error
	Count(ctx context.Context) int64
}

// GenericCrudService is a Generic Crud Service implementation.
type GenericCrudService[E GormEntity[ID], ID GormEntityId] struct {
	R ICrudRepository[E, ID]
	//R *GpaGormRepository[E, ID]
}

// NewService creates a new IGenericCrudService.
func NewService[E GormEntity[ID], ID GormEntityId](R ICrudRepository[E, ID]) IGenericCrudService[E, ID] {
	return &GenericCrudService[E, ID]{
		R: R,
	}
}

// NewServiceImpl creates a new GenericCrudService.
func NewServiceImpl[E GormEntity[ID], ID GormEntityId](R *GpaGormRepository[E, ID]) *GenericCrudService[E, ID] {
	return &GenericCrudService[E, ID]{
		R: R,
	}
}

// Save provides save entity to database.
func (s *GenericCrudService[E, ID]) Save(entity *E, ctx context.Context) (*E, error) {
	return s.R.Save(entity, ctx)
}

// SaveAll provides save entities to database.
func (s *GenericCrudService[E, ID]) SaveAll(entity *[]E, ctx context.Context) (*[]E, error) {
	return s.R.SaveAll(entity, ctx)
}

// FindById provides find entity by id.
func (s *GenericCrudService[E, ID]) FindById(id ID, ctx context.Context) (*E, error) {
	return s.R.FindById(id, ctx)
}

// FindAll provides find all entities.
func (s *GenericCrudService[E, ID]) FindAll(ctx context.Context) (*[]E, error) {
	return s.R.FindAll(ctx)
}

// Update provides update entity.
func (s *GenericCrudService[E, ID]) Update(entity *E, ctx context.Context) (*E, error) {
	return s.R.Update(entity, ctx)
}

// UpdateAll provides update entities.
func (s *GenericCrudService[E, ID]) UpdateAll(entities *[]E, ctx context.Context) (*[]E, error) {
	return s.R.UpdateAll(entities, ctx)
}

// Delete provides delete entity.
func (s *GenericCrudService[E, ID]) Delete(id ID, ctx context.Context) error {
	return s.R.Delete(id, ctx)
}

// HardDelete provides hard delete entity.
func (s *GenericCrudService[E, ID]) SoftDelete(id ID, ctx context.Context) error {
	return s.R.SoftDelete(id, ctx)
}

// Count provides count entities.
func (s *GenericCrudService[E, ID]) Count(ctx context.Context) int64 {
	return s.R.Count(ctx)
}
