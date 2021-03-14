package microsoft

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 0 {
		return (float64(getKthInTwoArrays(nums1, nums2, totalLen/2)) + float64(getKthInTwoArrays(nums1, nums2, totalLen/2+1))) / 2
	}
	return float64(getKthInTwoArrays(nums1, nums2, totalLen/2+1))
}

func getKthInTwoArrays(nums1, nums2 []int, k int) int {
	var index1, index2 int
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		newIndex1 := min(index1+k/2-1, len(nums1)-1)
		newIndex2 := min(index2+k/2-1, len(nums2)-1)
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 < pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	var i, j, k, num, num1, num2, index1, index2 int
	if (len1+len2)%2 == 0 {
		index1 = (len1+len2)/2 - 1
		index2 = index1 + 1
	} else {
		index1 = (len1 + len2) / 2
		index2 = index1
	}
	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			num = nums1[i]
			i++
		} else {
			num = nums2[j]
			j++
		}
		if k == index1 {
			num1 = num
		}
		if k == index2 {
			num2 = num
			k++
			break
		}
		k++
	}
	for k <= index2 && i < len1 {
		num = nums1[i]
		if k == index1 {
			num1 = num
		}
		if k == index2 {
			num2 = num
			break
		}
		i++
		k++
	}
	for k <= index2 && j < len2 {
		num = nums2[j]
		if k == index1 {
			num1 = num
		}
		if k == index2 {
			num2 = num
			break
		}
		j++
		k++
	}
	return (float64(num1) + float64(num2)) / 2
}
