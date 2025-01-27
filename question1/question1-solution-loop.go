package question1

func TriangleSumValueWithLoop(triangle [][]int) int {
	dpTriangle := make([][]int, len(triangle))

	for i := range triangle {
		dpTriangle[i] = make([]int, len(triangle[i]))
		copy(dpTriangle[i], triangle[i])
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			dpTriangle[row][col] += max(dpTriangle[row+1][col], dpTriangle[row+1][col+1])
		}
	}
	return dpTriangle[0][0]
}
