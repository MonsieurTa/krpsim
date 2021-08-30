package store

import "github.com/MonsieurTa/krpsim/internal/entity"

type Store interface {
	Duplicate() Store
	Consume(name string, amount int) (int, bool)
	ConsumeIfAvailable(stocks entity.Stocks) bool
	Store(name string, amount int) int
	BatchStore(stocks entity.Stocks)
	Value() store
	Get(name string) int
}
