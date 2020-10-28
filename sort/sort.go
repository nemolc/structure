package sort

type Sort struct {
	compare func(*interface{}, *interface{}) int
	hash    func(*interface{}) int
}

func NewSort(compare func(*interface{}, *interface{}) int, hash func(*interface{}) int) *Sort {
	return &Sort{compare: compare, hash: hash}
}


func (this *Sort) Bubble(list []interface{}) {
	length := len(list)
	for i := 0; i < length-1; i++ {
		for j := length - 1; j > i; j-- {
			if this.compare(&list[j], &list[j-1]) > 0 {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

func (this *Sort) Insert(list []interface{}) {
	length := len(list)
	for i := 1; i < length; i++ {
		tmp := list[i]
		j := i - 1
		for ; j >= 0; j-- {
			if this.compare(&list[j], &tmp) < 0 {
				list[j+1] = list[j]
			} else {
				break
			}
		}
		list[j+1] = tmp
	}
}

func (this *Sort) Select(list []interface{}) {
	length := len(list)
	var min_index int
	for i := 0; i < length-1; i++ {
		min_index = i
		for j := length - 1; j > i; j-- {
			if this.compare(&list[j], &list[min_index]) > 0 {
				min_index = j
			}
		}
		list[i], list[min_index] = list[min_index], list[i]
	}
}

func (this *Sort) Bubble2(list []interface{}) {
	for i := len(list) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if this.compare(&list[j], &list[j+1]) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func (this *Sort) Insert2(list []interface{}) {
	length := len(list)
	for i := len(list) - 1; i > 0; i-- {
		tmp := list[i-1]
		j := i
		for ; j < length; j++ {
			if this.compare(&list[j], &tmp) > 0 {
				list[j-1] = list[j]
			} else {
				break
			}
		}
		list[j-1] = tmp
	}
}

//效率低，乱序
func (this *Sort) Select2(list []interface{}) {
	var max_index int
	for i := len(list) - 1; i > 0; i-- {
		max_index = i
		for j := 0; j < i; j++ {
			if this.compare(&list[j], &list[max_index]) < 0 {
				max_index = j
			}
		}
		list[i], list[max_index] = list[max_index], list[i]
	}
}

func (this *Sort) Shell(list []interface{}) {
	length := len(list)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			tmp := list[i]
			j := i - gap
			for j >= 0 && this.compare(&list[j], &tmp) < 0 {
				list[j+gap] = list[j]
				j -= gap
			}
			list[j+gap] = tmp
		}
	}
}

func (this *Sort) Shell2(list []interface{}) {
	length := len(list)
	for i := 2; ; i *= 2 {
		gap := length / i
		if gap < 1 {
			break
		}
		for j := gap; j < length; j++ {
			tmp := list[j]
			k := j - gap
			for k >= 0 && this.compare(&list[k], &tmp) < 0 {
				list[k+gap] = list[k]
				k -= gap
			}
			list[k+gap] = tmp
		}
	}
}

func (this *Sort) Quick(list []interface{}) {
	this.quickHelp(list, 0, len(list)-1)
}

func (this *Sort) quickHelp(list []interface{}, start, end int) {
	if !(start < end) {
		return
	}
	direction := true
	pre := start
	post := end
	for pre < post {
		if direction {
			if this.compare(&list[post], &list[pre]) > 0 {
				list[pre], list[post] = list[post], list[pre]
				direction = false
				pre++
			} else {
				post--
			}
		} else {
			if this.compare(&list[pre], &list[post]) < 0 {
				list[post], list[pre] = list[pre], list[post]
				direction = true
				post--
			} else {
				pre++
			}
		}
	}
	this.quickHelp(list, start, pre-1)
	this.quickHelp(list, post+1, end)
}

//归并排序
func (this *Sort) Merge(list []interface{}) {
	tmp_list := make([]interface{}, len(list))
	copy(tmp_list, list)
	this.mergeHelp(list, 0, len(list)-1, tmp_list)
	return
}

func (this *Sort) mergeHelp(list []interface{}, start, end int, tmp_list []interface{}) {
	if !(start < end) {
		return
	}
	mid := (end + start) / 2
	this.mergeHelp(list, start, mid, tmp_list)
	this.mergeHelp(list, mid+1, end, tmp_list)
	i := start
	j := mid + 1
	begin := start
	for i <= mid && j <= end {
		if this.compare(&list[i], &list[j]) >= 0 {
			tmp_list[begin] = list[i]
			begin++
			i++
		} else {
			tmp_list[begin] = list[j]
			begin++
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp_list[begin] = list[i]
		begin++
	}
	for ; j <= end; j++ {
		tmp_list[begin] = list[j]
		begin++
	}
	for i := start; i <= end; i++ {
		list[i] = tmp_list[i]
	}
}

//堆排序
func (this *Sort) Heap(list []interface{}) {
	end := len(list) - 1
	start := (end - 1) / 2
	for i := start; i >= 0; i-- {
		this.siftAdjust(list, i, end)
	}
	for end > 0 {
		list[0], list[end] = list[end], list[0]
		end--
		this.siftAdjust(list, 0, end)
	}
}

func (this *Sort) siftAdjust(list []interface{}, start, end int) {
	for start < end {
		left := start*2 + 1
		if left > end {
			return
		}
		var big = left
		if left+1 <= end && this.compare(&list[left], &list[left+1]) > 0 {
			big = left + 1
		}
		if this.compare(&list[big], &list[start]) >= 0 {
			return
		}
		list[big], list[start] = list[start], list[big]
		start = big
	}
}


func (this *Sort) Count(list []interface{}) {
	length := len(list)
	if length < 2 {
		return
	}
	var min = this.hash(&list[0])
	var max = this.hash(&list[0])
	var tmp int
	for i := 1; i < length; i++ {
		tmp = this.hash(&list[i])
		if tmp < min {
			min = tmp
		}
		if tmp > max {
			max = tmp
		}
	}
	cache := make([]int, max-min+1)
	for i := 0; i < length; i++ {
		tmp = this.hash(&list[i])
		cache[tmp-min] ++
	}
	index := 0
	for k, count := range cache {
		for i := 0; i < count; i++ {
			list[index] = k + min
			index++
		}
	}
	return
}
