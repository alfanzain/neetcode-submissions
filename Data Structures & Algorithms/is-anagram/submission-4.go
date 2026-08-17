func isAnagram(s string, t string) bool {
    counts := make(map[rune]int)
    
    for _, ch := range s {
        counts[ch]++
    }

    for _, ch := range t {
        counts[ch]--
    }

    for _, count := range counts {
        fmt.Print(count)
        if count != 0 {
            return false
        }
    }

    return true
}
