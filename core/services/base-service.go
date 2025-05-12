package services

import (
	"blog/core/repositories"
)

type Service[T any] interface {
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
	// Parameters:
	// - conds: optional conditions to filter the entities.
	// Returns a slice of entities and an error if any occurs during the retrieval.
	FindAll(conds ...interface{}) ([]T, error)

	// FindPaginated retrieves a paginated list of entities based on the provided conditions.
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

type service[T any] struct {
	r repositories.Repository[T]
}

func NewService[T any](r repositories.Repository[T]) Service[T] {
	return &service[T]{r}
}

// Create inserts a new entity into the repository.
// Parameters:
// - entity: the entity to insert.
// Returns an error if any occurs during the insertion.
func (s *service[T]) Create(entity *T) error {
	return s.r.Create(entity)
}

// FindByID retrieves an entity by its unique identifier.
// Parameters:
// - id: the unique identifier of the entity to retrieve.
// Returns a pointer to the entity if found, and an error if any occurs during the retrieval.
func (s *service[T]) FindByID(id int64) (*T, error) {
	return s.r.FindByID(id)
}

// FindAll retrieves all entities from the repository, optionally filtered by the provided conditions.
// Parameters:
// - conds: optional conditions to filter the entities.
// Returns a slice of entities and an error if any occurs during the retrieval.
func (s *service[T]) FindAll(conds ...interface{}) ([]T, error) {
	return s.r.FindAll(conds)
}

// FindOneBy retrieves a single entity based on the provided conditions.
// Parameters:
// - conds: optional conditions to filter the entity.
// Returns a pointer to the entity if found, and an error if any occurs during the retrieval.
func (s *service[T]) FindOneBy(conds ...interface{}) (*T, error) {
	return s.r.FindOneBy(conds)
}

// FindPaginated retrieves a paginated list of entities based on the provided conditions.
// Parameters:
// - page: the page number to retrieve.
// - pageSize: the number of entities per page.
// - conds: optional conditions to filter the entities.
// Returns a slice of entities and an error if any occurs during the retrieval.
func (s *service[T]) FindPaginated(page int, pageSize int, conds ...interface{}) ([]T, error) {
	return s.r.FindPaginated(page, pageSize, conds)
}

// Delete removes an entity from the repository.
// Parameters:
// - id: the unique identifier of the entity to remove.
// Returns an error if any occurs during the deletion.
func (s *service[T]) Delete(id int64) error {
	return s.r.Delete(id)
}

// Update updates an entity in the repository.
// Parameters:
// - entity: the entity to update.
// Returns an error if any occurs during the update.
func (s *service[T]) Update(entity *T) error {
	return s.r.Update(entity)
}

// Count returns the number of entities in the repository that match the provided conditions.
// Parameters:
// - conds: optional conditions to filter the entities.
// Returns the count of entities and an error if any occurs during the counting.
func (s *service[T]) Count(conds ...interface{}) (int64, error) {
	return s.r.Count(conds)
}

// Where retrieves entities from the repository based on the specified query conditions.
// Parameters:
// - query: conditions to filter the entities.
// Returns a slice of entities that match the query and an error if any occurs during the retrieval.
func (s *service[T]) Where(query interface{}) ([]T, error) {
	return s.r.Where(query)
}
