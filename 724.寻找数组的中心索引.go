/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 */

// @lc code=start
func pivotIndex(nums []int) int {
	var sum,leftNum int
	for _, num := range nums {
		sum += num 
	}
	for i := 0; i < len(nums); i++ {
		if leftNum == sum - nums[i] - leftNum {
			return i
		}
		leftNum += nums[i]
	}
	return -1

}
// @lc code=end

