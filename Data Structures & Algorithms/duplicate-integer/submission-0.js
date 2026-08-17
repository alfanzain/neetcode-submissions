class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        let uniqueNums = new Set(nums)

        return (nums.length !== uniqueNums.size)
    }
}
