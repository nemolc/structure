package tree

import "structure/list"

type AVL struct {
	root    *AVLNode
	compare func(*interface{}, *interface{}) (int)
}

func NewAVL(compare func(*interface{}, *interface{}) (int)) *AVL {
	return &AVL{root: nil, compare: compare}
}

type AVLNode struct {
	height int
	Data   interface{}
	//parent  *TreeNode
	l_child *AVLNode
	r_child *AVLNode
}

func (this *AVL) Height() int {
	if this.root == nil {
		return 0
	}
	return this.root.height
}

func (this *AVL) Insert(data interface{}) {
	this.insertHelp(&this.root, &data)
}

func (this *AVL) Delete(data interface{}) {
	this.deleteHelp(&this.root, &data)
}

func (this *AVL) Search(data interface{}) *AVLNode {
	return this.searchHelp(this.root, &data)
}

func (this *AVL) IntraverseNonRecur(visit func(*interface{})) {
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
			visit(&cur.Data)
			if cur.r_child != nil {
				//有右子数，则从右子树继续寻找
				cur = cur.r_child
			} else {
			again:
				if ok, data := stack.Pop(); ok {
					cur = data.(*AVLNode)
					visit(&cur.Data)
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

func (this *AVL) InTraverse(visit func(*interface{})) {
	avl_inTraverseHelp(this.root, visit)
}

func (this *AVL) insertHelp(node **AVLNode, data *interface{}) {
	if *node == nil {
		*node = &AVLNode{Data: *data, height: 1}
		return
	}
	cmp := this.compare(&(*node).Data, data)
	if cmp < 0 {
		this.insertHelp(&(*node).l_child, data)
		if (*node).l_child.Height()-(*node).r_child.Height() == 2 {
			if (this.compare(&(*node).l_child.Data, data) < 0) {
				*node = LLRotationHelp(*node)
			} else {
				*node = LRRotationHelp(*node)
			}
		} else {
			(*node).setHeight()
		}
	} else if cmp > 0 {
		this.insertHelp(&(*node).r_child, data)
		if (*node).r_child.Height()-(*node).l_child.Height() == 2 {
			if (this.compare(&(*node).r_child.Data, data) < 0) {
				*node = RLRotationHelp(*node)
			} else {
				*node = RRRotationHelp(*node)
			}
		} else {
			(*node).setHeight()
		}
	} else {
		return
	}
	return
}

func (this *AVL) deleteHelp(node **AVLNode, data *interface{}) {
	if *node == nil {
		return
	}
	cmp := this.compare(&(*node).Data, data)
	if cmp < 0 {
		this.deleteHelp(&(*node).l_child, data)
		if (*node).r_child.Height()-(*node).l_child.Height() == 2 {
			r := (*node).r_child
			if (r.l_child.Height() > r.r_child.Height()) {
				*node = RLRotationHelp(*node)
			} else {
				*node = RRRotationHelp(*node)
			}
		} else {
			(*node).setHeight()
		}
	} else if cmp > 0 {
		this.deleteHelp(&(*node).r_child, data)
		if (*node).l_child.Height()-(*node).r_child.Height() == 2 {
			l := (*node).l_child
			if (l.r_child.Height() > l.l_child.Height()) {
				*node = LRRotationHelp(*node)
			} else {
				*node = LLRotationHelp(*node)
			}
		} else {
			(*node).setHeight()
		}
	} else {
		if (*node).l_child != nil && (*node).r_child != nil {
			if (*node).l_child.Height() > (*node).r_child.Height() {
				max := (*node).l_child.findMax()
				(*node).Data = max.Data
				this.deleteHelp(&(*node).l_child, &max.Data)
			} else {
				min := (*node).r_child.findMin()
				(*node).Data = min.Data
				this.deleteHelp(&(*node).r_child, &min.Data)
			}
			(*node).setHeight()
		} else {
			if (*node).l_child != nil {
				*node = (*node).l_child
			} else {
				*node = (*node).r_child
			}
		}
	}
	return
}

func (this *AVL) searchHelp(node *AVLNode, data *interface{}) *AVLNode {
	if node == nil {
		return nil
	}
	cmp := this.compare(&node.Data, data)
	if cmp < 0 {
		return this.searchHelp(node.l_child, data)
	} else if cmp > 0 {
		return this.searchHelp(node.r_child, data)
	} else {
		return node
	}
}

func (this *AVLNode) Height() int {
	if this == nil {
		return 0
	}
	return this.height
}

func (this *AVLNode) setHeight() {
	if this == nil {
		return
	}
	if this.l_child.Height() > this.r_child.Height() {
		this.height = this.l_child.Height() + 1
	} else {
		this.height = this.r_child.Height() + 1
	}
}

func (this *AVLNode) findMax() *AVLNode {
	cur := this
	for (cur.r_child != nil) {
		cur = cur.r_child
	}
	return cur
}

func (this *AVLNode) findMin() *AVLNode {
	cur := this
	for (cur.l_child != nil) {
		cur = cur.l_child
	}
	return cur
}

//func (this *AVL) Balance() {
//	this.balanceHelp(this.root)
//}
//
//func (this *AVL) balanceHelp(node *AVLNode) {
//	if node == nil {
//		return
//	}
//	bla := node.l_child.Height() - node.r_child.Height()
//	if -2 <= bla && bla <= 2 {
//		this.balanceHelp(node.l_child)
//		this.balanceHelp(node.r_child)
//	} else {
//		panic(fmt.Sprintln(bla))
//	}
//}

func avl_inTraverseHelp(node *AVLNode, visit func(*interface{})) {
	if node == nil {
		return
	}
	avl_inTraverseHelp(node.l_child, visit)
	visit(&node.Data)
	avl_inTraverseHelp(node.r_child, visit)
}

func rightRotationHelp(node *AVLNode) (root *AVLNode) {
	root = node.l_child
	node.l_child = root.r_child
	root.r_child = node
	if node.l_child.Height() > node.r_child.Height() {
		node.height = node.l_child.Height() + 1
	} else {
		node.height = node.r_child.Height() + 1
	}
	if root.l_child.Height() > root.r_child.Height() {
		root.height = root.l_child.Height() + 1
	} else {
		root.height = root.r_child.Height() + 1
	}
	return
}

func leftRotationHelp(node *AVLNode) (root *AVLNode) {
	root = node.r_child
	node.r_child = root.l_child
	root.l_child = node
	if node.l_child.Height() > node.r_child.Height() {
		node.height = node.l_child.Height() + 1
	} else {
		node.height = node.r_child.Height() + 1
	}
	if root.l_child.Height() > root.r_child.Height() {
		root.height = root.l_child.Height() + 1
	} else {
		root.height = root.r_child.Height() + 1
	}
	return
}

func LLRotationHelp(node *AVLNode) (root *AVLNode) {
	return rightRotationHelp(node)
}

func RRRotationHelp(node *AVLNode) (root *AVLNode) {
	return leftRotationHelp(node)
}

func LRRotationHelp(node *AVLNode) (root *AVLNode) {
	node.l_child = leftRotationHelp(node.l_child)
	return rightRotationHelp(node)
}

func RLRotationHelp(node *AVLNode) (root *AVLNode) {
	node.r_child = rightRotationHelp(node.r_child)
	return leftRotationHelp(node)
}
