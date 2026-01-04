package addtwonumbers

import "testing"

func buildListNode(l []int) *ListNode {
	dump := &ListNode{}
	curr := dump
	for _, val := range l {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dump.Next
}

func TestAddTwoNumbers(t *testing.T) {
	list1 := []int{2, 4, 3}
	list2 := []int{5, 6, 4}
	want := []int{7, 0, 8}
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			buildListNode(list1),
			buildListNode(list2),
			buildListNode(want),
		},
	}
	for _, tt := range tests {
		got := addTwoNumbers(tt.l1, tt.l2)
		if !(compareListNode(got, tt.want)) {
			t.Errorf("addTwoNumbers false")
		}
	}
}

func compareListNode(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
