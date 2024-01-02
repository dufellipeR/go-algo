package kata

func CalculateYears(years int) (result [3]int) {
	result = [3]int{}
	var catYears int
	var dogYears int

	if years > 2 {
		catYears = (years-2)*4 + 24
		dogYears = (years-2)*5 + 24
	} else if years == 2 {
		catYears = 24
		dogYears = 24
	} else if years == 1 {
		catYears = 15
		dogYears = 15
	}

	result[0] = years
	result[1] = catYears
	result[2] = dogYears

	return result
}
