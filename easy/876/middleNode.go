package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{}
	cur := &ListNode{}
	cur.Next = head
	pre.Next = head
	for pre.Next != nil && pre.Next.Next != nil {
		cur = cur.Next
		pre = pre.Next.Next
	}
	return cur.Next
}
