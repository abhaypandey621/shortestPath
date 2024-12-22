Shortest Path Finder (DFS)
Project Description
This project is a full-stack application that calculates the shortest path between two points on a 20x20 grid using the Depth-First Search (DFS) algorithm. The grid is displayed on the frontend using React.js, and the shortest path is calculated and returned by a Go backend. The path is then visually represented on the frontend.

Key Features:
Interactive Grid Layout: A 20x20 grid with clickable cells.
Tile Selection: Users can select a start and end point, and the backend calculates the shortest path.
Path Visualization: The path is displayed on the frontend with a different color once calculated by the backend.
API Interaction: The frontend communicates with the backend via HTTP requests to calculate and display the shortest path.
Frontend Requirements
Grid Layout:

Create a 20x20 grid (400 cells) using a frontend framework such as React.js.
Each grid cell is clickable and can be selected as either the start or end point.
Tile Selection:

Allow the user to select exactly two tiles: one for the start point and one for the end point.
These selected tiles should be visually distinguishable from the other cells.
Path Display:

Once the shortest path is calculated, highlight the path between the start and end points with a different color (e.g., blue).
API Interaction:

Upon selecting the start and end points, send the coordinates to the backend via an HTTP request.
Receive the calculated path from the backend, which will be a list of grid coordinates.
Visualize the path by coloring the appropriate cells in the grid.
Backend Requirements (Golang)
API Setup:

Implement a Go server to handle requests from the frontend.
Create an API endpoint (e.g., /find-path) that receives the start and end coordinates and responds with the calculated path.
DFS Algorithm:

Implement the Depth-First Search (DFS) algorithm to calculate the shortest path between the selected start and end points on the grid.
Ensure the DFS algorithm can handle a 2D grid.
Response Format:

The backend should return the shortest path as a list of grid coordinates from the start to the end tile.
API Example
Request (via curl)
To test the backend, you can use curl to send a POST request to the /find-path endpoint with a grid, start, and end points:

bash
Copy code
curl --location 'http://localhost:8080/find-path' \
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
}'
Example Response
The backend will respond with the calculated path in the following format:

json
Copy code
{
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
Running the Application
Backend (Go)
Initialize the Go Module: First, initialize the Go module if you haven't already:

bash
Copy code
go mod init dfs
Install Dependencies: Run the following command to tidy up the Go module and install dependencies:

bash
Copy code
go mod tidy
Run the Backend: To start the Go backend, run the following command:

bash
Copy code
go run main.go
The backend will start running on http://localhost:8080.
