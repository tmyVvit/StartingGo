package leetcode

//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/solution/mian-shi-ti-09-yong-liang-ge-zhan-shi-xian-dui-l-3/

// CQueue is a queue that using two stacks
type CQueue struct {
	stack1, stack2 []int
}

// NewCQueue return a new CQueue
func NewCQueue() CQueue {
	return CQueue{stack1: []int{}, stack2: []int{}}
}

// AppendTail add a value to the end of the queue
func (cq *CQueue) AppendTail(value int) {
	cq.stack1 = append(cq.stack1, value)
}

// DeleteHead delete the value in the head of the queue and return it
// if queue is empty return -1
func (cq *CQueue) DeleteHead() int {
	if len(cq.stack2) == 0 {
		for i := len(cq.stack1) - 1; i >= 0; i-- {
			cq.stack2 = append(cq.stack2, cq.stack1[i])
		}
		cq.stack1 = []int{}
	}
	if len(cq.stack2) == 0 {
		return -1
	}
	ret := cq.stack2[len(cq.stack2)-1]
	cq.stack2 = cq.stack2[:len(cq.stack2)-1]
	return ret
}
