package main

import "math"

// Angle Between Hands of a Clock

func angleClock(hour int, minutes int) float64 {
	perMinute := 6.0
	hourPerMinute := 0.5
	minutesAngle := float64(minutes) * perMinute
	hourAngle := float64(hour*60+minutes) * hourPerMinute

	res := math.Abs(minutesAngle - hourAngle)
	if res > 180 {
		res = 360.0 - res
	}
	return res
}
