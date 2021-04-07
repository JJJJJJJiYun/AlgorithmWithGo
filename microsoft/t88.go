package microsoft

func merge2(nums1 []int, m int, nums2 []int, n int) {
	if n <= 0 {
		return
	}
	i, j, k := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		k--
		i--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
