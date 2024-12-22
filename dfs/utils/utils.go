package utils

func IsValidGrid(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}
