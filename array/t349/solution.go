package t349

func intersection(nums1 []int, nums2 []int) []int {
	record := make(map[int]struct{})
	for _, num := range nums1 {
		record[num] = struct{}{}
	}
	var res []int
	for _, num := range nums2 {
		if _, ok := record[num]; ok {
			res = append(res, num)
			delete(record, num)
		}
	}
	return res
}
