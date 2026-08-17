func groupAnagrams(strs []string) [][]string {
    guide := "00000000000000000000000000" // guide a-z as zeros
    dict := make(map[string][]string)

    // fmt.Println(int(guide[int('a')-97] - '0'))

    for _, str := range strs {
        // ASCII a-z -> 97-122
        // ASCII 0-9 -> 48-58?
        hash := []byte(guide)
        for _, ch := range str {
            hash[int(ch)-97] =  hash[int(ch)-97] + 1
        }
            
        dict[string(hash)] = append(dict[string(hash)], str)
    }

    // For debugging. Don't delete for the future
    // for _, strs := range dict {
    //     for _, str := range strs {
    //         fmt.Println("%+v", str)
    //     }
    //     fmt.Println()
    // }

    result := make([][]string, len(dict))
    i := 0
    for _, strs := range dict {
        for _, str := range strs {
            result[i] = append(result[i], str)
        }
        i++
    }

    return result
}

// func groupAnagrams(strs []string) [][]string {
//     dict := make(map[string][]int)

//     for idx, str := range strs {
//         str = sortStr(str)

//         dict[str] = append(dict[str], idx)
//     }

//     result := make([][]string, len(dict))
//     resultIdx := 0
//     for _, idxs := range dict {
//         for _, idx := range idxs {
//             result[resultIdx] = append(result[resultIdx], strs[idx])
//         }

//         resultIdx++
//     }
//     return result
// }

// // Bubble sort w/ flag
// func sortStr(str string) string {
// 	chars := []byte(str)
// 	n := len(chars)

// 	for i := 0; i < n; i++ {
// 		swapped := false

// 		for j := 0; j < n-1-i; j++ {
// 			if chars[j] > chars[j+1] {
// 				chars[j], chars[j+1] = chars[j+1], chars[j]
// 				swapped = true
// 			}
// 		}

// 		if !swapped {
// 			break
// 		}
// 	}

// 	return string(chars)
// }
