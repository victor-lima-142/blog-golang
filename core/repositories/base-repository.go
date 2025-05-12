package repositories

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func checkConds(conds ...interface{}) bool {
	if conds == nil {
		return false
	}

	if len(conds) == 0 {
		return false
	}

	if len(conds) == 1 {
		firstElement := conds[0]

		val := reflect.ValueOf(firstElement)
		strVal := val.String()
		if strVal == "<[]interface {} Value>" || strVal == "[[]]" {
			return false
		}
	}

	return true
}

type Repository[T any] interface {
	// Create inserts a new entity into the repository.
	// Parameters:
	// - entity: the entity to insert.
	// Returns an error if any occurs during the insertion.
	Create(entity *T) error

	// FindByID retrieves an entity by its unique identifier.
	// Parameters:
	// - id: the unique identifier of the entity to retrieve.
	// Returns a pointer to the entity if found, and an error if any occurs during the retrieval.
	FindByID(id int64) (*T, error)

	// FindAll retrieves all entities from the repository, optionally filtered by the provided conditions.
	// If no conditions are provided, all entities are retrieved.
	// Parameters:
	// - conds: optional conditions to filter the entities.
	// Returns a slice of entities and an error if any occurs during the retrieval.
	FindAll(conds ...interface{}) ([]T, error)

	// FindPaginated retrieves a paginated list of entities based on the provided conditions.
	// Parameters:
	// - page: the page number to retrieve.
	// - pageSize: the number of entities per page.
	// - conds: optional conditions to filter the entities.
	// Returns a slice of entities and an error if any occurs during the retrieval.
	FindPaginated(page int, pageSize int, conds ...interface{}) ([]T, error)

	// FindOneBy retrieves a single entity based on the provided conditions.
	// Parameters:
	// - conds: optional conditions to filter the entity.
	// Returns a pointer to the entity if found, and an error if any occurs during the retrieval.
	FindOneBy(conds ...interface{}) (*T, error)

	// Delete removes an entity from the repository.
	// Parameters:
	// - id: the unique identifier of the entity to remove.
	// Returns an error if any occurs during the deletion.
	Delete(id int64) error

	// Update updates an entity in the repository.
	// Parameters:
	// - entity: the entity to update.
	// Returns an error if any occurs during the update.
	Update(entity *T) error

	// Count returns the number of entities in the repository that match the provided conditions.
	// Parameters:
	// - conds: optional conditions to filter the entities.
	// Returns the count of entities and an error if any occurs during the counting.
	Count(conds ...interface{}) (int64, error)

	// Where retrieves entities from the repository based on the specified query conditions.
	// Parameters:
	// - query: conditions to filter the entities.
	// Returns a slice of entities that match the query and an error if any occurs during the retrieval.
	Where(query interface{}) ([]T, error)
}

type repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &repository[T]{db}
}

// Create inserts a new entity into the repository.
// Parameters:
// - entity: the entity to insert.
// Returns an error if any occurs during the insertion.
func (r *repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

// FindByID retrieves an entity by its unique identifier.
// Parameters:
// - id: the unique identifier of the entity to retrieve.
// Returns a pointer to the entity if found, and an error if any occurs during the retrieval.
func (r *repository[T]) FindByID(id int64) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll retrieves all entities from the repository, optionally filtered by the provided conditions.
// If no conditions are provided, all entities are retrieved.
// Parameters:
// - conds: optional conditions to filter the entities.
// Returns a slice of entities and an error if any occurs during the retrieval.
func (r *repository[T]) FindAll(conds ...interface{}) ([]T, error) {
	var entities []T
	check := checkConds(conds)
	fmt.Println(check)
	if !check {
		result := r.db.Find(&entities)
		if result.Error != nil {
			return nil, result.Error
		}
		return entities, nil
	}
	err := r.db.Find(&entities, conds).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// FindPaginated retrieves a paginated list of entities based on the provided conditions.
// Parameters:
// - page: the page number to retrieve.
// - pageSize: the number of entities per page.
// - conds: optional conditions to filter the entities.
// Returns a slice of entities and an error if any occurs during the retrieval.
func (r *repository[T]) FindPaginated(page int, pageSize int, conds ...interface{}) ([]T, error) {
	var entities []T
	check := checkConds(conds)
	if !check {
		result := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&entities)
		if result.Error != nil {
			return nil, result.Error
		}
		return entities, nil
	}
	err := r.db.Offset((page-1)*pageSize).Limit(pageSize).Find(&entities, conds).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// FindOneBy retrieves a single entity based on the provided conditions.
// Parameters:
// - conds: optional conditions to filter the entity.
// Returns a pointer to the entity if found, and an error if any occurs during the retrieval.
func (r *repository[T]) FindOneBy(conds ...interface{}) (*T, error) {
	var entity T
	if err := r.db.First(&entity, conds).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete removes an entity from the repository.
// Parameters:
// - id: the unique identifier of the entity to remove.
// Returns an error if any occurs during the deletion.
func (r *repository[T]) Delete(id int64) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}

// Update updates an entity in the repository.
// Parameters:
// - entity: the entity to update.
// Returns an error if any occurs during the update.
func (r *repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

// Count returns the number of entities in the repository that match the provided conditions.
// Parameters:
// - conds: optional conditions to filter the entities.
// Returns the count of entities and an error if any occurs during the counting.
func (r *repository[T]) Count(conds ...interface{}) (int64, error) {
	var count int64
	var model T

	if err := r.db.Model(&model).Where(conds).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Where retrieves entities from the repository based on the specified query conditions.
// Parameters:
// - query: conditions to filter the entities.
// Returns a slice of entities that match the query and an error if any occurs during the retrieval.
func (r *repository[T]) Where(query interface{}) ([]T, error) {
	var entities []T
	result := r.db.Where(query).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}
