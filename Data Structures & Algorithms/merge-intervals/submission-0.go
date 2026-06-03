import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
    // Edge case: If there are 0 or 1 intervals, there's nothing to merge
	if len(intervals) <= 1 {
		return intervals
	}

	// Step 1: Sort the intervals by their start time
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	var merged [][]int
	
	// Step 2: Add the first interval to start the process
	merged = append(merged, intervals[0])

	// Step 3: Iterate through the remaining intervals
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		
		// Grab a reference to the last interval we added to `merged`
		lastMerged := merged[len(merged)-1]

		// Check for overlap
		if current[0] <= lastMerged[1] {
			// They overlap! Merge them by updating the end time.
			// In Go, since slices are references, updating `lastMerged[1]` 
			// directly updates the value inside the `merged` slice.
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			// No overlap! Add the new interval to our list.
			merged = append(merged, current)
		}
	}

	return merged
}
