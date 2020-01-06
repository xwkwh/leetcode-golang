package main

import "leetcode-golang/lib"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *lib.ListNode) *ListNode {
	cur := head
	var already *lib.ListNode

	for cur != nil {
		next := cur.Next   // 保存next节点 一会可能就没了
		cur.Next = already // 当前的next指针指向已经反转好的
		already = cur      // 移动已经反转好的节点到头节点
		cur = next         // cur  用保存好的next 进行遍历
	}

	// for cur != nil {
	// 	cur.Next, already, cur = already, cur, cur.Next
	// }

	dummy := new(lib.ListNode)
	dummy.Next = head
	point := dummy

	for point.Next != nil {
		tmp := point.Next
		point.Next = tmp.Next
		tmp.Next = point
		point = tmp
	}
	return already
}
