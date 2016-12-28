package main

import (
	"github.com/Baabah/maze/generator"
	"github.com/Baabah/maze/solver"
)

func main() {
	var maze = generator.BuildMaze()
	solver.Handle(maze)
}