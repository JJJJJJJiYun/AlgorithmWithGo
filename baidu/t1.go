package baidu

// 50亿数据，内容是请求延时0-500ms的数字，50，90，99分位

var (
	nums = [501]int{}
)

const (
	all = 5000000000
)

func getNum(n int) {
	nums[n]++
}

func getMPercentOfNums(m float32) int {
	index := int(all * m)
	count := 0
	for i, v := range nums {
		count += v
		if count >= index {
			return i
		}
	}
	return -1
}
