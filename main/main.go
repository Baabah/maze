package main

import (
	"github.com/Baabah/maze/generator"
	"github.com/Baabah/maze/solver"
	"github.com/Baabah/maze/model"
	"fmt"
)

func main() {
	var maze = generator.BuildMaze()
	dataChannel := make(chan model.Solution)
	for i := 0; i < 30; i++ {
		go worker(maze, dataChannel)
	}

	var best model.Solution
	for {
		solution := <- dataChannel
		if best.Steps == nil {
			best = solution
			fmt.Println(best.Steps)
			fmt.Println("Best solution is:", len(best.Steps))
			continue
		}
		if len(solution.Steps) < len(best.Steps) {
			best = solution
			fmt.Println(best.Steps)
			fmt.Println("Best solution is:", len(best.Steps))
		}
	}
}

func worker(maze model.Maze, dataChannel chan model.Solution) {
	for {
		solution, errorResponse := solver.Solve(maze)
		if errorResponse == nil {
			dataChannel <- solution
			continue
		}
		fmt.Println(errorResponse)
	}
}