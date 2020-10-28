package search

type Search struct {
	compare func(*interface{}, *interface{}) (int)
}

func NewSearch(compare func(*interface{}, *interface{}) (int)) *Search {
	return &Search{compare: compare}
}

func (this *Search) Search2(list []interface{}, v interface{}) int {
	if len(list) == 0 {
		return -1
	}
	x := 0
	y := len(list) - 1
	for x <= y {
		i := (x + y) / 2
		cmp := this.compare(&list[i], &v)
		if cmp == 0 {
			return i
		} else if cmp > 0 {
			x = i + 1
		} else {
			y = i - 1
		}
	}
	return -1
}

func (this *Search) Search(list []interface{}, v interface{}) int {
	if len(list) == 0 {
		return -1
	}
	return this.searchHelp(list, v, 0, len(list)-1)
}

func (this *Search) searchHelp(list []interface{}, v interface{}, start, end int) int {
	i := (start + end) / 2
	cmp := this.compare(&list[i], &v)
	if cmp == 0 {
		return i
	} else if cmp > 0 {
		if i < end {
			return this.searchHelp(list, v, i, end)
		} else {
			return -1
		}
	} else {
		if start < i {
			return this.searchHelp(list, v, start, i)

		} else {
			return -1
		}
	}

}
