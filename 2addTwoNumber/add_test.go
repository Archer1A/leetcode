package addTwoNumber

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T)  {
	i := &ListNode{
		Val:  9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: &ListNode{
					Val:  9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	j :=&ListNode{
		Val:  9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	res := addTwoNumbers(i,j)
	for res!= nil {
		fmt.Print(res.Val)
		res = res.Next
	}

}
