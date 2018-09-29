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
	return MyCircularQueue{make([]int, 0, k), 0, 0, k}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	return 0
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	return 0
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return true
}

//IsFull Checks whether the circular queue is full or not.
func (this *MyCircularQueue) IsFull() bool {
	return len(this.data) == this.size
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
