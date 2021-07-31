package store

type Store interface {
	Duplicate() Store
	Take(name string, amount int) (int, bool)
	Store(name string, amount int) int
}
