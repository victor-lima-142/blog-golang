package seeders

type Seeder[T any] interface {
	Seed(count *int) ([]T, error)
	SeedOne() (T, error)
}
