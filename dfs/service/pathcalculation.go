package service

import "dfs/utils"

var direction = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func findPath(grid [][]int, x, y, ex, ey int, visitedpath [][]bool, path *[][]int) bool {

	if x == ex && y == ey {
		*path = append(*path, []int{x, y})
		return true
	}

	visitedpath[x][y] = true
	*path = append(*path, []int{x, y})

	for _, d := range direction {
		nx, ny := x+d[0], y+d[1]
		if utils.IsValidGrid(grid, nx, ny) && grid[nx][ny] == 1 && !visitedpath[nx][ny] {
			if findPath(grid, nx, ny, ex, ey, visitedpath, path) {
				return true
			}
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func CalculateThePath(grid [][]int, sx, sy, ex, ey int) [][]int {
	visitedpath := make([][]bool, len(grid))
	for i := range visitedpath {
		visitedpath[i] = make([]bool, len(grid[0]))
	}
	path := [][]int{}
	findPath(grid, sx, sy, ex, ey, visitedpath, &path)
	return path
}
