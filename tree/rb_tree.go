package tree

import (
	"fmt"
	"structure/list"
)

type RBTree struct {
	root    *RBNode
	compare func(*interface{}, *interface{}) (int)
}

type RBNode struct {
	color   color
	Data    interface{}
	l_child *RBNode
	r_child *RBNode
	parent  *RBNode
}

type color bool

const (
	red   color = true
	black color = false
)

type position bool

const (
	left  position = true
	right position = false
)

func NewRB(compare func(*interface{}, *interface{}) (int)) *RBTree {
	return &RBTree{root: nil, compare: compare}
}

func (this *RBTree) Search(data interface{}) *RBNode {
	return this.searchHelp(this.root, &data)
}

func (this *RBTree) searchHelp(node *RBNode, data *interface{}) *RBNode {
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

/*
1：如果是空树，直接插入，设为黑节点，结束
2：如果已插入，结束
3：插入红色新节点
	3.1：如果是根节点，设置为黑色，结束
	3.2：如果父节点是黑色，结束
	3.3：父节点为红色：
		3.3.1 如果叔叔节点是红色，把父亲节点和叔叔节点设置为黑色，祖父节点为红色，祖父节点为新的插入节点->3
		3.3.2 如果叔叔节点是黑色或者不存在
			3.3.2.1 如果父节点是左节点，新节点是左节点 设置祖父节点为红色，父亲节点为黑色，对祖父节点右旋
			3.3.2.2 如果父节点是左节点，新节点是右节点 对父亲节点左旋->6.2.1
			3.3.2.3 如果父节点是右节点，新节点是左节点 对父亲节点右旋->6.2.4
			3.3.2.4 如果父节点是右节点，新节点是右节点 设置祖父节点为红色，父亲节点为黑色，对祖父节点左旋
*/

func (this *RBTree) Insert(data interface{}) {
	//如果为空树，直接插入
	if this.root == nil {
		this.root = &RBNode{Data: data, color: black}
		return
	}
	cur := this.root
	var cmp int
	for {
		cmp = this.compare(&cur.Data, &data)
		if cmp < 0 {
			if cur.l_child == nil {
				break
			}
			cur = cur.l_child
		} else if cmp > 0 {
			if cur.r_child == nil {
				break
			}
			cur = cur.r_child

		} else {
			//如果已存在，忽略
			return
		}
	}
	new_node := RBNode{parent: cur, Data: data, color: red}
	if cmp < 0 {
		cur.l_child = &new_node
	} else {
		cur.r_child = &new_node
	}
	cur = &new_node
	for {
		//如果是根节点，设为黑，结束
		parent := cur.parent
		if parent == nil {
			cur.color = black
			return
		}
		//如果父节点是黑色，结束
		if parent.color == black {
			return
		}

		var parent_left bool
		var uncle *RBNode
		if parent.parent.l_child == parent {
			uncle = parent.parent.r_child
			parent_left = true
		} else {
			uncle = parent.parent.l_child
			parent_left = false
		}
		if uncle != nil && uncle.color == red {
			//6.1
			parent.color = black
			uncle.color = black
			parent.parent.color = red
			cur = parent.parent
			continue
		}
		var cur_left bool
		if parent.l_child == cur {
			cur_left = true
		} else {
			cur_left = false
		}
		if parent_left {
			if !cur_left {
				//6.2.2
				parent = parent.leftRotation()
			}
			//6.2.1
			gf := parent.parent.rightRotation()
			if gf.parent == nil {
				this.root = gf
			}
		} else {
			if cur_left {
				//6.2.3
				parent = parent.rightRotation()
			}
			//6.2.4
			gf := parent.parent.leftRotation()
			if gf.parent == nil {
				this.root = gf
			}
		}
		return
	}
}

func (this *RBTree) Check() {
	f := func(node *RBNode) {
		if node.l_child != nil {
			if node.l_child.parent != node {
				fmt.Println("left error:", node.Data, node.l_child.Data, node.l_child.parent.Data)
			}
			if node.r_child.parent != node {
				fmt.Println("right error:", node.Data, node.r_child.Data, node.r_child.parent.Data)
			}
		}
	}
	this.InTraverse(f)
}

/*
(空树或者未找到直接结束)
1: 如果节点有两个子节点，进行替换->2or3
2: 如果节点有一个子节点，这个子节点必定是红色，进行替换删除，结束
3: 如果节点没有子节点
	3.1 如果节点是红色，直接删除，结束
	3.2 如果节点是黑色，如果不是根节点，则必然有兄弟节点
		3.2.1 如果是根节点，直接删除，结束
		3.2.2 如果兄弟节点是红色
			3.2.2.1 节点在左，交换兄弟节点和父节点颜色，对父节点左旋->3.2.3
			3.2.2.2 节点在右，交换兄弟节点和父节点颜色，对父节点右旋->3.2.3
		3.2.3 如果兄弟节点是黑色
			3.2.3.1 如果兄弟节点没有红色节点，删除节点，将兄弟节点设置为红色，父节点设置为黑色，结束
			3.2.3.2 如果兄弟节点有红色节点
				3.2.3.2.1 兄弟节点在左，兄弟节点左节点存在，删除节点，交换兄弟节点和父节点颜色，兄弟节点左节点设置黑色，对父亲节点右旋，结束
				3.2.3.2.2 兄弟节点在左，兄弟节点左节点不存在，交换兄弟节点和兄弟节点右节点颜色，对兄弟节点进行左旋->3.2.3.2.1
				3.2.3.2.3 兄弟节点在右，兄弟节点右节点存在，删除节点，交换兄弟节点和父节点颜色，兄弟节点右节点设置黑色，对父亲节点左旋，结束
				3.2.3.2.4 兄弟节点在右，兄弟节点右节点不存在，交换兄弟节点和兄弟节点左节点颜色，对兄弟节点进行右旋->3.2.3.2.3
*/

func (this *RBTree) Delete(data interface{}) {
	if this.root == nil {
		return
	}
	cur := this.root
	var cmp int
	for {
		if cur == nil {
			return
		}
		cmp = this.compare(&cur.Data, &data)
		if cmp < 0 {
			cur = cur.l_child
		} else if cmp > 0 {
			cur = cur.r_child
		} else {
			break
		}
	}
	if cur.l_child != nil && cur.r_child != nil {
		l_max := (*cur).l_child.findMax()
		cur.Data = l_max.Data
		cur = l_max
	}
	if cur.l_child != nil {
		cur.Data = cur.l_child.Data
		cur.l_child = nil
		return
	} else if cur.r_child != nil {
		cur.Data = cur.r_child.Data
		cur.r_child = nil
		return
	}
	if cur.color == red {
		del_node_help(cur)
		return
	}
	if cur == this.root {
		this.root = nil
		return
	}
	var sibling_pos position
	var sibling *RBNode
	sibling_pos, sibling = get_sibling(cur)

	if sibling.color == red {
		var parent *RBNode
		if sibling_pos == left {
			parent = cur.parent.rightRotation()
		} else {
			parent = cur.parent.leftRotation()
		}
		if parent.parent == nil {
			this.root = parent
		}
		_, sibling = get_sibling(cur)

	}
	parent := cur.parent
	if sibling.l_child == nil && sibling.r_child == nil {
		parent.color = black
		sibling.color = red
		if sibling_pos == left {
			parent.r_child = nil
		} else {
			parent.l_child = nil
		}
		return
	}
	if sibling_pos == left {
		if sibling.l_child == nil {
			sibling = sibling.leftRotation()
		}
		parent.r_child = nil
		sibling.l_child.color = black
		parent = parent.rightRotation()
	} else {
		if sibling.r_child == nil {
			sibling = sibling.rightRotation()
		}
		parent.l_child = nil
		sibling.r_child.color = black
		parent = parent.leftRotation()
	}
	if parent.parent == nil {
		this.root = parent
	}
	return
}

func (this *RBTree) IntraverseNonRecur(visit func(*interface{})) {
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
					cur = data.(*RBNode)
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

func (this *RBTree) InTraverse(visit func(node *RBNode)) {
	rb_inTraverseHelp(this.root, visit)
}

func rb_inTraverseHelp(node *RBNode, visit func(node *RBNode)) {
	if node == nil {
		return
	}
	rb_inTraverseHelp(node.l_child, visit)
	visit(node)
	rb_inTraverseHelp(node.r_child, visit)
}

//返回旋转后此的父节点
//必须有右节点
func (this *RBNode) leftRotation() (*RBNode) {
	r_child := this.r_child
	this.color, r_child.color = r_child.color, this.color
	this.r_child = r_child.l_child
	if this.r_child != nil {
		this.r_child.parent = this
	}
	r_child.l_child = this
	r_child.parent = this.parent
	this.parent = r_child

	if r_child.parent != nil {
		if r_child.parent.l_child == this {
			r_child.parent.l_child = r_child
		} else {
			r_child.parent.r_child = r_child
		}
	}
	return r_child
}

//必须有左节点
func (this *RBNode) rightRotation() (*RBNode) {
	l_child := this.l_child
	this.color, l_child.color = l_child.color, this.color
	this.l_child = l_child.r_child
	if this.l_child != nil {
		this.l_child.parent = this
	}
	l_child.r_child = this
	l_child.parent = this.parent
	this.parent = l_child
	if l_child.parent != nil {
		if l_child.parent.l_child == this {
			l_child.parent.l_child = l_child
		} else {
			l_child.parent.r_child = l_child
		}
	}
	return l_child
}

func del_node_help(cur *RBNode) {
	if cur.parent.l_child == cur {
		cur.parent.l_child = nil
	} else {
		cur.parent.r_child = nil
	}
}

func get_sibling(cur *RBNode) (position, *RBNode) {
	if cur.parent.l_child == cur {
		return right, cur.parent.r_child
	} else {
		return left, cur.parent.l_child
	}
}

func (this *RBTree) deleteHelp(parent **RBNode, node **RBNode, data *interface{}) {
	if *node == nil {
		return
	}
	cmp := this.compare(&(*node).Data, data)
	if cmp < 0 {
		this.deleteHelp(node, &(*node).l_child, data)
		return
	} else if cmp > 0 {
		this.deleteHelp(node, &(*node).r_child, data)
		return
	}
	if (*node).l_child != nil && (*node).r_child != nil {
		//有两个子节点
		l_max := (*node).l_child.findMax()
		(*node).Data = l_max.Data
		this.deleteHelp(node, &(*node).l_child, &l_max.Data)
		return
	} else if (*node).l_child != nil {
		//只有一个子节点
		//结束
		(*node).l_child.parent = (*node).parent
		*node = (*node).l_child
		(*node).color = black
		return
	} else if (*node).r_child != nil {
		//只有一个子节点
		//结束
		(*node).r_child.parent = (*node).parent
		*node = (*node).r_child
		(*node).color = black
		return
	}
	if (*node).color == red {
		//节点为红
		//结束
		*node = nil
		return
	}
	if (*parent) == nil {
		//节点为根节点
		//结束
		(*node) = nil
		return
	}
	var sibling_is_left bool
	var sibling *RBNode
	if (*parent).l_child == *node {
		sibling_is_left = false
		sibling = (*parent).r_child
	} else {
		sibling_is_left = true
		sibling = (*parent).l_child
	}
	if sibling.color == black {
		//兄弟节点为黑节点
		if sibling.l_child == nil && sibling.r_child == nil {
			//兄弟节点无红子节点
			//结束
			(*node) = nil
			(*parent).color = black
			sibling.color = red
		} else {
			//兄弟节点有红子节点
			//结束
			(*node) = nil
			if sibling_is_left {
				(*parent).l_child = sibling.r_child
				sibling.r_child = *parent
				if sibling.l_child != nil {
					sibling.l_child.color = black
				}
			} else {
				(*parent).r_child = sibling.l_child
				sibling.l_child = *parent
				if sibling.r_child != nil {
					sibling.r_child.color = black
				}
			}
			sibling.color = (*parent).color
			(*parent).color = black
			(*parent) = sibling
		}
		return
	}
	//如果兄弟节点是红色
	(*parent).color = red
	sibling.color = black
	var tmp_parent **RBNode
	if sibling_is_left {
		(*parent).l_child = sibling.r_child
		sibling.r_child = (*parent)
		tmp_parent = &sibling.r_child
		node = &(*parent).r_child
	} else {
		(*parent).r_child = sibling.l_child
		sibling.l_child = (*parent)
		tmp_parent = &sibling.l_child
		node = &(*parent).l_child
	}
	(*parent) = sibling
	this.deleteHelp(tmp_parent, node, data)
	return
}

func (this *RBNode) findMax() *RBNode {
	cur := this
	for (cur.r_child != nil) {
		cur = cur.r_child
	}
	return cur
}
