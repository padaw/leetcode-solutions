type StackNode struct {
	value int
	prev  *StackNode
}

type MyStack struct {
	tail *StackNode
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	node := StackNode{value: x}
	if this.tail != nil {
		node.prev = this.tail
	}
	this.tail = &node
}

func (this *MyStack) Pop() int {
	node := this.tail
	this.tail = node.prev
	node.prev = nil
	return node.value
}

func (this *MyStack) Top() int {
	return this.tail.value
}

func (this *MyStack) Empty() bool {
	return this.tail == nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
