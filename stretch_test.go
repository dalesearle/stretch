package main

import (
	"fmt"
	"math"
	"playground/stretch/tree"
	"strconv"
	"testing"
)

func TestStretchRoot(t *testing.T) {
	node := tree.New(1, 31)
	fmt.Println("----------before----------")
	tmp := *node
	printTree(&tmp)
	fmt.Println("----------after-----------")
	stretch(node, 5)
	depth := getTreeDepth(node)
	if depth != 5 {
		t.Fatalf("expected 5, got %d", depth)
	}
	printTree(node)
}

func TestStretch2NodeTree(t *testing.T) {
	parent := tree.New(1, 4)
	parent.AddRightChild(tree.New(0, 8))
	fmt.Println("----------before----------")
	printTree(parent)
	fmt.Println("----------after-----------")
	stretch(parent, 2)
	depth := getTreeDepth(parent)
	if depth != 4 {
		t.Fatalf("expected 4, got %d", depth)
	}
	printTree(parent)
}

func TestStretch4NodeTree(t *testing.T) {
	parent := tree.New(1, 4)
	pr := tree.New(0, 8)
	parent.AddRightChild(pr)
	pr.AddLeftChild(tree.New(0, 6))
	pr.AddRightChild(tree.New(1, 10))
	fmt.Println("----------before----------")
	printTree(parent)
	fmt.Println("----------after-----------")
	stretch(parent, 2)
	depth := getTreeDepth(parent)
	if depth != 6 {
		t.Fatalf("expected 6, got %d", depth)
	}
	printTree(parent)
}

func TestStretchFull4NodeTree(t *testing.T) {
	parent := tree.New(1, 4)
	pl := tree.New(0, 16)
	pr := tree.New(0, 8)
	parent.AddLeftChild(pl)
	parent.AddRightChild(pr)
	pl.AddLeftChild(tree.New(0, 12))
	pl.AddRightChild(tree.New(0, 20))
	pr.AddLeftChild(tree.New(0, 6))
	pr.AddRightChild(tree.New(1, 10))
	fmt.Println("----------before----------")
	printTree(parent)
	fmt.Println("----------after-----------")
	stretch(parent, 2)
	depth := getTreeDepth(parent)
	if depth != 6 {
		t.Fatalf("expected 6, got %d", depth)
	}
	printTree(parent)
}

func getTreeDepth(node *tree.Node) int {
	rval := 0
	if node == nil {
		return rval
	}
	//fmt.Println(node.String())
	depths := []int{
		node.Depth(),
		getTreeDepth(node.LeftChild()),
		getTreeDepth(node.RightChild()),
	}
	for _, depth := range depths {
		if depth > rval {
			rval = depth
		}
	}
	return rval
}

func printTree(node *tree.Node) {
	treeDepth := getTreeDepth(node)
	maxRowWidth := math.Pow(2, float64(treeDepth-1))*2 - 2
	rows := make(map[int][]string)
	fillTree(node, treeDepth)
	mapTree(rows, node)
	for i := 1; i <= len(rows); i++ {
		row := rows[i]
		rowLen := len(row)
		spaces := ((int(maxRowWidth) - rowLen) / rowLen) + 1
		for index, v := range row {
			if index == 0 {
				for s := 0; s < spaces/2; s++ {
					fmt.Print(" ")
				}
			} else {
				for s := 0; s < spaces; s++ {
					fmt.Print(" ")
				}
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func mapTree(rows map[int][]string, node *tree.Node) {
	if node == nil {
		return
	}
	row, ok := rows[node.Depth()]
	if !ok {
		row = make([]string, 0)
	}
	rows[node.Depth()] = append(row, strconv.Itoa(node.Value))
	if node.HasChildren() {
		tmpNode := node.LeftChild()
		if tmpNode == nil {
			tmpNode = tree.New(node.Depth()+1, 0)
		}
		mapTree(rows, tmpNode)
		tmpNode = node.RightChild()
		if tmpNode == nil {
			tmpNode = tree.New(node.Depth()+1, 0)
		}
		mapTree(rows, tmpNode)
	}
}

func maxRowLength(rows map[int][]string) int {
	max := 0
	for _, row := range rows {
		len := len(row)
		if len > max {
			max = len
		}
	}
	return max
}

func fillTree(node *tree.Node, treeDepth int) {
	if node == nil {
		return
	}
	if node.LeftChild() == nil {
		node.AddLeftChild(tree.New(0, 0))
	}
	if node.Depth()+1 < treeDepth {
		fillTree(node.LeftChild(), treeDepth)
	}
	if node.RightChild() == nil {
		node.AddRightChild(tree.New(0, 0))
	}
	if node.Depth()+1 < treeDepth {
		fillTree(node.RightChild(), treeDepth)
	}
}
