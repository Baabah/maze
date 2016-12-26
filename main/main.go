package main

import (
	"github.com/Baabah/maze/solver"
	"github.com/Baabah/maze/generator"
	"fmt"
	"github.com/Baabah/maze/model"
)

func main() {
	var maze = generator.BuildMaze()
	printMaze(maze)
	solver.Solve(maze)
}

func printMaze(maze model.Maze) {
	for _, row := range maze.Points {
		for index, element := range row {
			printableElement := fmt.Sprintf("%c", element)
			if (index == 0) {
				fmt.Println(printableElement)
			} else {
				fmt.Print(printableElement)
			}
		}
	}
}
