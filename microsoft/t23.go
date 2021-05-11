package microsoft

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeHelper(lists, 0, len(lists)-1)
}

func mergeHelper(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists(mergeHelper(lists, l, mid), mergeHelper(lists, mid+1, r))
}
