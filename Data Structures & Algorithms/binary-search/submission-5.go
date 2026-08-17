func search(nums []int, target int) int {
	// Binary search:
	// From middle, go to left part or right part
	// whenever the target is less than or
	// more than the middle.
	// Loop until found.
	//
	// Binary search only works correctly when
	// nums is sorted.
	//
	// Example:
	// 1 2 3 4 5 6 7 8 9
	// s       m       e
	//
	// s = start
	// m = mid
	// e = end
	//
	// target: 8
	//
	// 1st loop:
	// 
	// Indexes:
	// start = 0
	// end = 8
	// mid = 4
	//
	// nums[mid] = 5
	//
	// 5 < 8, so target must be on the right side
	// Move start to mid + 1.
	//
	// start = 5
	// end = 8
	//
	// 1 2 3 4 5 6 7 8 9
	//           s m   e
	//
	// Now search only:
	// 6 7 8 9
	//
	// 2nd loop:
	// start = 5
	// end = 8
	// mid = 6
	// nums[mid] = 7
	//
	// 7 < 8, so target must be on the right side
	// Move start to mid + 1.
	//
	// start = 7
	// end = 8
	//
	// 1 2 3 4 5 6 7 8 9
	//               s e
	//				 m
	// 3rd loop:
	// start = 7
	// end = 8
	// mid = 7
	// nums[mid] = 8
	//
	// 8 == target, return index 7.
	//
	// If start becomes greater than end,
	// it means the target does not exist in nums.


	start := 0
	end := len(nums) - 1

	for start <= end {
		// Find the middle index of the current search range
		mid := start + (end-start)/2

		if nums[mid] > target {
			// Target is smaller, ignore the right half
			end = mid - 1
		} else if nums[mid] < target {
			// Target is larger, ignore the left half
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}