package pkg

import (
	"strconv"
)

func MaxSumPath(triangle [][]int) string {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if triangle[i+1][j] > triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}
		}
	}
	result := strconv.Itoa(triangle[0][0])

	return result
}
