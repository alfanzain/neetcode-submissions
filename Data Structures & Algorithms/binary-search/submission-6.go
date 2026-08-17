func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, start int, end int) int {
	// Base case
	if start > end {
		return -1
	}

	// Not base case
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

	return binarySearch(nums, target, start, end)
}