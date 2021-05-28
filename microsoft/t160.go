package microsoft

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	var lenA, lenB int
	for a != nil {
		a = a.Next
		lenA++
	}
	for b != nil {
		b = b.Next
		lenB++
	}
	aa, bb := headA, headB
	if lenA > lenB {
		aa, bb = bb, aa
		lenA, lenB = lenB, lenA
	}
	for i := 0; i < lenB-lenA; i++ {
		bb = bb.Next
	}
	for aa != nil {
		if aa == bb {
			return aa
		}
		aa = aa.Next
		bb = bb.Next
	}
	return nil
}
