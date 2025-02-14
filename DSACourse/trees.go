package main

import (
	"fmt"
	"math"
)

type StackInt struct {
	data []int
}

func (s *StackInt) Length() int {
	return len(s.data)
}
func (s *StackInt) IsEmpty() bool {
	return s.Length() == 0
}

func (s *StackInt) Print() {
	length := len(s.data)
	for i := 0; i < length; i++ {
		fmt.Print(s.data[i], " ")
	}
}

func (s *StackInt) Push(value int) {
	s.data = append(s.data, value)
}

func (s *StackInt) Pop() int {
	if s.IsEmpty() {
		fmt.Println("The stack is empty.")
		return -9999
	} else {
		topind := s.Length() - 1
		element := s.data[topind]
		s.data = s.data[:topind]

		return element
	}
}

func (s *StackInt) Top() int {
	if s.IsEmpty() {
		fmt.Println("The stack is empty.")
		return -9999
	} else {
		topind := s.Length() - 1
		element := s.data[topind]
		return element
	}
}

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

type Tree struct {
	root *treeNode
}

func PrintInOrder(root *treeNode) {
	if root != nil {
		PrintInOrder(root.left)
		fmt.Print(root.value, " ")
		PrintInOrder(root.right)
	}

}

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *treeNode {
	if start >= size {
		return nil
	}

	root := new(treeNode)
	root.value = arr[start]
	root.left = levelOrderBinaryTree(arr, 2*start+1, size)
	root.right = levelOrderBinaryTree(arr, 2*start+2, size)

	return root

}

func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
}

func printPreOrder(n *treeNode) {

	//Implement your solution here

	if n != nil {
		fmt.Print(n.value, " ")
		printPreOrder(n.left)
		printPreOrder(n.right)
	}

}

func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
}

func printPostOrder(n *treeNode) {

	//Implement your solution here
	if n != nil {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Print(n.value, " ")

	}

	//Implement your solution here

}

// func (t *Tree) PrintBreadthFirst() {

// 	//Implement your solution here
// 	que := Queue{}

// 	if t.root != nil {
// 		que.Enqueue(t.root)
// 	}

// 	for que.Length() > 0 {
// 		temp := que.Dequeue()
// 		temp2 := temp.(*treeNode)
// 		fmt.Print(temp)
// 		if temp2.left != nil {
// 			que.Enqueue(temp2.left)
// 		}
// 		if temp2.right != nil {
// 			que.Enqueue(temp2.right)
// 		}

// 	}

// 	for que.Length() > 0 {
// 		fmt.Print(que.Dequeue())
// 	}

// }

type node struct {
	nvalue *treeNode
	next   *node
}

type DepthStack struct {
	head *node
	size int
}

func (s *DepthStack) Length() int {
	return s.size
}

func (s *DepthStack) IsEmpty() bool {
	return s.size == 0
}

func (s *DepthStack) Peek() (*treeNode, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil, false
	}
	return s.head.nvalue, true
}

func (s *DepthStack) Push(nvalue *treeNode) {
	s.head = &node{&treeNode{nvalue.value, nvalue.left, nvalue.right}, s.head}
	s.size++
}

func (s *DepthStack) Pop() *treeNode {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil
	}
	temp := new(treeNode)
	temp.value = s.head.nvalue.value
	temp.left = s.head.nvalue.left
	temp.right = s.head.nvalue.right
	s.head = s.head.next
	s.size--
	return temp
}

func (t *Tree) PrintDepthFirst() {

	//Implement your solution here

	current := new(treeNode)

	current = t.root

	stk := new(DepthStack)

	for {
		if current != nil {
			fmt.Print(current.value, " ")
			if current.right != nil {
				stk.Push(current.right)
			}
			current = current.left
		} else {
			if stk.Length() != 0 {
				current = stk.Pop()
			} else {
				break
			}
		}

	}
}

func (t *Tree) PrintSpiralTree() {
	//Implement your solution here

	stk1 := DepthStack{}
	stk2 := DepthStack{}

	if t.root != nil {
		stk1.Push(t.root)
	}

	for stk1.Length() > 0 || stk1.Length() > 0 {
		for stk1.Length() > 0 {
			temp := stk1.Pop()
			fmt.Print(temp.value, " ")

			if temp.left != nil {
				stk2.Push(temp.left)
			}
			if temp.right != nil {
				stk2.Push(temp.right)
			}

		}
		fmt.Print("; ")
		for stk2.Length() > 0 {
			temp := stk2.Pop()
			fmt.Print(temp.value, " ")

			if temp.right != nil {
				stk1.Push(temp.right)
			}
			if temp.left != nil {
				stk1.Push(temp.left)
			}

		}
		fmt.Print("; ")

	}

	//Kindly change this line of code as per your needs
	fmt.Print("Sample Output")

	//Implement your solution here
}

// func (t *Tree) PrintLevelOrderLineByLine() {
// 	que1 := Queue{}
// 	que2 := Queue{}

// 	if t.root != nil {
// 		que1.Enqueue(t.root)
// 	}
// 	fmt.Print(t.root.value, " ; ")
// 	for que1.Length() > 0 {
// 		temp := que1.Dequeue()
// 		temp2 := temp.(*treeNode)
// 		if temp2.left != nil {
// 			que2.Enqueue(temp2.left)
// 		}

// 		if temp2.right != nil {
// 			que2.Enqueue(temp2.right)
// 		}

// 		if que1.Length() != 0 {
// 			continue
// 		}

// 		for que2.Length() > 0 {
// 			temp := que2.Dequeue()
// 			temp2 := temp.(*treeNode)
// 			fmt.Print(temp2.value, " ")
// 			que1.Enqueue(temp2)

// 		}
// 		fmt.Print("; ")

// 	}
// }

// func (t *Tree) PrintLevelOrderLineByLine2() {
// 	//Implement your solution here

// 	que := Queue{}

// 	var count int

// 	if t.root != nil {
// 		que.Enqueue(t.root)
// 	}

// 	for que.Length() > 0 {
// 		count = que.Length()
// 		for count > 0 {
// 			temp := que.Dequeue()
// 			temp2 := temp.(*treeNode)

// 			fmt.Print(temp2.value, " ")
// 			if temp2.left != nil {
// 				que.Enqueue(temp2.left)
// 			}

// 			if temp2.right != nil {
// 				que.Enqueue(temp2.right)
// 			}

// 			count--
// 		}
// 		fmt.Print("; ")

// 	}

// 	//Kindly change this line of code as per your needs
// 	fmt.Print("Sample Output")

// 	//Implement your solution here
// }

func (t *Tree) PrintDepthFirst2() {
	stk := new(DepthStack)

	if t.root != nil {
		stk.Push(t.root)
	}
	for stk.Length() != 0 {
		temp := stk.Pop()
		fmt.Print(temp.value, " ")
		if temp.right != nil {
			stk.Push(temp.right)
		}
		if temp.left != nil {
			stk.Push(temp.left)
		}
	}
}

func (t *Tree) NthPreOrder(index int) {
	var counter int = 0
	nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *treeNode, index int, counter *int) {
	//Implement your solution here

	if node != nil {
		(*counter)++
		//fmt.Print(*counter)
		if *counter == index {
			fmt.Print(node.value)
			return
		}
		nthPreOrder(node.left, index, counter)
		nthPreOrder(node.right, index, counter)
	}
}

func (t *Tree) NthPostOrder(index int) {
	var counter int = 0
	nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *treeNode, index int, counter *int) {
	//Implement your solution here

	if node != nil {

		nthPostOrder(node.left, index, counter)

		nthPostOrder(node.right, index, counter)
		(*counter)++

		if *counter == index {
			fmt.Print(node.value)
			return
		}

	}

	//Implement your solution here
}

func (t *Tree) NthInOrder(index int) {
	var counter int = 0
	nthInOrder(t.root, index, &counter)
}

func nthInOrder(node *treeNode, index int, counter *int) {
	//Implement your solution here

	if node != nil {
		nthInOrder(node.left, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Print(node.value)
			return
		}

		nthInOrder(node.right, index, counter)

	}

	//Implement your solution here
}

func (t *Tree) PrintAllPath() {
	stk := new(StackInt)
	printAllPath(t.root, stk)
}

func printAllPath(curr *treeNode, stk *StackInt) {

	stk.Print()
	fmt.Println()

	if curr == nil {
		return
	}
	stk.Push(curr.value)
	if curr.right == nil && curr.left == nil {
		stk.Print()
		fmt.Print("; ")
		stk.Pop()
		return

	}
	printAllPath(curr.right, stk)
	printAllPath(curr.left, stk)
	stk.Pop()
}

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numNodes(curr *treeNode) int {
	//Implement your solution here
	if curr != nil {
		return (1 + numNodes(curr.right) + numNodes(curr.left))
	} else {
		return 0
	}
}

func (t *Tree) SumAllBT() int {
	return sumAllBT(t.root)
}

func sumAllBT(curr *treeNode) int {
	//Implement your solution here

	if curr != nil {
		return (curr.value + sumAllBT(curr.right) + sumAllBT((curr.left)))
	} else {
		return 0

	}

}

func (t *Tree) NumLeafNodes() int {
	return numLeafNodes(t.root)
}

func numLeafNodes(curr *treeNode) int {
	//Implement your solution here

	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	} else {
		return (0 + numLeafNodes(curr.left) + numLeafNodes(curr.right))
	}
}

func (t *Tree) NumFullNodesBT() int {
	return numFullNodesBT(t.root)
}

func numFullNodesBT(curr *treeNode) int {

	//Implement your solution here

	var count int
	if curr == nil {
		return 0
	}

	count = numFullNodesBT(curr.right) + numFullNodesBT(curr.left)
	if curr.left != nil && curr.right != nil {
		count++
	}

	return count
}

func (t *Tree) SearchBT(value int) bool {
	return searchBT(t.root, value)
}

func searchBT(root *treeNode, value int) bool {

	var left bool
	var right bool
	//Implement your solution here
	if root == nil {
		return false
	}

	if root.value == value {
		return true
	}

	left = searchBT(root.left, value)
	right = searchBT(root.right, value)

	if right {
		return true
	}
	if left {
		return true
	}
	return false

}

func (t *Tree) FindMaxBT() int {
	return findMaxBT(t.root)
}

func findMaxBT(curr *treeNode) int {
	var max int
	var left int
	var right int

	if curr == nil {
		return math.MinInt32
	}
	max = curr.value
	left = findMaxBT(curr.left)
	right = findMaxBT(curr.right)

	if left > max {
		max = left
	}

	if right > max {
		max = right
	}

	return max
}

func (t *Tree) TreeDepth() int {
	return treeDepth(t.root)
}

func treeDepth(root *treeNode) int {

	//Implement your solution here

	if root == nil {
		return -1
	}

	left := treeDepth(root.left)
	right := treeDepth(root.right)

	if left > right {
		return left + 1
	}

	return right + 1 //Kindly change the return value as per needs
}

func (t *Tree) IsEqual(t2 *Tree) bool {
	return isEqual(t.root, t2.root)
}

func isEqual(node1 *treeNode, node2 *treeNode) bool {
	//Implement your solution here

	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return ((node1.value == node2.value) && isEqual(node1.left, node2.left) && isEqual(node1.right, node2.right))
	}
}

func (t *Tree) CopyTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = copyTree(t.root)
	return Tree2
}

func copyTree(curr *treeNode) *treeNode {
	temp := new(treeNode)

	if curr != nil {
		temp.value = curr.value
		temp.left = copyTree(curr.left)
		temp.right = copyTree(curr.right)
		return temp
	}
	return nil
}

func (t *Tree) CopyMirrorTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = copyTree(t.root)
	return Tree2
}

func copyMirrorTree(curr *treeNode) *treeNode {
	temp := new(treeNode)

	if curr != nil {
		temp.value = curr.value
		temp.left = copyTree(curr.right)
		temp.right = copyTree(curr.left)
		return temp
	}
	return nil
}

func (t *Tree) Free() {
	//Implement your solution here
	t.root = nil

}

func (t *Tree) IsCompleteTree() bool {
	count := t.NumNodes()
	return isCompleteTreeUtil(t.root, 0, count)
}

func isCompleteTreeUtil(curr *treeNode, index int, count int) bool {
	//Implement your solution here

	if curr == nil {
		return true
	}
	if index > count {
		return false
	}

	return isCompleteTreeUtil(curr.left, 2*index+1, count) && isCompleteTreeUtil(curr.right, 2*index+2, count)
}

func (t *Tree) IsHeap() bool {
	parentValue := -99999999
	return t.IsCompleteTree() && isHeapUtil(t.root, parentValue)
}

func isHeapUtil(curr *treeNode, parentValue int) bool {
	//Implement your solution here

	if curr == nil {
		return true
	}

	if curr.left.value > parentValue || curr.right.value > parentValue {
		return false
	}
	return isHeapUtil(curr.left, curr.left.value) && isHeapUtil(curr.right, curr.right.value)
}
func maxf(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *Tree) TreeToListRec() *treeNode {
	return treeToListRec(t.root)
}

func treeToListRec(curr *treeNode) *treeNode {
	//Implement your solution here

	if curr == nil {
		return nil
	}

	var head, tail, temphead *treeNode

	if curr.left == nil && curr.right == nil {
		curr.left = curr
		curr.right = curr
		return curr
	}

	if curr.left != nil {
		head = treeToListRec(curr.left)
		tail = head.left
		curr.left = tail
		tail.right = curr
	} else {
		head = curr
	}

	if curr.right != nil {
		temphead = treeToListRec(curr.right)
		tail = temphead.left
		curr.right = tail
		temphead.left = curr
	} else {
		tail = curr
	}

	head.left = tail
	tail.right = head

	return head

}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	tree := LevelOrderBinaryTree(arr)

// 	//PrintInOrder(tree.root)
// 	// fmt.Println()

// 	// printPreOrder(tree.root)
// 	// fmt.Println()

// 	//printPostOrder(tree.root)

// 	// //tree.PrintBreadthFirst()
// 	//fmt.Print(tree.root.value)
// 	// tree.PrintDepthFirst2()

// 	// tree.PrintSpiralTree()

// 	// tree.PrintPostOrder()

// 	// tree.NthPostOrder(5)

// 	// tree.NthInOrder(5)

// 	//tree.PrintAllPath()

// 	fmt.Print(tree.NumFullNodesBT())

// }
