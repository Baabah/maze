package main

import (
	"github.com/Baabah/maze/solver"
	"github.com/Baabah/maze/generator"
)

func main() {
	var maze = generator.BuildMaze()
	maze.Print()
	solver.Solve(maze)
}