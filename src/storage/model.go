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

func (s *Storage) SetKey(key string) {
	s.key = key
}

func (s *Storage) SetValue(value string) {
	s.value = value
}

func (s *Storage) SetKeyValue(key string, value string) {
	s.key = key
	s.value = value
}

func NewStorage(key string, value string) *Storage {
	return &Storage{
		key:   key,
		value: value,
	}
}
