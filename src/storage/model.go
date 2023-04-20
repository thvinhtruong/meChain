package storage

type Storage struct {
	key   string
	value string
}

func (s *Storage) Key() string {
	return s.key
}

func (s *Storage) Value() string {
	return s.value
}
