//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
//
//
// 示例 1:
//
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//
//
// 示例 2:
//
//
//输入: "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 105
// 字符串 s 由 ASCII 字符组成
//
// Related Topics 双指针 字符串
// 👍 405 👎 0

package leetcode

import (
	"strings"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    left,right := 0,len(s)-1 //设置首尾指针
    for left < right{
        for left < right && !isValid(s[left]){
            left++
        }
        for left < right && !isValid(s[right]){
            right--
        }
        if left < right && s[left] != s[right]{
            return false
        }
        left++
        right--
    }
    return true
}
func isValid(b byte)bool{
    if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z'){
        return true
    }
    return false
}

//leetcode submit region end(Prohibit modification and deletion)
func isPalindrome1(s string) bool {
    //转小写
    //只保留数字和字符串
    //翻转后和原字符串比较
	s2 := strip([]byte(strings.ToLower(s)))
    return s2 == reverseStr(s2)
}
func isPalindrome2(s string) bool {
    s = strings.ToLower(s)
    var s2 string
	for i := 0; i < len(s); i++ {
		if isValid(s[i]){
			s2 += string(s[i])
		}
	}
	l := len(s2)
	for i := 0; i < l/2; i++ {

		if s2[i] != s2[l-1-i]{
			return false
		}
	}

    return true
}

func strip(s []byte) string {
	n := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||('0' <= b && b <= '9') {
			s[n] = b
			n++
		}
	}
	return string(s[:n])
}
func reverseStr(s string) string {
	a := func(s string) *[]rune {
		var b []rune
		for _, k := range []rune(s) {
			defer func(v rune) {
				b = append(b, v)
			}(k)
		}
		return &b
	}(s)
	return string(*a)
}
func Test_validPalindrome(t *testing.T) {
	t.Log(isPalindrome3("A man, a plan, a canal: Panama"))
}
