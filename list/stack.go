package list

type Stack struct {
	top *StackNode
}

type StackNode struct {
	Data interface{}
	last *StackNode
}

func NewStack()*Stack{
	return &Stack{}
}

func (this *Stack) Length() int {
	length := 0
	for cur := this.top; cur != nil; cur = cur.last {
		length++
	}
	return length
}

func (this *Stack) Empty() bool {
	return this.top == nil
}

func (this *Stack) Clear() {
	this.top = nil
}

func (this *Stack) Traverse(visit func(*StackNode)) {
	for cur := this.top; cur != nil; cur = cur.last {
		visit(cur)
	}
	return
}

func (this *Stack) Push(data interface{}) {
	node := StackNode{Data: data, last: this.top}
	this.top = &node
	return
}

func (this *Stack) Pop() (ok bool, data interface{}) {
	if this.top == nil {
		return
	}
	ok = true
	data = this.top.Data
	this.top = this.top.last
	return
}

func (this *Stack) Top() (ok bool, data interface{}) {
	if this.top == nil {
		return
	}
	ok = true
	data = this.top.Data
	return
}
