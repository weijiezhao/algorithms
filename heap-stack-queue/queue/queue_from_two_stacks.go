package queue

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		for i := 0; i < len(stack1); i++ {
			stack2 = append(stack2, stack1[len(stack1)-i-1])
		}
		stack1 = []int{}
	}
	res := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return res
}
