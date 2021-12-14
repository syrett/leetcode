package main

func generateMatrix(n int) [][]int {
	max := n * n
	num := 1
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	top, bottom, left, right := 0, n-1, 0, n-1

	for num <= max {
		if left <= right {
			for i := left; i <= right; i++ {
				matrix[top][i] = num
				num++
			}
			top++
		}

		if top <= bottom {
			for i := top; i <= bottom; i++ {
				matrix[i][right] = num
				num++
			}
			right--
		}

		if right > left {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = num
				num++
			}
			bottom--
		}

		if bottom >= top {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}
