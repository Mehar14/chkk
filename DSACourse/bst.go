package main

import "fmt"

type bstNode struct {
	value       int
	left, right *bstNode
}

type BSTree struct {
	root *bstNode
}

func CreateBinarySearchTree(arr []int) *BSTree {
	t := new(BSTree)
	size := len(arr)
	t.root = createBinarySearchTreeUtil(arr, 0, size-1)
	return t
}

func createBinarySearchTreeUtil(arr []int, start int, end int) *bstNode {

	if start > end {
		return nil
	}
	mid := (start + end) / 2
	curr := new(bstNode)
	curr.value = arr[mid]

	curr.left = createBinarySearchTreeUtil(arr, start, mid-1)
	curr.right = createBinarySearchTreeUtil(arr, mid+1, end)

	return curr
}

func (t *BSTree) Add(value int) {
	t.root = addNode(t.root, value)
}

func addNode(n *bstNode, value int) *bstNode {
	//Implement your solution here

	if n == nil {
		n = new(bstNode)
		n.value = value
		return n
	}
	if value < n.value {
		n.left = addNode(n.left, value)

	} else if value > n.value {
		n.right = addNode(n.right, value)
	}
	return n
}

func (t *BSTree) Find(value int) bool {
	//Implement your solution here

	var curr *bstNode
	curr = t.root

	for {
		if curr == nil {
			break
		} else if value > curr.value {
			curr = curr.right
		} else if value < curr.value {
			curr = curr.left
		} else {
			return true
		}
	}

	return false
}

func (t *BSTree) FindMinNode() *bstNode {

	var node *bstNode = t.root

	for node.left != nil {
		node = node.left
	}
	//Implement your solution here

	return node
}

// func IsBST(root *treeNode) bool {
// 	//Implement your solution here

// 	if root == nil {
// 		return true
// 	}

// 	if root.left != nil && findMaxNode(root.left).value > root.value {
// 		return false
// 	}

// 	if root.right != nil && findMinMode(root.right).value <= root.value {
// 		return false
// 	}

// 	return (IsBST(root.left) && IsBST(root.right))
// }

// func (t *BSTree) DeleteNode(value int) {
// 	t.root = DeleteNode(t.root, value)
// }

// func DeleteNode(node *bstNode, value int) *bstNode {
// 	//Implement your solution here

// 	//temp = new(bstNode)

// 	if node != nil {
// 		if node.value == value {
// 			if node.left == nil && node.right == nil {
// 				return nil //leaf node
// 			}
// 			if node.left == nil {
// 				//no left node
// 				temp = node.right
// 				return temp
// 			}
// 			if node.right == nil {
// 				temp = node.left
// 				return temp
// 			}

// 			//maxnode := FindMax(node.left)
// 			maxval := maxnode.value
// 			node.value = maxval
// 			node.left = DeleteNode(node.left, maxval)

// 		} else {
// 			if node.value > value {
// 				node.left = DeleteNode(node.left, value)
// 			} else {
// 				node.right = DeleteNode(node.right, value)
// 			}
// 		}

// 	}
// 	return node
// }

func (t *BSTree) PrintDataInRange(min int, max int) {
	printDataInRange(t.root, min, max)
}

func printDataInRange(root *bstNode, min int, max int) {

	//Implement your solution here

	if root != nil {
		if root.value >= min && root.value <= max {
			fmt.Print(root.value, " ")
		}
		printDataInRange(root.left, min, max)
		printDataInRange(root.right, min, max)
	}

	//Implement your solution here
}
