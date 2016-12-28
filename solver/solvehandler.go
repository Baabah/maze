package solver

import (
	"github.com/Baabah/maze/model"
	"time"
	"fmt"
)

func Handle(maze model.Maze)  {

	dataChannel := make(chan model.Solution)
	for i := 0; i < 300; i++ {
		go worker(maze, dataChannel)
	}

	var best model.Solution
	startTime := time.Now().UTC().UnixNano()
	solutionLimit := 100
	solutionCount := 0
	for {
		if (solutionCount == solutionLimit) {
			printSolution(best, startTime)
			return
		}
		solutionCount++
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
		solution, errorResponse := Solve(maze)
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
	timeDiff := float64(endTime - startTime) / float64(time.Millisecond)
	fmt.Println("Milliseconds taken to solve:", timeDiff)
}