import "slices"

func isValidSudoku(board [][]byte) bool {
	// 1. Each row must contain the digits 1-9 without duplicates.
	// 2. Each column must contain the digits 1-9 without duplicates.
	// 3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
	var rows [9][]byte
	var cols [9][]byte
	var subBoxes [9][]byte

	// Move diagonally
	// 0,0 <-- top-left
	// 1,1
	// until
	// 8,8 <-- bottom-right
	//
	// Loop i: 0-8
	// Row: board[i, 0-8]
	// Col: board[0-8, i]
	//
	// For sub-box, the index count like this:
	//
	// 0 1 2
	// 3 4 5
	// 6 7 8
	//
	// Sub-box pattern:
	//
	// First sub-box:
	// 0,0  0,1  0,2
	// 1,0  1,1  1,2
	// 2,0  2,1  2,2
	//
	// Second sub-box:
	// 0,3  0,4  0,5
	// 1,3  1,4  1,5
	// 2,3  2,4  2,5
	//
	// Fourth sub-box:
	// 3,0  3,1  3,2
	// 4,0  4,1  4,2
	// 5,0  5,1  5,2

	for i := range 9 {
		// Check row
		for j := range 9 {
			// Check cell in current row
			currentTile := board[i][j]
			fmt.Printf("(%d,%d): %c\n", i, j, currentTile)
			if currentTile == '.' {
				continue
			}
			// Check duplicate
			if slices.Contains(rows[i], currentTile) {
				return false
			}

			rows[i] = append(rows[i], currentTile)
		}

		// Check col
		for j := range 9 {
			// Check cell in current col
			currentTile := board[j][i]
			if currentTile == '.' {
				continue
			}
			// Check duplicate
			if slices.Contains(cols[i], currentTile) {
				return false
			}

			cols[i] = append(cols[i], currentTile)
		}

		// Check sub-box
		// Find row per i -> (i / 3) * 3
		// Find col per i -> (i mod 3) * 3
		initialRow := (i / 3) * 3
		finalRow := initialRow + 2
		initialCol := (i % 3) * 3
		finalCol := initialCol + 2
		for j := initialRow; j <= finalRow; j++ {
			for k := initialCol; k <= finalCol; k++ {
				// Check cell in current sub-box
				currentTile := board[j][k]
				if currentTile == '.' {
					continue
				}
				// Check duplicate
				if slices.Contains(subBoxes[i], currentTile) {
					return false
				}

				subBoxes[i] = append(subBoxes[i], currentTile)
			}
		}
	}

	return true
}