package main

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func solution(l *ListNode) bool {
	cp := l
	stack := []interface{}{}
	for cp != nil {
		stack = append(stack, cp.Value)
		cp = cp.Next
	}
	for l != nil {
		cur := l.Value
		l = l.Next
		last := stack[len(stack)-1]
		if last != cur {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return true
}
