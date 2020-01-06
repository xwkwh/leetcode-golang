package main

import "leetcode-golang/lib"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *lib.ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := head.Next
	head.Next = swapPairs(head.Next.Next)
	n.Next = head
	return n
}

func swapPairs2(head *lib.ListNode) *ListNode {
	dummy := new(lib.ListNode)
	dummy.Next = head
	point := dummy

	for point.Next != nil && point.Next.Next != nil {
		swap1 := point.Next
		swap2 := point.Next.Next
		point.Next = swap2
		swap1.Next = swap2.Next
		swap2.Next = swap1
		point = swap1
	}
	return dummy.Next
}
