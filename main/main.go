package main

import (
	"github.com/Baabah/maze/generator"
	"github.com/Baabah/maze/solver"
	"github.com/Baabah/maze/model"
	"fmt"
	"time"
)

func main() {
	var maze = generator.BuildMaze()
	dataChannel := make(chan model.Solution)
	for i := 0; i < 300; i++ {
		go worker(maze, dataChannel)
	}

	var best model.Solution
	startTime := time.Now().UTC().UnixNano()
	for {
		solution := <- dataChannel
		if best.Steps == nil {
			best = solution
			printSolution(best, startTime)
			continue
		}
		if len(solution.Steps) < len(best.Steps) {
			best = solution
			printSolution(best, startTime)
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

func printSolution(solution model.Solution, startTime int64) {
	fmt.Println(solution.Steps)
	fmt.Println("Best solution is:", len(solution.Steps))
	endTime := time.Now().UTC().UnixNano()
	fmt.Println("Nanoseconds taken to solve:", endTime - startTime)
}