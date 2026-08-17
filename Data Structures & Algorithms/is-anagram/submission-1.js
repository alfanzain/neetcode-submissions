class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
        if (s.length != t.length) {
            return false
        }

        let h1 = {}
        let h2 = {}
        for (let i = 0; i < s.length; i++) {
            if (h1[s[i]] === undefined) {
                h1[s[i]] = 0
            }
            if (h2[t[i]] === undefined) {
                h2[t[i]] = 0
            }

            h1[s[i]]++
            h2[t[i]]++
        }

        for (let i = 0; i < s.length; i++) {
            if (h1[s[i]] === undefined) {
                h1[s[i]] = 0
            }
            if (h2[t[i]] === undefined) {
                h2[t[i]] = 0
            }

            h1[s[i]]++
            h2[t[i]]++
        }

        for (const key in h1) {
            if (h1[key] !== h2[key]) {
                return false;
            }
        }

        return true;
    }
}
