func searchMatrix(matrix [][]int, target int) bool {
	// w/ Binary Search (?)
	//
	// Because it's increasing order, and maybe we can use (kinda) binary search
	// The pointers point the index of the row
	// So, the start is first row, the end is the last row, and the mid is the mid row
	//
	// If found, use the row for the list. Search with binary search again
	//
	// matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]], target = 12
	//
	// In this case:
	// Indexes:
	// - start = 0
	// - mid = 1
	// - end = 2
	//
	// If target < nums[mid][0], ignore mid to the right rows
	// If target > nums[mid][0], ignore left rows. Don't ignore mid since it might be in the mid row
	// If target > nums[mid][0], get the mid row.

	// fmt.Println("target:", target)

	n := len(matrix)
	m := len(matrix[0])
	startRow, endRow := 0, n-1

	for startRow <= endRow {
		midRow := startRow + (endRow-startRow)/2
		currMidRow := matrix[midRow]
		// 1st ->
		// 	startRow = 0
		// 	endRow = 2
		// 	midRow = endRow/2 = 2/2 = 1
		//
		// 	row = [10, »11«, 12, 13]
		// 	startCol = 0
		// 	endCol = 3
		// 	midCol = endCol/2 = 1
		// fmt.Println("startRow:", startRow, "endRow:", endRow, "midRow:", midRow)

		startCol, endCol := 0, m-1

		if currMidRow[startCol] > target {
			endRow = midRow - 1
			continue
		} else if currMidRow[endCol] < target {
			startRow = midRow + 1
			continue
		}

		midCol := endCol / 2
		for startCol <= endCol {
			midCol = startCol + (endCol-startCol)/2
			currNum := currMidRow[midCol]
			// fmt.Println("startCol:", startCol, "endCol:", endCol, "midCol:", midCol)

			if currNum > target {
				endCol = midCol - 1
			} else if currNum < target {
				startCol = midCol + 1
			} else {
				// fmt.Println("num:", currNum)
				return true
			}
		}

		return false

		// fmt.Println()
	}

	return false
}