package section1

func filterAges(minAge int, maxAge int, ages []int) []int {
	var agesFiltered []int
	for _, age := range ages {
		if minAge <= age && age <= maxAge {
			agesFiltered = append(agesFiltered, age)
		}
	}
	return agesFiltered
}