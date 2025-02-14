package main

import "fmt"

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]interface{}
	front int
	back  int
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Length() int {
	return q.size
}

func (q *Queue) Enqueue(value interface{}) {

	//check if capacity reached
	if q.size == capacity {
		return
	} else {
		q.data[q.back] = value
		q.size++

		q.back = (q.back + 1) % (capacity - 1) //act as a circular queue
	}
}

func (q *Queue) Dequeue() interface{} {

	var element interface{}
	if q.Length() == 0 {
		return -9999 //if empty
	} else {
		q.size--
		element = q.data[q.front]
		q.front = (q.front + 1) % (capacity - 1)
		return element
	}
}

type QueueNode struct {
	value int
	next  *QueueNode
}
type QueueLinkedList struct {
	head *QueueNode
	tail *QueueNode
	size int
}

func (q *QueueLinkedList) Enqueue(root *treeNode) {
	panic("unimplemented")
}

func (q *QueueLinkedList) Length() int {
	return q.size
}

func (q *QueueLinkedList) isEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Peek() int {
	if q.isEmpty() {
		fmt.Println("QueueEmptyException")
		return -999
	}

	return q.head.value
}

func (q *QueueLinkedList) Print() {
	if q.isEmpty() {
		fmt.Println("QueueEmptyException")
		return
	}
	temp := q.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (q *QueueLinkedList) Add(value int) {
	newnode := &QueueNode{}
	//check if empty

	newnode.value = value

	if q.isEmpty() {
		q.head = newnode
		q.tail = newnode
	} else {
		q.tail.next = newnode
		q.tail = newnode
	}

	q.size++

}

func (q *QueueLinkedList) Remove() int {
	if q.isEmpty() {
		fmt.Println("QueueEmptyException")
		return -999
	}
	temp := q.head.value
	q.head = q.head.next

	q.size--

	return temp
}

func (q *QueueLinkedList) Addcircular(value int) {
	newnode := &QueueNode{}
	newnode.value = value
	if q.head == nil { //queue is empty
		q.head = newnode
		q.tail = newnode

		q.tail.next = q.head

	} else {
		q.tail.next = newnode
		newnode.next = q.head
		q.tail = newnode
	}
	q.size++
}

func (q *QueueLinkedList) DeleteCircular() int {
	if q.isEmpty() {
		fmt.Println("Empty queue exception")
		return -9999
	}
	element := q.head.value
	if q.head == q.tail { //one node
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.tail.next = q.head
	}
	q.size--
	return element

}

func (q *QueueLinkedList) PrintCircular() {
	if q.isEmpty() {
		return
	}

	current := q.head

	for {
		fmt.Println(current.value)
		current = current.next

		if current == q.head {
			break
		}
	}
}

type DequeNode struct {
	value int
	next  *DequeNode
	prev  *DequeNode
}

type Dequeue struct {
	size int
	head *DequeNode
	tail *DequeNode
}

func (d *Dequeue) isEmpty() bool {
	return d.size == 0
}

func (d *Dequeue) length() int {
	return d.size
}

func (d *Dequeue) PeekFront() int {
	if d.size == 0 {
		return -9999
	}
	return d.head.value
}

func (d *Dequeue) PeekBack() int {
	if d.size == 0 {
		return -9999
	}
	return d.tail.value
}

func (d *Dequeue) AddFront(value int) {
	newnode := &DequeNode{}
	newnode.value = value

	if d.size == 0 { //empty
		d.head = newnode
		d.tail = newnode
	} else {
		d.head.prev = newnode
		newnode.next = d.head
		d.head = newnode
	}
	d.size++
}

func (d *Dequeue) AddBack(value int) {
	newnode := &DequeNode{}
	newnode.value = value

	if d.size == 0 { //empty
		d.head = newnode
		d.tail = newnode
	} else {
		d.tail.next = newnode
		newnode.prev = d.tail
		d.tail = newnode
	}
	d.size++
}

func (d *Dequeue) DeleteBack() int {
	if d.size == 0 {
		return -9999
	}
	element := d.tail.value
	if d.tail == d.head {
		d.tail = nil
		d.head = nil
	} else {
		d.tail = d.tail.prev
		d.tail.next = nil
	}

	d.size--

	return element
}

func (d *Dequeue) DeleteFront() int {
	if d.size == 0 {
		return -9999
	}
	element := d.head.value
	if d.tail == d.head {
		d.tail = nil
		d.head = nil
	} else {
		d.head = d.head.next
		d.head.prev = nil
	}

	d.size--

	return element
}

func (q *Dequeue) Print() {
	temp := q.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

// func main() {
// 	// q := new(Queue)
// 	// q.Enqueue(1)
// 	// q.Enqueue(2)
// 	// q.Enqueue(3)
// 	// for q.IsEmpty() == false {
// 	// 	val, _ := q.Dequeue().(int)
// 	// 	fmt.Println(val, " ")
// 	// }

// 	// q := new(QueueLinkedList)
// 	// q.Addcircular(1)
// 	// q.Addcircular(2)
// 	// q.Addcircular(3)

// 	// for q.isEmpty() == false {
// 	// 	val := q.DeleteCircular()
// 	// 	fmt.Println(val, " ")
// 	// }

// 	q := new(Dequeue)
// 	q.AddFront(1)
// 	q.AddFront(1)
// 	q.AddBack(2)
// 	q.AddBack(2)
// 	q.AddFront(3)
// 	q.Print()
// 	fmt.Println(q.DeleteBack())
// 	q.Print()
// 	fmt.Println(q.DeleteFront())
// 	q.Print()
// }
