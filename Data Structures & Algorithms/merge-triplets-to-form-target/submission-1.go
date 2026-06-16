func mergeTriplets(triplets [][]int, target []int) bool {
	foundX, foundY, foundZ := false, false, false

	for _, t := range triplets {
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue
		}

		if t[0] == target[0] {
			foundX = true
		}

		if t[1] == target[1] {
			foundY = true
		}

		if t[2] == target[2] {
			foundZ = true
		}

		if foundX && foundY && foundZ {
			return true
		}
	}

	return false
}