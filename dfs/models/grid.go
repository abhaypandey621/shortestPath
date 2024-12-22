package models

type GridPathRequest struct {
	Grid  [][]int `json:"grid"`
	Start []int   `json:"start"`
	End   []int   `json:"end"`
}

type GridPathResponse struct {
	Path [][]int `json:"path"`
}
