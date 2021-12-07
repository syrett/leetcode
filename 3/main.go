package main

import (
	"fmt"
	"math"
	"strings"
)

/*
无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0


提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	strMap := make(map[rune]int, 0)
	tmpStrSlice := make([]rune, 0)
	for _, v := range s {
		repeatPoint, ok := strMap[v]
		if !ok {
			strMap[v] = len(tmpStrSlice)
			tmpStrSlice = append(tmpStrSlice, v)
			continue
		}

		fmt.Printf("============\n")
		for i, j := range tmpStrSlice {
			fmt.Printf("%v:%v\n", i, j)
			strMap[j] = i
		}

		maxLen = int(math.Max(float64(len(tmpStrSlice)), float64(maxLen)))
		tmpStrSlice = append(tmpStrSlice[repeatPoint+1:], v)
		strMap = make(map[rune]int, 0)

		fmt.Printf("---------------\n")
		for i, j := range tmpStrSlice {
			fmt.Printf("%v:%v\n", i, j)
			strMap[j] = i
		}
	}
	maxLen = int(math.Max(float64(len(tmpStrSlice)), float64(maxLen)))
	return maxLen
}

func LengthOfLongestSubstring2(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}

	return end - start
}
