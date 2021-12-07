package main

import "strconv"

func addStrings(num1 string, num2 string) string {
    len1 := len(num1)
    len2 := len(num2)
    isTen := 0
    result := ""

    for i, j := len1-1, len2-1; i>=0 || j>=0 || isTen != 0;i,j = i-1,j-1{
        var x,y int
        if i >= 0{
			x = int(num1[i] - '0')
        }
        if j >= 0 {
			y = int(num2[j] - '0')			
        }
        tmpNum := x+y+isTen
        result = strconv.Itoa(tmpNum%10) + result
        isTen = tmpNum/10
		println(isTen, i, j, x, y)
    }
    return result
}