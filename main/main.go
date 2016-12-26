package main

import "github.com/Baabah/maze/generator"

func main() {
	var maze = generator.BuildMaze()
	generator.PrintMaze(maze)
}
