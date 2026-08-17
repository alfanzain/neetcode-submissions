func groupAnagrams(strs []string) [][]string {
    dict := make(map[[26]int][]string)

    for _, str := range strs {
        // fmt.Println(str)
        hash := [26]int{}
        for _, char := range str {
            letterIdx := char-97
            hash[letterIdx]++
        }
        // fmt.Printf("%+v\n", hash)

        dict[hash] = append(dict[hash], str)
    }

    var result [][]string
    for _, strs := range dict {
        result = append(result, strs)
    }

    return result
}