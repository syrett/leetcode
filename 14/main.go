package main

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
*/

func longestCommonPrefix(strs []string) string {
	lenStr := len(strs)
	if lenStr < 1 {
		return ""
	}
	prefix := strs[0]
	for _, item := range strs[1:] {
		if len(prefix) < 1 || len(item) < 1 {
			return ""
		}
		if item[0] != prefix[0] {
			return ""
		}

		if item == prefix {
			continue
		}

		lenPrefix := len(prefix)
		lenItem := len(item)
		i := 0
		isContinue := true
		for isContinue {
			if i > lenItem-1 || i > lenPrefix-1 {
				isContinue = false
				continue
			}

			if prefix[i] != item[i] {
				isContinue = false
				continue
			}
			i++
		}
		prefix = prefix[0:i]
	}

	return prefix

}
