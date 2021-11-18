package addTwoNumber

// ListNode Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	result := new(ListNode)
	cur := result
	for l1 != nil || l2 != nil {
		a1 := 0
		if l1 != nil {
			a1 = l1.Val
		}

		a2 := 0
		if l2 != nil {
			a2 = l2.Val
		}

		sum := a1 + a2  + add
		add = sum / 10
		n := sum % 10

		cur.Next = &ListNode{
			Val:  n,
		}
		cur = cur.Next
		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if add > 0 {
		cur.Next =   &ListNode{
			Val:  add,
			Next: nil,
		}
	}
	return result.Next
}
