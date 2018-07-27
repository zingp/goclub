package main

import (
	"fmt"
)

// Node is a representation of binary tree node
type Node struct {
	value int
	left  *Node
	right *Node
}

func gen2Tree() (node *Node) {
	node9 := &Node{value: 9, left: nil, right: nil}
	node8 := &Node{value: 8, left: nil, right: nil}
	node7 := &Node{value: 7, left: nil, right: nil}
	node6 := &Node{value: 6, left: nil, right: nil}
	node5 := &Node{value: 5, left: node8, right: node9}
	node4 := &Node{value: 4, left: nil, right: nil}
	node3 := &Node{value: 3, left: node6, right: node7}
	node2 := &Node{value: 2, left: node4, right: node5}
	node1 := &Node{value: 1, left: node2, right: node3}
	return node1
}

// 广度优先
func breadthFirstSearch(node Node) []int {
	var result []int
	nodes := []Node{node}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]  // nodes 切片缩小，不断向后推移
		result = append(result, node.value)
		if (node.left != nil) {
			nodes = append(nodes, *node.left)
		}
		if (node.right != nil) {
			nodes = append(nodes, *node.right)
		}
	}
	return result
}

// 前序遍历
func preOrderTraverse(node *Node) {

	fmt.Println(node.value)
	if node.left != nil {
		preOrderTraverse(node.left)
	}
	if node.right != nil {
		preOrderTraverse(node.right)
	}
}

func main() {
	root := gen2Tree()
	preOrderTraverse(root)

	array := breadthFirstSearch(*root)
	for _, i := range array {
		fmt.Println(i)
	}
}
