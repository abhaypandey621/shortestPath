package main

import (
	"dfs/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/find-path", handlers.CalculatePathHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
