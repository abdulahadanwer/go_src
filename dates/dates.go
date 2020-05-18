package dates

//DaysInweek total days in a week
const DaysInweek int = 7

// WeekToDays : convert weeks into days
func WeekToDays(weeks int) int {
	return weeks * DaysInweek
}

// DaysToWeek : convert days into weeks
func DaysToWeek(days int) float64 {
	return float64(days) / float64(DaysInweek)
}
