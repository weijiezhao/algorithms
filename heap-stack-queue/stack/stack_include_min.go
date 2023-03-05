package stack

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
	if len(stack2) == 0 {
		stack2 = append(stack2, node)
		return
	}
	if node > stack2[len(stack2)-1] {
		stack2 = append(stack2, stack2[len(stack2)-1])
	} else {
		stack2 = append(stack2, node)
	}
}

func Pop() {
	stack1 = stack1[:len(stack1)-1]
	stack2 = stack2[:len(stack2)-1]
}

func Top() int {
	return stack1[len(stack1)-1]
}

func Min() int {
	return stack2[len(stack2)-1]
}
