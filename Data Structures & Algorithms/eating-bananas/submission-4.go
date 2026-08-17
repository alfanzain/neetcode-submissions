func minEatingSpeed(piles []int, h int) int {
	// More banana / hour (k) = less total hours

	// Now find the k
	start := 1

	// DON'T DELETE THIS COMMENT, FUTURE ME!
	// THIS NOTE IS DEPCREATED SINCE IDK WHY THE SORT DOESNT WORK
	// BUT I NEED TO REDO THE EXPERIMENT IN THE FUTURE
	//
	// The slice should be sorted before using max()
	// The syntax arr[:] is called a slice expression.
	// It turns your fixed-size array into a slice that points to the exact same data in memory.
	// slices.Sort(piles[:])
	// end := max(piles) // The most banana in the piles // Idk why max() doesn't work

	// Find the most banana in the piles
	end := piles[0]
	for _, pile := range piles {
		end = max(pile, end)
	}
	// candidates := make([]int, 0)
	candidate := 0

	for start <= end {
		mid := start + (end-start)/2

		// Test to all piles
		totalHours := 0
		for _, pile := range piles {
			hours := math.Ceil(float64(pile) / float64(mid))
			totalHours += int(hours)
		}

		// fmt.Println("k:", mid)
		// fmt.Println("totalHours:", totalHours)

		// If total hours less than h, so increase k
		// Which mean, go to the right side
		if totalHours > h {
			start = mid + 1
		} else if totalHours <= h {
			end = mid - 1
			candidate = mid
		}

		// fmt.Println()

		if start > end {
			break
		}
	}

	// fmt.Println(candidates)

	return candidate
}