package linked_list_test

import (
	"gotour/dast/linked_list"
	"testing"
)

func TestAdd(t *testing.T) {
	list := linked_list.LinkedList{}
	ok := list.Add(1, 1)
	if ok {
		t.Error("error")
	}

	ok = list.Add(1, 0)
	if !ok {
		t.Error("error")
	}

	ok = list.Add(2, 1)
	if !ok {
		t.Error("error")
	}
}

func TestGet(t *testing.T) {
	list := linked_list.LinkedList{}
	list.Add(1, 0)
	list.Add(3, 1)
	list.Add(2, 2)

	if list.Length != 3 {
		t.Error("Error")
	}

	v, ok := list.Get(1)
	if !ok || v != 3 {
		t.Error("Error")
	}

	v, ok = list.Get(0)
	if !ok || v != 1 {
		t.Error("Error")
	}
}

func TestRemove(t *testing.T) {
	list := linked_list.LinkedList{}
	list.Add(1, 0)
	list.Add(3, 1)
	list.Add(2, 2)

	ok := list.Remove(1)
	if !ok {
		t.Error("Error")
	}

	if list.Length != 2 {
		t.Error("Error")
	}

	v, _ := list.Get(1)
	if v != 2 {
		t.Error("Error")
	}
	v, _ = list.Get(0)
	if v != 1 {
		t.Error("Error")
	}

	list.Add(5, 2)
	if list.Length != 3 {
		t.Error("Error")
	}

	ok = list.Remove(3)
	if ok {
		t.Error("Error")
	}

	ok = list.Remove(0)
	if !ok {
		t.Error("Error")
	}

	v, _ = list.Get(0)
	if v != 2 {
		t.Error("Error")
	}
}
