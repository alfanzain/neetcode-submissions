func topKFrequent(nums []int, k int) []int {
	dictNumberFreq := make(map[int]int) // key -> number; value -> freq
	dictFreq := make(map[int][]int)     // key -> freq; value -> list of numbers
	maxK := 0

	for _, num := range nums {
		dictNumberFreq[num]++

		if dictNumberFreq[num] > maxK {
			maxK = dictNumberFreq[num]
		}
	}

	fmt.Println("maxK:", maxK)

	for num, freq := range dictNumberFreq {
		dictFreq[freq] = append(dictFreq[freq], num)
	}

	// fmt.Printf("dictFreq: %+v\n", dictFreq)

	result := []int{}
	for k > 0 {
		if listOfNums, exists := dictFreq[maxK]; exists {
			result = append(result, listOfNums...)
			k = k - len(listOfNums)
		}

		maxK--
	}

	return result
}
