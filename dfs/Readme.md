# Calculate Shortest Path

# Project Description

`Build a full-stack application where the user interacts with a 20x20 grid to select two tiles (start and end), and the shortest path between them is calculated using Depth-First Search (DFS). The path will be calculated on the backend (Go) and returned to the frontend, where it will be displayed by highlighting the path in a different color.

Frontend Requirements:

1. Grid Layout:
   - Create a 20x20 grid (400 cells in total) using a frontend framework like React.js.
   - Each grid cell should be clickable.
2. Tile Selection:

   - Allow the user to select exactly two tiles: one for the start point and one for the end point.
   - The selected start and end tiles should be visually distinguishable from other cells.

3. Path Display:
   - Once the shortest path is calculated by the backend, highlight the path between the start and end tiles using a different color (e.g., blue).
4. API Interaction:
   - When the start and end points are selected, send their grid coordinates (e.g., (x1, y1) and (x2, y2)) to the backend via an HTTP request.
   - Receive the calculated path from the backend as a series of grid coordinates and display the path by coloring the appropriate cells.

Backend Requirements (Golang):

1. API Setup:

   - Implement a Go server to handle requests from the frontend.
   - Create an API endpoint (e.g., `/find-path`) to receive the start and end coordinates and respond with the calculated path.

2. DFS Algorithm:
   - Implement Depth-First Search (DFS) to calculate the shortest path between the selected start and end points on the grid.
   - Ensure the DFS algorithm is capable of handling a 2D grid.
3. Response Format:
   - Return the shortest path as a list of grid coordinates from the start to the end tile.

---

Deliverables:

1. Frontend:
   - A functional 20x20 grid interface that allows tile selection and path visualization.
2. Backend:

   - A Go server capable of calculating the shortest path between two points on the grid using DFS.
   - The backend should return the calculated path as a list of coordinates to the frontend.

3. Integration:
   - Ensure the frontend and backend communicate effectively. Upon selecting the start and end tiles, the frontend should send the data to the backend, and once the path is calculated, it should be visualized on the frontend.

Extra Features (Optional):

- Add obstacles to the grid, where the path must navigate around them.
- Allow the user to reset the grid and try different start and end points.

Note: Write read.me and add screenshots and explain what you used to solve this problem.
`

# Curl

curl --location 'http://localhost:8080/find-path' \
--header 'Content-Type: application/json' \
--data '{
"grid":[
[1,1,0,0,0],
[1,1,0,1,1],
[0,1,1,1,0],
[0,0,0,1,1],
[1,1,1,1,1]

    ],
    "start":[0,0],
    "end":[4,4]

}'

# Response

`{
    "path": [
        [
            0,
            0
        ],
        [
            0,
            1
        ],
        [
            1,
            1
        ],
        [
            2,
            1
        ],
        [
            2,
            2
        ],
        [
            2,
            3
        ],
        [
            3,
            3
        ],
        [
            3,
            4
        ],
        [
            4,
            4
        ]
    ]
}`

# Run command

go mod init dfs
go mod tidy
go run main.go
