package datastructure

import "fmt"

func TestMyCircularQueue() {
	param1 := true

	obj := ConstructorCQ(3)
	fmt.Println(obj)
	param1 = obj.EnQueue(1)
	fmt.Println(obj, param1)
	param1 = obj.EnQueue(2)
	fmt.Println(obj, param1)
	param1 = obj.EnQueue(3)
	fmt.Println(obj, param1)
	param1 = obj.EnQueue(4)
	fmt.Println(obj, param1)
	// param1 = obj.DeQueue()
	// fmt.Println(obj, param1)
	// param1 = obj.EnQueue(4)
	// fmt.Println(obj, param1)
	// param1 = obj.EnQueue(2)
	// fmt.Println(obj, param1)
}

//MyCircularQueue :
type MyCircularQueue struct {
	data []int
	head int
	tail int
	size int
}

//Constructor Initialize your data structure here. Set the size of the queue to be k.
func ConstructorCQ(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k), -1, -1, k}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.tail = (this.tail + 1) % this.size
	this.data[this.tail] = value

	if this.head == -1 {
		this.head = 0
	}

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.data[this.head] = 0
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
	} else {
		this.head = (this.head + 1) % this.size
	}

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1 && this.tail == -1
}

//IsFull Checks whether the circular queue is full or not.
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.size == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
