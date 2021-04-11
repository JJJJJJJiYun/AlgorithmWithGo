package microsoft

var phoneMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	helper(digits, "", 0, &res)
	return res
}

func helper(digits, cur string, index int, res *[]string) {
	if index == len(digits) {
		*res = append(*res, cur)
		return
	}
	for _, c := range phoneMap[digits[index]] {
		helper(digits, cur+string(c), index+1, res)
	}
}
