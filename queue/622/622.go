package main

type MyCircularQueue struct {
	front int
	tail int
	Slice []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		front: 0,
		tail:  0,
		Slice: make([]int, k+1),
	}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Slice[this.tail] = value
	this.tail = (this.tail + 1) % cap(this.Slice)
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % cap(this.Slice)
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Slice[this.front]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Slice[(this.tail - 1 + cap(this.Slice)) % cap(this.Slice)]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.tail
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail + 1) % cap(this.Slice) == this.front
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

func main() {
	
}
