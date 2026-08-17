class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        char_count = dict()
        for ch in s:
            if ch not in char_count:
                char_count[ch] = 1
            else:
                char_count[ch] += 1

        for ch in t:
            if ch not in char_count:
                char_count[ch] = 1
            else:
                char_count[ch] -= 1

        for char in char_count:
            if char_count[char] != 0:
                return False

        return True

            