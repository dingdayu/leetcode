package main

import "fmt"

// Definition for singly-linked list.
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	first := dummy
	second := dummy
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range []int{2, 3, 4, 5} {
		a := &ListNode{
			Val:  v,
			Next: cur,
		}
		cur = a
	}
	PrintListReverse(cur)
	fmt.Println()
	PrintListReverse(removeNthFromEnd(cur, 2))
}
