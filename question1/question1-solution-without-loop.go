package question1

func TriangleSumValueWithoutLoop(triangle [][]int) int {
	memo := make(map[[2]int]int)
	return getMaxPathSum(triangle, 0, 0, memo)
}

func getMaxPathSum(triangle [][]int, row, col int, memo map[[2]int]int) int {
	if row == len(triangle)-1 {
		return triangle[row][col]
	}

	if val, ok := memo[[2]int{row, col}]; ok {
		return val
	}

	leftPathSum := getMaxPathSum(triangle, row+1, col, memo)
	rightPathSum := getMaxPathSum(triangle, row+1, col+1, memo)

	result := triangle[row][col] + max(leftPathSum, rightPathSum)
	memo[[2]int{row, col}] = result

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
