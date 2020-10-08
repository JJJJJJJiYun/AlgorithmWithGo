package o59

func MaxOfQueue(nums []int, k int) []int {
	res := make([]int, 0)
	record := make([]int, 0)
	for i := range nums {
		if len(record) != 0 && i-record[0] >= k {
			record = record[1:]
		}
		for len(record) != 0 && nums[record[len(record)-1]] < nums[i] {
			record = record[:len(record)-1]
		}
		record = append(record, i)
		if len(record) != 0 && i >= k-1 {
			res = append(res, nums[record[0]])
		}
	}
	return res
}
