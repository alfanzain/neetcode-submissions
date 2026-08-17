type MinStack struct {
	items []int
	size  int
}

func Constructor() MinStack {
	return MinStack{
		items: make([]int, 0),
		size:  0,
	}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	this.size++
}

func (this *MinStack) Pop() {
	if this.size == 0 {
		return
	}
	this.items = this.items[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	if this.size == 0 {
		return 0
	}

	return this.items[this.size-1]
}

func (this *MinStack) GetMin() int {
	if this.size == 0 {
		return 0
	}

	min := this.items[0]
	for _, val := range this.items {
		if val < min {
			min = val
		}
	}

	return min
}