package bs_test

import (
	"gotour/bs"
	"testing"
)

func TestSearch(t *testing.T) {
	l := bs.List{}
	if _, ok := l.Search(0); ok {
		t.Error("Error")
	}

	l2 := bs.List{0, 3, 5, 10}
	if v, ok := l2.Search(0); !ok || v != 0 {
		t.Error("Error")
	}
	if v, ok := l2.Search(10); !ok || v != 3 {
		t.Error("Error")
	}
}
