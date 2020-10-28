package collection

import (
	"fmt"
	"hash/fnv"
)

type HashMap struct {
	size int
	ht   []*MapNode
}

type MapNode struct {
	key   interface{}
	value interface{}
	next  *MapNode
}

const (
	MIN_LEN = 16
)

func NewMap() *HashMap {
	m := HashMap{size: 0, ht: make([]*MapNode, MIN_LEN, MIN_LEN)}
	return &m
}

func (this *HashMap) Insert(k, v interface{}) {
	fnv.New32()
	if this.insertHelp(k, v) {
		this.size++
	}
	this.extend()
}

func (this *HashMap) Print() {
	fmt.Printf("size:%d,\nlen:%d\nclashed_count:%d\n,", this.size, len(this.ht), this.clashedCount())
}

func (this *HashMap) clashedCount() int {
	count := 0
	for _, cur := range this.ht {
		for cur != nil {
			if cur.next != nil {
				count++
				cur = cur.next
			} else {
				break
			}
		}
	}
	return count
}

func (this *HashMap) insertHelp(k, v interface{}) ( bool) {
	hash_code := this.getAddr(&k)
	var cur = this.ht[hash_code]
	if cur == nil {
		this.ht[hash_code] = &MapNode{key: k, value: v, next: nil}
		return true
	}
	for {
		if cur.key == k {
			cur.value = v
			return false
		}
		if cur.next != nil {
			cur = cur.next
		} else {
			break
		}
	}
	cur.next = &MapNode{key: k, value: v, next: nil}
	return true
}

func (this *HashMap) Delete(k interface{}) {
	hash_code := this.getAddr(&k)
	var cur = this.ht[hash_code]
	var last *MapNode = nil
	for cur != nil {
		if cur.key == k {
			goto del
		}
		last = cur
		cur = cur.next
	}
	return
del:
	defer func() {
		this.size--
		this.reduce()
	}()
	if last == nil {
		this.ht[hash_code] = cur.next
	} else {
		last.next = cur
	}
	return
}

func (this *HashMap) Get(k interface{}) (bool, interface{}) {
	hash_code := this.getAddr(&k)
	cur := this.ht[hash_code]
	for ; cur != nil; cur = cur.next {
		if cur.key == k {
			return true, cur.value
		}
	}
	return false, nil
}

func (this *HashMap) hashCode(k *interface{}) int {
	var code int
	switch real_key := (*k).(type) {
	case uint8:
		code = int(real_key)
	case uint16:
		code = int(real_key)
	case uint32:
		code = int(real_key)
	case uint64:
		code = int(real_key)
	case int8:
		code = int(real_key)
	case int16:
		code = int(real_key)
	case int32:
		code = int(real_key)
	case int64:
		code = int(real_key)
	case int:
		code = real_key
	default:
		panic("未实现此关键词类型散列")
	}
	return code
}

func (this *HashMap) getAddr(k *interface{}) int {
	return this.hashCode(k) % len(this.ht)
}

func (this *HashMap) extend() {
	if this.size > (len(this.ht) / 4 * 3) {
		tmp := this.ht
		this.ht = make([]*MapNode, len(this.ht)*2, len(this.ht)*2)
		for _, node := range tmp {
			for ; node != nil; node = node.next {
				this.insertHelp(node.key, node.value)
			}
		}
	}
}

func (this *HashMap) reduce() {
	if this.size < (len(this.ht)>>3) && len(this.ht) > MIN_LEN {
		tmp := this.ht
		size := this.size * 2
		if size < MIN_LEN {
			size = MIN_LEN
		}
		this.ht = make([]*MapNode, size, size)
		for _, node := range tmp {
			for ; node != nil; node = node.next {
				this.insertHelp(node.key, node.value)
			}
		}
	}
}
