// файл persistent/persistent_test.go

package persistent

import (
	"github.com/golang/mock/gomock"
	"github.com/polosaty/go_developer_sprint3_gomock/mocks"
	"github.com/polosaty/go_developer_sprint3_gomock/store"
	"testing"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockStore(ctrl)

	// гарантируем, что заглушка
	// при вызове с аргументом "Key" вернёт "Value"
	s.EXPECT().Get("Key").Return([]byte("Value"), nil)

	// тестируем функцию
	someCond := store.Cond{Key: "Key"}
	if string(Lookup(s, someCond)) != "Value" {
		t.Fatal("fail")
	}

}

func TestGetWithStabs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockStore(ctrl)

	// при вызове с произвольным аргументом
	// заглушка будет возвращать слайс длиной 8 (точно не nil)
	// метод может быть вызван не более 5 раз

	s.EXPECT().
		Get(gomock.Any()).
		Return(
			gomock.All(
				gomock.Not(gomock.Nil()),
				gomock.Len(8),
			),
			nil).
		MaxTimes(5)

	// тестируем функцию
	someCond := store.Cond{Key: "Key"}
	x := Lookup(s, someCond)
	if len(x) != 8 {
		t.Fatal("fail", string(x), "len != 8")
	}
}
