package handlers

import (
	"encoding/json"
	"net/http"

	"dfs/models"
	"dfs/service"
)

func CalculatePathHandler(w http.ResponseWriter, r *http.Request) {
	var request models.GridPathRequest
	var response models.GridPathResponse

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(request.Grid) == 0 {
		http.Error(w, "Grid is empty", http.StatusBadRequest)
		return
	}

	if len(request.Start) != 2 || len(request.End) != 2 {
		http.Error(w, "Invalid start or end point", http.StatusBadRequest)
		return
	}

	sx, sy := request.Start[0], request.Start[1]
	ex, ey := request.End[0], request.End[1]

	rows := len(request.Grid)
	columns := len(request.Grid[0])

	if sx < 0 || sx >= rows || sy < 0 || sy >= columns {
		http.Error(w, "Invalid start point", http.StatusBadRequest)
		return
	}

	if ex < 0 || ex >= rows || ey < 0 || ey >= columns {
		http.Error(w, "Invalid end point", http.StatusBadRequest)
		return
	}

	Path := service.CalculateThePath(request.Grid, sx, sy, ex, ey)

	if len(Path) > 1 {
		response.Path = Path
	} else {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
