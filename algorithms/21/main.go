package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// PrintListReverse 递归打印链表
func PrintListReverse(ptr *ListNode) {
	if ptr != nil {
		fmt.Print(ptr.Val, " ")
		if ptr.Next != nil {
			PrintListReverse(ptr.Next)
		}
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preheat := &ListNode{Val: 0, Next: nil}

	prev := preheat

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return preheat.Next
}

func main() {
	curA := &ListNode{}
	for _, v := range []int{2, 3, 4, 5} {
		a := &ListNode{
			Val:  v,
			Next: curA,
		}
		curA = a
	}

	curB := &ListNode{}
	for _, v := range []int{4, 9, 40, 53} {
		a := &ListNode{
			Val:  v,
			Next: curB,
		}
		curB = a
	}

	PrintListReverse(curA)
	fmt.Println()
	PrintListReverse(curB)
	fmt.Println()
	PrintListReverse(mergeTwoLists(curA, curB))
}
