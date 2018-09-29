package DataStructure

//MyCircularQueue :
type MyCircularQueue struct {
	data []int
	head int
	tail int
	size int
}

//Constructor Initialize your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k), -1, -1, k}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	tail := 0
	if this.tail < this.size-1 {
		tail = this.tail + 1
	}
	this.data[tail] = value
	this.tail = tail

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
		if this.head == this.size-1 {
			this.head = 0
		} else {
			this.head = this.head + 1
		}
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
	return this.head+this.tail == this.size
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
