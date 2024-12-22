# Shortest Path Finder Backend (DFS)

## vProject Description

This is the backend implementation of a full-stack application that calculates the shortest path between two points on a 20x20 grid using the Depth-First Search (DFS) algorithm. It provides an API to receive the start and end points, calculates the path, and returns it as a list of grid coordinates.

# Backend Requirements:

Go server that exposes an API to calculate the shortest path using DFS.
DFS Algorithm implemented to find the shortest path between two points on the grid.
The backend should return the calculated path as a list of grid coordinates.

## API Endpoint

POST /find-path
Request:

The frontend sends a POST request to the /find-path endpoint with the following payload:

`{
  "grid": [
    [1, 1, 0, 0, 0],
    [1, 1, 0, 1, 1],
    [0, 1, 1, 1, 0],
    [0, 0, 0, 1, 1],
    [1, 1, 1, 1, 1]
  ],
  "start": [0, 0],
  "end": [4, 4]
}`

grid: A 2D array representing the grid where 1 represents an open cell and 0 represents a blocked cell.
start: The starting point coordinates [x1, y1].
end: The ending point coordinates [x2, y2].

## Response:

The server will return a JSON response containing the path from the start point to the end point:

`{
  "path": [
    [0, 0],
    [0, 1],
    [1, 1],
    [2, 1],
    [2, 2],
    [2, 3],
    [3, 3],
    [3, 4],
    [4, 4]
  ]
}
`

path: An array of coordinates representing the shortest path between the start and end points.

# Running the Backend

1. Initialize Go Module
   To initialize the Go module for the project, run the following command:

`go mod init dfs`

`go mod tidy`

`go run main.go`

# 4. Test the API

To test the backend functionality, you can use curl or any HTTP client like Postman. Here's an example using curl:

`curl --location 'http://localhost:8080/find-path' \
--header 'Content-Type: application/json' \
--data '{
  "grid": [
    [1, 1, 0, 0, 0],
    [1, 1, 0, 1, 1],
    [0, 1, 1, 1, 0],
    [0, 0, 0, 1, 1],
    [1, 1, 1, 1, 1]
  ],
  "start": [0, 0],
  "end": [4, 4]
}'`

# DFS Algorithm (Backend)

The DFS algorithm is implemented in the backend to search for the shortest path between the start and end coordinates in the grid.

Key Points:
DFS Search explores paths from the start point, checking neighboring cells recursively.
It ensures to avoid revisiting already visited cells or blocked cells (cells with 0).
The algorithm backtracks once a valid path to the end point is found.
