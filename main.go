package main

import (
	"playground/stretch/tree"
)

func main() {
}

func stretch(node *tree.Node, factor int) {
	lChild := node.LeftChild()
	rChild := node.RightChild()
	node.OrphanChildren()
	newValue := node.Value / factor
	node.Value = newValue
	tmpNode := node
	for i := 1; i < factor; i++ {
		child := tree.New(0, newValue)
		tmpNode.AddChild(child, tmpNode.Orientation())
		tmpNode = child
	}
	if lChild != nil {
		tmpNode.AddLeftChild(lChild)
		stretch(lChild, factor)
	}
	if rChild != nil {
		tmpNode.AddRightChild(rChild)
		stretch(rChild, factor)
	}
}
