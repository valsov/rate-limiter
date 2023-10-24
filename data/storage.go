package data

type Storage[TConfig, TValue any] interface {
	Get(id string) (Data[TConfig, TValue], error)
	Set(id string, data Data[TConfig, TValue]) error
}

type Data[TConfig, TValue any] struct {
	HasConfig bool
	Config    TConfig
	Value     TValue
}

type memoryStore[TConfig, TValue any] struct {
	store map[string]Data[TConfig, TValue]
}

func NewMemoryStore[TConfig, TValue any]() Storage[TConfig, TValue] {
	return &memoryStore[TConfig, TValue]{make(map[string]Data[TConfig, TValue])}
}

func (s *memoryStore[TConfig, TValue]) Get(id string) (Data[TConfig, TValue], error) {
	return s.store[id], nil
}

func (s *memoryStore[TConfig, TValue]) Set(id string, data Data[TConfig, TValue]) error {
	s.store[id] = data
	return nil
}
