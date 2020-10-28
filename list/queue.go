package list

type Queue struct {
	head *QueueNode
}

type QueueNode struct {
	Data interface{}
	last *QueueNode
	next *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (this *Queue) Length() int {
	if this.head == nil {
		return 0
	}
	length := 1
	for current := this.head.next; current != this.head; current = current.next {
		length++
	}
	return length
}

func (this *Queue) Empty() bool {
	return this.head == nil
}

func (this *Queue) Traverse(visit func(*QueueNode)) {
	if this.head == nil {
		return
	}
	next := this.head.next
	visit(this.head)
	for ; next != this.head; next = next.next {
		visit(next)
	}

}

func (this *Queue) OutQueue() (ok bool, data interface{}) {
	if this.head == nil {
	} else if this.head == this.head.next {
		ok = true
		data = this.head.Data
		this.head = nil
	} else {
		head := this.head
		ok = true
		data = head.Data
		this.head = head.next
		head.last.next = head.next
		head.next.last = head.last
	}
	return
}

func (this *Queue) GetHead() (ok bool, data interface{}) {
	if this.head == nil {
		return
	}
	return true, this.head.Data

}

func (this *Queue) InQueue(data interface{}) {
	new_node := QueueNode{Data: data}
	if this.head == nil {
		new_node.last = &new_node
		new_node.next = &new_node
		this.head = &new_node
	} else {
		new_node.last = this.head.last
		new_node.next = this.head
		this.head.last.next = &new_node
		this.head.last = &new_node
	}
}
