package main

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func solution(l *ListNode, k int) *ListNode {
	if l == nil {
		return l
	}
	for l.Value == k && l != nil {
		l = l.Next
		if l == nil {
			break
		}
	}
	if l == nil {
		return l
	}
	cp := l
	for cp.Next != nil {
		if cp.Next.Value == k {
			cp.Next = cp.Next.Next
		} else {
			cp = cp.Next
		}
	}
	return l
}

func main() {
	solution()
}
