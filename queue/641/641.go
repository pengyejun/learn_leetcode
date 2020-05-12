package main

type MyCircularDeque struct {
	front int
	tail int
	Slice []int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		front: 0,
		tail:  0,
		Slice: make([]int, k+1),
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front - 1 + cap(this.Slice)) % cap(this.Slice)
	this.Slice[this.front] = value
	return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Slice[this.tail] = value
	this.tail = (this.tail + 1) % cap(this.Slice)
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % cap(this.Slice)
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + cap(this.Slice)) % cap(this.Slice)
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Slice[this.front]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Slice[(this.tail - 1 + cap(this.Slice)) % cap(this.Slice)]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.tail
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.tail + 1) % cap(this.Slice) == this.front
}

func main() {
	
}
