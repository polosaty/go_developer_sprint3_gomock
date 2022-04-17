package persistent

// файл persistent/persistent.go
import (
	"github.com/polosaty/go_developer_sprint3_gomock/store"
)

func Lookup(s store.Store, c store.Cond) []byte {
	//...
	val, _ := s.Get(c.Key)
	return val
}
