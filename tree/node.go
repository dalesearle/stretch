package tree

import (
	"fmt"
	"strconv"
)

const (
	Left  = 1
	Right = 2
)

type Node struct {
	depth       int
	left        *Node
	orientation int
	parent      *Node
	right       *Node
	Value       int
}

func New(depth, value int) *Node {
	return &Node{
		depth: depth,
		Value: value,
	}
}

func (n *Node) AddChild(child *Node, orientation int) {
	if orientation == Left {
		n.AddLeftChild(child)
	} else {
		n.AddRightChild(child)
	}
}

func (n *Node) AddLeftChild(child *Node) {
	child.depth = n.depth + 1
	child.parent = n
	child.orientation = Left
	n.left = child
}

func (n *Node) AddRightChild(child *Node) {
	child.depth = n.depth + 1
	child.parent = n
	child.orientation = Right
	n.right = child
}

func (n *Node) RightChild() *Node {
	return n.right
}

func (n *Node) LeftChild() *Node {
	return n.left
}

func (n *Node) Depth() int {
	return n.depth
}

func (n *Node) HasChildren() bool {
	return n.left != nil || n.right != nil
}

func (n *Node) Orientation() int {
	if n.parent == nil {
		return Left
	}
	return n.orientation
}

func (n *Node) OrphanChildren() {
	n.left = nil
	n.right = nil
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) String() string {
	left := ""
	if n.left != nil {
		left = strconv.Itoa(n.left.Value) + "<-"
	}
	right := ""
	if n.right != nil {
		right = "->" + strconv.Itoa(n.right.Value)
	}
	return fmt.Sprintf("%s%d%s", left, n.Value, right)
}
