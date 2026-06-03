import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 1. Sort the intervals
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	// 2. Use a pointer to track the position of the last merged interval
	insertPos := 0 

	// 3. Iterate through the array starting from the second element
	for i := 1; i < len(intervals); i++ {
		// Check for overlap with the interval at insertPos
		if intervals[i][0] <= intervals[insertPos][1] {
			// Overlap: update the end time
			if intervals[i][1] > intervals[insertPos][1] {
				intervals[insertPos][1] = intervals[i][1]
			}
		} else {
			// No overlap: move the insert pointer forward and copy the interval
			insertPos++
			intervals[insertPos] = intervals[i]
		}
	}

	// 4. Return the sub-slice containing only the merged intervals
	return intervals[:insertPos+1]
}

