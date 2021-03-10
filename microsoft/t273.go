package microsoft

import "fmt"

const (
	Hundred  = "Hundred"
	Thousand = "Thousand"
	Million  = "Million"
	Billion  = "Billion"
	Zero     = "Zero"
)

var (
	lessThanTwenty = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	lessThanHundred = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
)

func numberToWords(num int) string {
	if num == 0 {
		return Zero
	}
	billion := num / 1000000000
	million := (num - billion*1000000000) / 1000000
	thousand := (num - billion*1000000000 - million*1000000) / 1000
	rest := num - billion*1000000000 - million*1000000 - thousand*1000
	res := ""
	if billion != 0 {
		res += threeDigitToWords(billion) + " " + Billion + " "
	}
	if million != 0 {
		res += threeDigitToWords(million) + " " + Million + " "
	}
	if thousand != 0 {
		res += threeDigitToWords(thousand) + " " + Thousand + " "
	}
	if rest != 0 {
		res += threeDigitToWords(rest)
	}
	return res
}

func twoDigitToWords(num int) string {
	if num < 20 {
		return lessThanTwenty[num]
	}
	oneDigit := num % 10
	twoDigit := num / 10
	if oneDigit == 0 {
		return lessThanHundred[twoDigit]
	}
	return fmt.Sprintf("%s %s", lessThanHundred[twoDigit], lessThanTwenty[oneDigit])
}

func threeDigitToWords(num int) string {
	if num < 100 {
		return twoDigitToWords(num)
	}
	threeDigit := num / 100
	twoDigits := num - threeDigit*100
	if twoDigits == 0 {
		return fmt.Sprintf("%s %s", lessThanTwenty[threeDigit], Hundred)
	}
	return fmt.Sprintf("%s %s %s", lessThanTwenty[threeDigit], Hundred, twoDigitToWords(twoDigits))
}
