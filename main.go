package main

import (
	"fmt"
	"math/rand"
	"structure/collection"
	"structure/graph"
	"structure/list"
	"structure/search"
	"structure/sort"
	"structure/tree"
	"time"
)

const SUM = 1000000

type Str struct {
	A int
	b int
}

const zero = 0.0

func main() {
	sort_test()
}

func A(a int) {}

func t(m map[int]int) {
	m[4] = 4
}

func f() (r int) {
	r = 5
	defer func() {
		r = r + 5
	}()
	return
}

//v0->v1->v2->v5->v3->v6->v7->v8->v4->end
//v0->v1->v2->v5->v3->v6->v7->v8->v4->end

func graph_test() {
	graph := graph.NewGraph(false)
	_, i0, _ := graph.InsertVex("v0")
	_, i1, _ := graph.InsertVex("v1")
	_, i2, _ := graph.InsertVex("v2")
	_, i3, _ := graph.InsertVex("v3")
	_, i4, _ := graph.InsertVex("v4")
	_, i5, _ := graph.InsertVex("v5")
	_, i6, _ := graph.InsertVex("v6")
	_, i7, _ := graph.InsertVex("v7")
	_, i8, _ := graph.InsertVex("v8")

	graph.AddEdge(i0, i1, 0)
	graph.AddEdge(i0, i3, 0)
	graph.AddEdge(i0, i4, 0)
	graph.AddEdge(i1, i2, 0)
	graph.AddEdge(i1, i4, 0)
	graph.AddEdge(i2, i5, 0)
	graph.AddEdge(i3, i6, 0)
	graph.AddEdge(i4, i6, 0)
	graph.AddEdge(i6, i7, 0)
	graph.AddEdge(i7, i8, 0)
	//
	//graph.BFSTraverse(i0, visit_graph)
	//fmt.Println("end")

	graph.DFSTraverse(i0, visit_graph)
	fmt.Println("end")

	graph.DFSTraverseNonRecur(i0, visit_graph)
	fmt.Println("end")
}

func visit_graph(vex *graph.GraphVex) {
	fmt.Print(vex.Key, "->")
}

func search_test() {
	search := search.NewSearch(compare)
	li := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(search.Search2(li, 7))
}

func sort_test() {
	sort := sort.NewSort(compare, hash9)
	li := []interface{}{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < SUM; i++ {
		li = append(li, rand.Intn(100000))
	}
	//fmt.Println(li)

	var l []interface{}
	var t0 time.Time
	l = make([]interface{}, len(li))
	copy(l, li)
	t0 = time.Now()
	sort.Quick(l)
	//fmt.Println(l)

	//li = l

	l = make([]interface{}, len(li))
	copy(l, li)
	t0 = time.Now()
	sort.Quick(l)
	fmt.Println("quick:", time.Since(t0))

	l = make([]interface{}, len(li))
	copy(l, li)
	t0 = time.Now()
	sort.Count(l)
	//fmt.Println(l)
	fmt.Println("count:", time.Since(t0))

	l = make([]interface{}, len(li))
	copy(l, li)
	t0 = time.Now()
	sort.Merge(l)
	//fmt.Println(l)
	fmt.Println("merge:", time.Since(t0))

	l = make([]interface{}, len(li))
	copy(l, li)
	t0 = time.Now()
	sort.Heap(l)
	//fmt.Println(l)
	fmt.Println("heap:", time.Since(t0))

	l = make([]interface{}, len(li))
	copy(l, li)
	t0 = time.Now()
	sort.Shell(l)
	//fmt.Println(l)
	fmt.Println("shell:", time.Since(t0))

	l = make([]interface{}, len(li))
	copy(l, li)
	t0 = time.Now()
	sort.Shell2(l)
	//fmt.Println(l)
	fmt.Println("shell2:", time.Since(t0))



/*	l = make([]interface{}, len(li))
	copy(l, li)
	t0 = time.Now()
	sort.Insert(l)
	//fmt.Println(l)
	fmt.Println("insert:", time.Since(t0))

	l = make([]interface{}, len(li))
	copy(l, li)
	t1 := time.Now()
	sort.Select(l)
	//fmt.Println(l)
	fmt.Println("select:", time.Since(t1))

	l = make([]interface{}, len(li))
	copy(l, li)
	t1 = time.Now()
	sort.Bubble(l)
	//fmt.Println(l)
	fmt.Println("bubble:", time.Since(t1))*/
}

func map_test() {
	m := collection.NewMap()
	t0 := time.Now()
	for i := 0; i < SUM; i++ {
		n := rand.Int()
		m.Insert(n, n)
	}
	m.Print()
	fmt.Println("map_insert:", time.Since(t0))

}

func rb_test() {
	rb := tree.NewRB(compare)
	t0 := time.Now()
	for i := 0; i < SUM; i++ {
		rb.Insert(rand.Int())
	}
	fmt.Println("rb_insert:", time.Since(t0))
}

func avl_test() {
	avl := tree.NewAVL(compare)
	//var tmp int
	//var tmp_node *tree.AVLNode
	sum := 1000000
	for i := 0; i < sum; i++ {
		if sum/2 == i {
			//tmp = rand.Int()
			avl.Insert(i)
		} else {
			avl.Insert(i)
		}
	}
	//fmt.Println(tmp)
	//a := avl.Search(tmp)
	//fmt.Println(a.Data)
	//avl.Delete(tmp)
	//b := avl.Search(tmp)
	//fmt.Println(b)
	//avl.InTraverse(visit_avl)
	avl.Delete(33)
	avl.Delete(89)
	avl.Delete(2)
	//avl.InTraverse(visit_avl2)
	//last = 0
	fmt.Println(avl.Search(33))
	fmt.Println(avl.Search(89))
	fmt.Println(avl.Search(2))
	fmt.Println(avl.Height())

}

func compare(x *interface{}, y *interface{}) int {
	return (*y).(int) - (*x).(int)
}

func hash9(x *interface{}) int {
	return (*x).(int)
}

var last int

func visit_avl(data *interface{}) {
	cur := (*data).(int)
	if cur < last {
		panic(fmt.Sprintln("wrong:", last, cur))
	}
	last = cur
}

func visit_avl2(data *interface{}) {
	fmt.Print(*data, " ")
}

func queue_test() {
	var queue list.Queue
	for i := 0; i < 10; i++ {
		queue.InQueue(i)

	}
	queue.Traverse(visit_queue)
	fmt.Println()
	fmt.Println(queue.Length())
	for {
		if ok, data := queue.OutQueue(); ok {
			fmt.Print(data, " ")
		} else {
			return
		}
	}
}

func visit_queue(node *list.QueueNode) {
	fmt.Print(node.Data, " ")
}
func visit_tree(node *tree.TreeNode) {
	fmt.Print(node.Data, " ")

}

func tree_test2() {
	sum := 20000000

	avl := tree.NewAVL(compare)
	t1 := time.Now()
	for i := 0; i < sum; i++ {
		avl.Insert(i)
	}
	fmt.Println("avl_tree insert:", time.Since(t1))

	rb := tree.NewRB(compare)
	t0 := time.Now()
	for i := 0; i < sum; i++ {
		rb.Insert(i)
	}
	fmt.Println("rb_tree insert:", time.Since(t0))

	t5 := time.Now()
	for i := 0; i < sum; i++ {
		avl.Search(i)
	}
	fmt.Println("avl_tree search:", time.Since(t5))

	t4 := time.Now()
	for i := 0; i < sum; i++ {
		rb.Search(i)
	}
	fmt.Println("rb_tree search:", time.Since(t4))

	t2 := time.Now()
	for i := 0; i < sum; i++ {
		avl.Delete(i)
	}
	fmt.Println("avl_tree delete:", time.Since(t2))

	t3 := time.Now()
	for i := 0; i < sum; i++ {
		rb.Delete(i)
	}
	fmt.Println("rb_tree delete:", time.Since(t3))
}

func tree_test() {
	var tree tree.BitTree;
	fmt.Println(tree)
	A := tree.InitRoot("A")
	B := A.InsertLChild("B")
	D := B.InsertLChild("D")
	D.InsertRChild("G")
	C := A.InsertRChild("C")
	C.InsertLChild("E")
	C.InsertRChild("F")
	//tree.PreTraverse(visit_tree)
	//fmt.Println()
	//tree.InTraverse(visit_tree)
	//fmt.Println()
	mirror:=tree.Mirror()
	mirror.PreTraverse(visit_tree)
	fmt.Println()
	mirror.InTraverse(visit_tree)
	fmt.Println()
	mirror.LevelTraverse(visit_tree)
	fmt.Println()
	mirror.PostTraverse(visit_tree)

}
