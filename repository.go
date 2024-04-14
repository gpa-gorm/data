// Package data provides generic repository, generic service and generic controller interface
// and their implementation with Gorm.

package data

import (
	"context"

	"gorm.io/gorm"
)

// IRepository is base generic repository interface.
type IRepository[E Entity[ID], ID Id] interface {
}

// ICrud IRepository is base generic crud repository interface.
type ICrudRepository[E Entity[ID], ID Id] interface {
	IRepository[E, ID] // implements

	// <CRUD> operations
	Save(entity *E, ctx context.Context) (*E, error)
	SaveAll(entity *[]E, ctx context.Context) (*[]E, error)
	FindById(id ID, ctx context.Context) (*E, error)
	FindAll(ctx context.Context) (*[]E, error)
	Update(entity *E, ctx context.Context) (*E, error)
	UpdateAll(entities *[]E, ctx context.Context) (*[]E, error)
	SoftDelete(id ID, ctx context.Context) error
	Delete(id ID, ctx context.Context) error

	// Aggregate operations
	Count(ctx context.Context) int64
}

// IGpaGormRepository is a Generic Gorm Crud IRepository interface.
type IGpaGormRepository[E GormEntity[ID], ID GormEntityId] interface {
	ICrudRepository[E, ID]

	// Gorm specific operations

	//Begin(opts ...*sql.TxOptions) GormRepository[E, ID]
	//Rollback() error
	//Commit() error
}

// GpaGormRepository is a Generic Gorm Crud IRepository implementation.
type GpaGormRepository[E GormEntity[ID], ID GormEntityId] struct {
	DB *gorm.DB
}

// NewGpaGormRepository creates a new IGpaGormRepository.
func NewGpaGormRepository[E GormEntity[ID], ID GormEntityId](DB *gorm.DB) IGpaGormRepository[E, ID] {
	return &GpaGormRepository[E, ID]{
		DB: DB,
	}
}

// NewGpaGormRepositoryImpl creates a new GpaGormRepository.
func NewGpaGormRepositoryImpl[E GormEntity[ID], ID GormEntityId](DB *gorm.DB) *GpaGormRepository[E, ID] {
	return &GpaGormRepository[E, ID]{
		DB: DB,
	}
}


// Save provides save entity to database.
func (r *GpaGormRepository[E, ID]) Save(entity *E, ctx context.Context) (*E, error) {

	err := r.DB.WithContext(ctx).Save(&entity).Error

	if err != nil {
		return nil, err
	}

	return entity, nil

}

// SaveAll provides save entities to database.
func (r *GpaGormRepository[E, ID]) SaveAll(entity *[]E, ctx context.Context) (*[]E, error) {

	err := r.DB.WithContext(ctx).Save(&entity).Error

	if err != nil {
		return nil, err
	}

	return entity, nil
}

// FindById provides find entity by id.
func (r *GpaGormRepository[E, ID]) FindById(id ID, ctx context.Context) (*E, error) {

	var entity E

	err := r.DB.WithContext(ctx).First(&entity, id).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// FindAll provides find all entities.
func (r *GpaGormRepository[E, ID]) FindAll(ctx context.Context) (*[]E, error) {

	var entities []E

	err := r.DB.WithContext(ctx).Find(&entities).Where("is_deleted = ?", false).Error
	if err != nil {
		return nil, err
	}

	return &entities, nil
}

// Update provides update entity.
func (r *GpaGormRepository[E, ID]) Update(entity *E, ctx context.Context) (*E, error) {

	err := r.DB.WithContext(ctx).Save(&entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil

}

// UpdateAll provides update entities.
func (r *GpaGormRepository[E, ID]) UpdateAll(entities *[]E, ctx context.Context) (*[]E, error) {

	err := r.DB.WithContext(ctx).Save(&entities).Error

	if err != nil {
		return nil, err
	}

	return entities, nil
}

// Delete provides delete entity from database.
func (r *GpaGormRepository[E, ID]) SoftDelete(id ID, ctx context.Context) error {

	entity, err := r.FindById(id, ctx)
	if err != nil {
		return err
	}

	return r.DB.WithContext(ctx).First(&entity).Where("id = ?", id).UpdateColumn("is_deleted", true).Error
}

// HardDelete provides hard delete entity from database.
func (r *GpaGormRepository[E, ID]) Delete(id ID, ctx context.Context) error {

	entity, err := r.FindById(id, ctx)
	if err != nil {
		return err
	}

	return r.DB.WithContext(ctx).Delete(&entity).Error
}

// Count provides count of entities.
func (r *GpaGormRepository[E, ID]) Count(ctx context.Context) int64 {

	var entity E
	var count int64

	err := r.DB.WithContext(ctx).Model(&entity).Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}
