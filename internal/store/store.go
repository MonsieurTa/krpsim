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
	for _, stock := range cfg.Stocks {
		s.Store(stock.Key, stock.Value)
	}

	for _, process := range cfg.Processes {
		for _, need := range process.Needs {
			s.Store(need.Key, 0)
		}
		for _, result := range process.Results {
			s.Store(result.Key, 0)
		}
	}
}

func (s store) Take(name string, amount int) (int, bool) {
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
