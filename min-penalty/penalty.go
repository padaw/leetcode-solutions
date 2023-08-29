func bestClosingTime(customers string) int {
	size := len(customers)
	score := 0
	best := 0
	holder := 0

	for i := 0; i < size; i++ {
		char := string(customers[i])

		if char == "Y" {
			score++
		} else {
			score--
		}

		if score > best {
			holder = i + 1
			best = score
		}
	}

	return holder
}
