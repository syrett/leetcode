package main

import (
	"fmt"
	"strconv"
)

/*
7. 整数反转

给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0


提示：

-231 <= x <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

func reverse(x int) int {
	str := fmt.Sprintf("%d", x)
	reverseStr := ""
	isNegative := false
	for _, i := range str {
		if i == 45 {
			isNegative = true
		} else {
			reverseStr = fmt.Sprintf("%c%s", i, reverseStr)
		}
	}
	if isNegative {
		reverseStr = fmt.Sprintf("-%s", reverseStr)
	}
	v, _ := strconv.ParseInt(reverseStr, 10, 64)

	if v <= (2<<30)-1 && v >= (-2)<<30 {
		return int(v)
	}
	return 0
}
