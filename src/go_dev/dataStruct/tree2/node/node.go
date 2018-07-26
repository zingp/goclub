package main

import (
	"fmt"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func testTree() {
	nodeG := Node{data: "g", left: nil, right: nil}
	nodeF := Node{data: "f", left: &nodeG, right: nil}
	nodeE := Node{data: "e", left: nil, right: nil}
	nodeD := Node{data: "d", left: &nodeE, right: nil}
	nodeC := Node{data: "c", left: nil, right: nil}
	nodeB := Node{data: "b", left: &nodeD, right: &nodeF}
	nodeA := Node{data: "a", left: &nodeB, right: &nodeC}

	preOrderRecursive(nodeA)

}

func preOrderRecursive(node Node) {
	fmt.Println(node.data)
	if node.left != nil {
		preOrderRecursive(*node.left)
	}
	// 在这里输出就是中序
	if node.right != nil {
		preOrderRecursive(*node.right)
	}
}

func main() {
	testTree()
}
