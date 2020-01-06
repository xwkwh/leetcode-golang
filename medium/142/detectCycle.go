package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	if head == nil || head.Next == nil {
		return nil
	}
	isCycle := false
	for fast != nil && slow != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}
	if !isCycle {
		return nil
	}
	slow = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
