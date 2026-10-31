/*

Problem Url: https://leetcode.com/problems/add-two-numbers/

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

*/

package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return addTwoNumbersUtil(l1, l2, 0)
}

func addTwoNumbersUtil(l1 *ListNode, l2 *ListNode, carry int) (summationNode *ListNode) {

	if l1 == nil && l2 == nil && carry == 0 {
		return summationNode
	}

	var l1Val, l2Val int
	var l1Next, l2Next *ListNode
	
	if l1 != nil {
		l1Val = l1.Val
		l1Next = l1.Next
	}
	if l2 != nil {
		l2Val = l2.Val
		l2Next = l2.Next
	}
	summation := l1Val + l2Val + carry
	summationNode = &ListNode{}
	summationNode.Val, summationNode.Next = summation % 10, addTwoNumbersUtil(l1Next, l2Next, summation/10)
	return summationNode
}

func main() {
	num1 := ListNode{}
}
