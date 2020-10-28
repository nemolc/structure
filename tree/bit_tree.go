package tree

import (
	"structure/list"
)

//type BitTreeInter interface {
//	Clear()
//	NodeCount()
//	Height()
//	PreTraverseNonRecur(visit func(*TreeNode))
//	PreTraverse(visit func(*TreeNode))
//	IntraverseNonRecur(visit func(*TreeNode))
//	InTraverse(visit func(*TreeNode))
//	PostTraverseNonRecur(visit func(*TreeNode))
//	PostTraverse(visit func(*TreeNode))
//	LevelTraverse(visit func(*TreeNode))
//}

type BitTree struct {
	root *TreeNode
}

type TreeNode struct {
	Data    interface{}
	parent  *TreeNode
	l_child *TreeNode
	r_child *TreeNode
}

func (this *BitTree) Clear() {
	this.root = nil
}

func (this *BitTree) NodeCount() (int) {
	return node_count_help(this.root)
}

func node_count_help(node *TreeNode) (int) {
	if node == nil {
		return 0
	}
	return 1 + node_count_help(node.l_child) + node_count_help(node.r_child)

}
func (this *BitTree) Height() (int) {
	return height_help(this.root)
}

func height_help(node *TreeNode) (int) {
	if node == nil {
		return 0
	}
	l_height := height_help(node.l_child)
	r_height := height_help(node.r_child)
	if l_height > r_height {
		return 1 + l_height
	} else {
		return 1 + r_height
	}

}

func (this *BitTree) PreTraverseNonRecur(visit func(*TreeNode)) {
	if this.root == nil {
		return
	}
	stack := list.Stack{}
	cur := this.root
	for {
		visit(cur)
		if cur.r_child != nil {
			stack.Push(cur.r_child)
		}
		if cur.l_child != nil {
			cur = cur.l_child
		} else {
			if ok, data := stack.Pop(); ok {
				cur = data.(*TreeNode)
			} else {
				return
			}
		}

	}
}

func (this *BitTree) PreTraverseNonRecur2(visit func(*TreeNode)) {
	if this.root == nil {
		return
	}
	stack := list.Stack{}
	cur := this.root
	for {
		visit(cur)
		if cur.l_child != nil {
			if cur.r_child != nil {
				stack.Push(cur.r_child)
			}
			cur = cur.l_child
		} else if cur.r_child != nil {
			cur = cur.r_child
		} else {
			if ok, data := stack.Pop(); ok {
				cur = data.(*TreeNode)
			} else {
				return
			}
		}
	}
}

func (this *BitTree) PreTraverse(visit func(*TreeNode)) {
	preTraverseHelp(this.root, visit)
}

func preTraverseHelp(node *TreeNode, visit func(*TreeNode)) {
	if node == nil {
		return
	}
	visit(node)
	preTraverseHelp(node.l_child, visit)
	preTraverseHelp(node.r_child, visit)
}

func (this *BitTree) IntraverseNonRecur(visit func(*TreeNode)) {
	if this.root == nil {
		return
	}
	stack := list.Stack{}
	cur := this.root
	for {
		if cur.l_child != nil {
			//如果有左子树则入栈
			stack.Push(cur)
			cur = cur.l_child
		} else {
			//无左子树访问
			visit(cur)
			if cur.r_child != nil {
				//有右子数，则从右子树继续寻找
				cur = cur.r_child
			} else {
			again:
				if ok, data := stack.Pop(); ok {
					cur = data.(*TreeNode)
					visit(cur)
					if cur.r_child != nil {
						cur = cur.r_child
					} else {
						goto again
					}
				} else {
					return
				}
			}
		}
	}

}

func (this *BitTree) InTraverse(visit func(*TreeNode)) {
	inTraverseHelp(this.root, visit)
}

func inTraverseHelp(node *TreeNode, visit func(*TreeNode)) {
	if node == nil {
		return
	}
	inTraverseHelp(node.l_child, visit)
	visit(node)
	inTraverseHelp(node.r_child, visit)
}

type modiNode struct {
	node     *TreeNode
	no_visit bool
}

func (this *BitTree) PostTraverseNonRecur(visit func(*TreeNode)) {
	if this.root == nil {
		return
	}
	stack := list.Stack{}
	cur := this.root
	for {
		if cur.l_child != nil {
			stack.Push(modiNode{node: cur})
			if cur.r_child != nil {
				stack.Push(modiNode{cur.r_child, true})
			}
			cur = cur.l_child
		} else if cur.r_child != nil {
			stack.Push(modiNode{node: cur})
			cur = cur.r_child
		} else {
		again:
			visit(cur)
			if ok, data := stack.Pop(); ok {
				m_node := data.(modiNode)
				cur = m_node.node
				if m_node.no_visit == false {
					goto again
				}
			} else {
				return
			}
		}
	}
}

func (this *BitTree) PostTraverse(visit func(*TreeNode)) {
	postTraverseHelp(this.root, visit)
}

func postTraverseHelp(node *TreeNode, visit func(*TreeNode)) {
	if node == nil {
		return
	}
	postTraverseHelp(node.l_child, visit)
	postTraverseHelp(node.r_child, visit)
	visit(node)
}

func (this *BitTree) LevelTraverse(visit func(*TreeNode)) {
	if this.root == nil {
		return
	}
	queue := list.Queue{}
	queue.InQueue(this.root)
	for {
		if ok, data := queue.OutQueue(); ok {
			node := data.(*TreeNode)
			visit(node)
			if node.l_child != nil {
				queue.InQueue(node.l_child)
			}
			if node.r_child != nil {
				queue.InQueue(node.r_child)
			}
		} else {
			return
		}
	}
}

func (this *BitTree) GetRoot() (*TreeNode) {
	return this.root
}

func (this *BitTree) InitRoot(data interface{}) (*TreeNode) {
	this.root = &TreeNode{Data: data}
	return this.root
}

func (this *BitTree) Mirror() (*BitTree) {
	mirror := &BitTree{}
	if this.root == nil {
		return mirror
	}
	mirror_node := *this.root
	mirror.root = &mirror_node
	this.MirrorHelp(mirror.root)
	return mirror
}

func (this *BitTree) MirrorHelp(mirror_node *TreeNode) {
	if mirror_node == nil {
		return
	}
	var mirror_right_child *TreeNode
	if mirror_node.l_child != nil {
		r_c := *mirror_node.l_child
		mirror_right_child = &r_c

	}
	var mirror_left_child *TreeNode
	if mirror_node.r_child != nil {
		l_c := *mirror_node.r_child
		mirror_left_child = &l_c

	}
	mirror_node.r_child, mirror_node.l_child = mirror_right_child, mirror_left_child
	this.MirrorHelp(mirror_node.l_child)
	this.MirrorHelp(mirror_node.r_child)
}

func (this *BitTree) MirrorNonRecur() (*BitTree) {
	return nil
}

func (this *TreeNode) InsertLChild(data interface{}) (*TreeNode) {
	new_node := TreeNode{Data: data, parent: this}
	current_l_child := this.l_child
	this.l_child = &new_node
	if current_l_child != nil {
		new_node.l_child = current_l_child
		current_l_child.parent = &new_node
	}
	return &new_node
}

func (this *TreeNode) InsertRChild(data interface{}) (*TreeNode) {
	new_node := TreeNode{Data: data, parent: this}
	current_r_child := this.r_child
	this.r_child = &new_node
	if current_r_child != nil {
		new_node.r_child = current_r_child
		current_r_child.parent = &new_node
	}
	return &new_node
}

func (this *TreeNode) DelLChild() {
	this.l_child = nil
}

func (this *TreeNode) DelRChild() {
	this.r_child = nil
}

func (this *TreeNode) GetParent() (*TreeNode) {
	return this.parent
}

func (this *TreeNode) GetLChild() (*TreeNode) {
	return this.l_child
}

func (this *TreeNode) GetRChild() (*TreeNode) {
	return this.r_child
}
