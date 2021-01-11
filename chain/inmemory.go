package chain

import "github.com/google/uuid"

type InMemoryDB struct {
	data map[string]map[string]string
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: map[string]map[string]string{},
	}
}

func (x *InMemoryDB) AppendItem(key, value string) error {
	m, ok := x.data[key]
	if !ok {
		m = make(map[string]string)
		x.data[key] = m
	}

	m[uuid.New().String()] = value
	return nil
}

func (x *InMemoryDB) FetchItems(key string) ([]string, error) {
	m, ok := x.data[key]
	if !ok {
		return nil, nil
	}

	var res []string
	for _, v := range m {
		res = append(res, v)
	}

	return res, nil
}
