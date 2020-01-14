package main

import "fmt"

type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

func main() {
	node1 := new(Node)
	node1.data = 1
	node2 := new(Node)
	node2.data = 2
	node3 := new(Node)
	node3.data = 3

	node1.left = node2
	node1.right = node3
	node2.right = node1
	node3.left = node1

	fmt.Println(node1)

	//
	user := map[string]interface{}{
		"name": "小明",
		"age":  18,
		"sex":  true,
	}

	fmt.Println(user)
}
