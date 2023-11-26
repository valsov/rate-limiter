package storage

type Storage[TConfig, TValue any] interface {
	Get(id string) (Data[TConfig, TValue], error)
	Set(id string, data Data[TConfig, TValue]) error
}

type Data[TConfig, TValue any] struct {
	HasConfig bool
	Config    TConfig
	Value     TValue
}
