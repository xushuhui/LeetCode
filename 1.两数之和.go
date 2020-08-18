package main
/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
    for k, num := range nums {
		diff :=  target - num
		if v,ok := m[diff];ok{
			return []int{v,k}
		}
		m[num] = k	
	}
	return nil
}
func main()  {
	twoSum([]int{2, 7, 11, 15},9)
}
// @lc code=end

//该题目主要考的是数组查找，数组查找一个元素时间复杂度是O(n)，查找两个时间复杂度是O(n^2)。
//我们是不是可以试着分两步，
//1.只查一个值，O(n)
//2.判断target-num的差值是否在剩余的数组中，如果还是数组，O(n)，如果我们能转换成map，时间复杂度就只有O(1)
//第一个解法
//暴力两重循环，时间复杂度是O(n^2)，明显这种复杂度是不能接受的
//第二个解法
//1.初始化map数据结构
//2.循环nums，找到target-num的差值diff
//3.如果map中查找到key是diff的值，就说明符合条件，返回该值和num对应的索引k
//4.如果找不到，就把num作为m的key，num的索引k作为map的value
// 

