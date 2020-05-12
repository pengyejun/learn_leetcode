package main

//同步辅助栈

type MinStack1 struct {
	data []int
	helper []int
}


/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		data:     nil,
		helper:   nil,
	}
}


func (this *MinStack1) Push(x int)  {
	this.data = append(this.data, x)
	if len(this.helper) == 0 || x <= this.helper[len(this.helper) - 1] {
		this.helper = append(this.helper, x)
	}
}


func (this *MinStack1) Pop()  {
	top := this.data[len(this.data) - 1]
	this.data = this.data[:len(this.data) - 1]
	if top == this.helper[len(this.helper) - 1] {
		this.helper = this.helper[:len(this.helper) - 1]
	}
}


func (this *MinStack1) Top() int {
	return this.data[len(this.data) - 1]
}


func (this *MinStack1) GetMin() int {
	return this.helper[len(this.helper) - 1]
}

func main() {

}
