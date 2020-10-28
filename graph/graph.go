package graph

import "structure/list"

type Graph struct {
	VexList     []GraphVex
	VertexCount int
	EdgeCount   int
	Directed    bool
}

type GraphVex struct {
	Key       string
	FirstFrom *GraphEdge
	FirstTo   *GraphEdge
}

type GraphEdge struct {
	Weight int
	Index  int
	Next   *GraphEdge
}

func NewGraph(directed bool) *Graph {
	return &Graph{VexList: make([]GraphVex, 0), VertexCount: 0, EdgeCount: 0, Directed: directed}
}

func (this *Graph) GetVexCount() int {
	return this.VertexCount
}

func (this *Graph) GetEdgeCount() int {
	if this.Directed == true {
		return this._GetEdgeCount()
	} else {
		return this._GetEdgeCount() / 2
	}
}

func (this *Graph) _GetEdgeCount() int {
	return this.EdgeCount
}

func (this *Graph) InsertVex(key string) (ok bool, index int, vex *GraphVex) {
	index, _ = this.GetIndex(key)
	if index == -1 {
		ok = true
		vex = &GraphVex{Key: key, FirstFrom: nil, FirstTo: nil}
		index = len(this.VexList)
		this.VexList = append(this.VexList, *vex)
		this.VertexCount++
		return
	} else {
		ok = false
		vex = &this.VexList[index]
		return
	}
}
func (this *Graph) AddEdge(v_i, u_i int, weight int) {
	if this.Directed == true {
		this._AddEdge(v_i, u_i, weight)
	} else {
		this._AddEdge(v_i, u_i, weight)
		this._AddEdge(u_i, v_i, weight)
	}
}

func (this *Graph) _AddEdge(v_i, u_i int, weight int) {
	v := &this.VexList[v_i]
	u := &this.VexList[u_i]
	to_edge := GraphEdge{Index: u_i, Weight: weight}
	if v.FirstTo == nil {
		v.FirstTo = &to_edge
	} else {
		cur := v.FirstTo
		for {
			if cur.Index == u_i {
				return
			}
			if cur.Next == nil {
				break
			} else {
				cur = cur.Next
			}
		}
		cur.Next = &to_edge
	}
	from_edge := GraphEdge{Index: v_i, Weight: weight}
	if u.FirstFrom == nil {
		u.FirstFrom = &from_edge
	} else {
		cur := u.FirstFrom
		for {
			if cur.Index == v_i {
				return
			}
			if cur.Next == nil {
				break
			} else {
				cur = cur.Next
			}
		}
		cur.Next = &from_edge
	}
	this.EdgeCount++
	return
}
func (this *Graph) DelEdge(v_i, u_i int) {
	if this.Directed == true {
		this._DelEdge(v_i, u_i)
	} else {
		this._DelEdge(v_i, u_i)
		this._DelEdge(u_i, v_i)
	}
}

func (this *Graph) _DelEdge(v_i, u_i int) {
	v := &this.VexList[v_i]
	u := &this.VexList[u_i]
	if v.FirstTo == nil {
		return
	}
	var last_cur **GraphEdge = &v.FirstTo
	var cur = v.FirstTo
	for cur != nil {
		if cur.Index == u_i {
			*last_cur = cur.Next
			goto del_from
		}
		last_cur = &cur
		cur = cur.Next
	}
	return
del_from:
	last_cur = &u.FirstFrom
	cur = v.FirstFrom
	for cur != nil {
		if cur.Index == v_i {
			*last_cur = cur.Next
			this.EdgeCount--
			return
		}
		last_cur = &cur
		cur = cur.Next
	}
	return
}

func (this *GraphVex) GetOutDegree() int {
	var n = 0
	for cur := this.FirstTo; cur != nil; cur = cur.Next {
		n++
	}
	return n
}

func (this *GraphVex) GetInDegree() int {
	var n = 0
	for cur := this.FirstFrom; cur != nil; cur = cur.Next {
		n++
	}
	return n
}

func (this *Graph) GetIndex(key string) (index int, vex *GraphVex) {
	for i, v := range this.VexList {
		if v.Key == key {
			return i, &this.VexList[i]
		}
	}
	return -1, nil
}

func (this *Graph) BFSTraverse(v_i int, visit func(*GraphVex)) {
	if v_i >= len(this.VexList) {
		return
	}
	tag_list := make([]bool, len(this.VexList))
	queue := list.NewQueue()
	queue.InQueue(v_i)
	tag_list[v_i] = true
	for {
		if ok, i := queue.OutQueue(); ok {
			index := i.(int)
			visit(&this.VexList[index])
			for cur := this.VexList[index].FirstTo; cur != nil; cur = cur.Next {
				if tag_list[cur.Index] == false {
					tag_list[cur.Index] = true
					queue.InQueue(cur.Index)
				}
			}
		} else {
			return
		}
	}
}

func (this *Graph) DFSTraverse(v_i int, visit func(*GraphVex)) {
	if v_i >= len(this.VexList) {
		return
	}
	tag_list := make([]bool, len(this.VexList))
	tag_list[v_i] = true
	this.DFSTraverseHelp(v_i, tag_list, visit)
}

func (this *Graph) DFSTraverseHelp(v_i int, tag_list []bool, visit func(*GraphVex)) {
	visit(&this.VexList[v_i])
	for cur := this.VexList[v_i].FirstTo; cur != nil; cur = cur.Next {
		if tag_list[cur.Index] == false {
			tag_list[cur.Index] = true
			this.DFSTraverseHelp(cur.Index, tag_list, visit)
		}
	}
}

func (this *Graph) DFSTraverseNonRecur(v_i int, visit func(*GraphVex)) {
	if v_i >= len(this.VexList) {
		return
	}
	tag_list := make([]bool, len(this.VexList))
	var i = v_i
	stack := list.NewStack()
	for {
		to_list := make([]int, 0, 4)
		if tag_list[i] == false {
			tag_list[i] = true
			visit(&this.VexList[i])
		} else {
			goto pop
		}
		for cur := this.VexList[i].FirstTo; cur != nil; cur = cur.Next {
			if tag_list[cur.Index] == false {
				to_list = append(to_list, cur.Index)
			}
		}
		if len(to_list) > 0 {
			for j := len(to_list) - 1; j > 0; j-- {
				stack.Push(to_list[j])
			}
			i = to_list[0]
			continue
		}
	pop:
		if ok, data := stack.Pop(); ok {
			i = data.(int)
		} else {
			return
		}

	}
}
