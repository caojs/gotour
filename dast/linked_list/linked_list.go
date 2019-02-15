package linked_list

type Node struct {
	next  *Node
	Value interface{}
}

type LinkedList struct {
	head   *Node
	Length int
}

func (list LinkedList) Get(position int) (interface{}, bool) {
	if position > list.Length-1 {
		return 0, false
	}

	h := list.head
	for count := 0; count != position; count = count + 1 {
		h = h.next
	}

	return h.Value, true
}

func (list *LinkedList) Add(value interface{}, position int) bool {
	if position > list.Length {
		return false
	}

	if position == 0 {
		list.head = &Node{list.head, value}
	} else {
		prev := list.head
		for count := 0; count != position-1; count = count + 1 {
			prev = prev.next
		}
		prev.next = &Node{prev.next, value}
	}

	list.Length++
	return true
}

func (list *LinkedList) Remove(position int) bool {
	if position > list.Length-1 {
		return false
	}

	if position == 0 {
		list.head = list.head.next
	} else {
		prev := list.head
		for count := 1; count != position; count = count + 1 {
			prev = prev.next
		}

		next := prev.next.next
		prev.next = next
	}

	list.Length--
	return true
}
