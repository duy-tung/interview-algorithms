package car_fueling

func car_fueling(dist int, tank int, stops []int) int {
	fills := 0
	recent := 0
	current := 0

	stops = append([]int{0}, append(stops, dist)...)
	for current < len(stops)-2 {
		recent = current
		for current < len(stops)-2 && stops[current+1]-stops[recent] <= tank {
			current++
		}
		if current == recent {
			return -1
		}
		if current < len(stops)-2 {
			fills++
		}
	}

	return fills
}
