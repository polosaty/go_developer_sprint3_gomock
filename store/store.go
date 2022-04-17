// файл store/store.go

package store

type Cond struct {
	Key string
}

type Store interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}
