package t1344

import "math"

func angleClock(hour int, minutes int) float64 {
	hourAngle := float64(30)
	minuteAngle := float64(6)
	floatMinutes := float64(minutes)
	hourAngles := (float64(hour%12) + floatMinutes/60) * hourAngle
	minuteAngles := minuteAngle * floatMinutes
	diff := math.Abs(minuteAngles - hourAngles)
	return math.Min(diff, 360-diff)
}
