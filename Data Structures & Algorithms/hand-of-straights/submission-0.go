import "slices"

func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand) < groupSize {
		return false
	}

	counts := make(map[int]int)
	for _, cards := range hand {
		counts[cards]++ 
	}

	uniqueCards := make([]int, len(counts))
	for count := range counts {
		uniqueCards = append(uniqueCards, count)
	}

	slices.Sort(uniqueCards)

	for _, cards := range uniqueCards {
		count := counts[cards]

		if count > 0 {
			for i := 0; i < groupSize; i++ {
				if counts[cards + i] < count {
					return false
				}

				counts[cards + i] -= count 
			}
		}
	}

	return true
}