package store

import "github.com/MonsieurTa/krpsim/internal/entity"

type store map[string]int

func New(cfg *entity.Config) Store {
	rv := store{}
	rv.init(cfg)
	return rv
}

func (s store) Duplicate() Store {
	rv := make(store, len(s))

	for key, value := range s {
		rv[key] = value
	}
	return rv
}

func (s store) init(cfg *entity.Config) {
	for name, value := range cfg.Stocks {
		s.Store(name, value)
	}

	for _, process := range cfg.Processes {
		for name := range process.Needs {
			s.Store(name, 0)
		}
		for name := range process.Results {
			s.Store(name, 0)
		}
	}
}

func (s store) ConsumeIfAvailable(stocks entity.Stocks) bool {
	for name, value := range stocks {
		if !s.available(name, value) {
			return false
		}
	}
	for name, value := range stocks {
		s.Consume(name, value)
	}
	return true
}

func (s store) available(name string, amount int) bool {
	v, ok := s[name]
	if !ok || v < amount {
		return false
	}
	return true
}

func (s store) Consume(name string, amount int) (int, bool) {
	v, ok := s[name]
	newAmount := v - amount
	if !ok || newAmount < 0 {
		return -1, false
	}
	s[name] = newAmount
	return amount, true
}

func (s store) Store(name string, amount int) int {
	v, ok := s[name]
	if !ok {
		s[name] = amount
		return amount
	}
	rv := v + amount
	s[name] = rv
	return rv
}

func (s store) BatchStore(stocks entity.Stocks) {
	for name, amount := range stocks {
		s.Store(name, amount)
	}
}

func (s store) Value() store {
	return s
}
