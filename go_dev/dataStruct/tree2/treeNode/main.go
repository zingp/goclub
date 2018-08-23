package main

import (
	"fmt"
)

// Node is a representation of binary tree node
type Node struct {
	value interface{}
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
func breadthFirstSearch(node Node) []interface{} {
	var result []interface{}
	nodes := []Node{node}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:] // nodes 切片缩小，不断向后推移
		result = append(result, node.value)
		if node.left != nil {
			nodes = append(nodes, *node.left)
		}
		if node.right != nil {
			nodes = append(nodes, *node.right)
		}
	}
	return result
}

// 前序遍历  根->左子树->右子树
func preOrderTraversal(node *Node) {
	fmt.Printf(" %v", node.value)
	if node.left != nil {
		preOrderTraversal(node.left)
	}
	if node.right != nil {
		preOrderTraversal(node.right)
	}
}

// 中序遍历 左子树 根 右子树
func midOrderTraversal(node *Node) {
	if node.left != nil {
		midOrderTraversal(node.left)
	}
	fmt.Printf(" %v", node.value)
	if node.right != nil {
		midOrderTraversal(node.right)
	}
}

// 后序遍历
func postOrderTraversal(node *Node) {
	if node.left != nil {
		postOrderTraversal(node.left)
	}
	if node.right != nil {
		postOrderTraversal(node.right)
	}
	fmt.Printf(" %v", node.value)
}

/*  // 非递归遍历有问题
func inOderTravelsal1(node *Node) {
	n := node
	stack := NewStack()
	for n != nil || stack.Size() != 0 {
		for n != nil {
			stack.Push(n)
			n = n.left
		}
		if stack.Size() != 0 {
			t := stack.Pop()
			fmt.Printf(" %v", t.(*Node))
			n = n.right
		}
	}
}
*/

func main() {
	root := gen2Tree()
	preOrderTraversal(root)
	fmt.Println()
	midOrderTraversal(root)
	fmt.Println()
	postOrderTraversal(root)
	fmt.Println()

	array := breadthFirstSearch(*root)
	fmt.Println(array)
}
