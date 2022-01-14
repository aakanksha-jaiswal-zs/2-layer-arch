package customer

import "sample-api/stores"

type store struct {
	db map[string]int
}

func New(db map[string]int) stores.Customer {
	return store{db: db}
}

func (s store) Get(name string) bool {
	_, ok := s.db[name]

	return ok
}

func (s store) Create(name string, id int) {
	s.db[name] = id
}
