package main


//同步辅助栈

type MinStack struct {
	data []int
	helper []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:     nil,
		helper:   nil,
	}
}


func (this *MinStack) Push(x int)  {
	this.data = append(this.data, x)
	if len(this.helper) == 0 || x <= this.helper[len(this.helper) - 1] {
		this.helper = append(this.helper, x)
	} else {
		this.helper = append(this.helper, this.helper[len(this.helper) - 1])
	}
}


func (this *MinStack) Pop()  {
	this.data = this.data[:len(this.data) - 1]
	this.helper = this.helper[:len(this.helper) - 1]
}


func (this *MinStack) Top() int {
	return this.data[len(this.data) - 1]
}


func (this *MinStack) GetMin() int {
	return this.helper[len(this.helper) - 1]
}

func main() {
	
}
