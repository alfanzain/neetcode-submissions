func searchMatrix(matrix [][]int, target int) bool {
	// Here we try to find the target by binary search 1d
	//
	// The different with previous solution is, the loop only 1.
	// Consider the 2d matrix as 1d matrix

	n := len(matrix)
	m := len(matrix[0])
	start := 0

	// Let say 2x5 matrix. So, it's 10-length matrix.
	// If the end is 10-1 = 9, the challenge is to define which cell is the end since it's 2d matrix
	end := n*m - 1

	// 0,0 0,1 0,2 0,3 0,4
	// 1,0 1,1 1,2 1,3 1,4
	//
	// 	0	1	2	3	4
	//	5	6	7	8	9

	for start <= end {
		// Find the middle index of the current search range
		mid := start + (end-start)/2
		fmt.Println("mid:", mid)
		fmt.Println("start:", start)
		fmt.Println("end:", end)

		// Define matrix indexes
		row := mid / m
		col := mid % m
		fmt.Println("row:", row)
		fmt.Println("col:", col)

		if matrix[row][col] > target {
			// Target is smaller, ignore the right half
			end = mid - 1
		} else if matrix[row][col] < target {
			// Target is larger, ignore the left half
			start = mid + 1
		} else {
			return true
		}

		fmt.Println()
	}

	return false
}
