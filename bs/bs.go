package bs

type List []int

func (list List) Search(value int) (int, bool) {
	l := 0
	h := len(list) - 1

	for {
		if l > h {
			return 0, false
		}

		m := (l + h) / 2
		v := list[m]

		switch true {
		case v == value:
			return m, true
		case v > value:
			h = m - 1
		default:
			l = m + 1
		}
	}
}
