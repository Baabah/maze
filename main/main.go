package main

import (
	"github.com/Baabah/maze/generator"
	"github.com/Baabah/maze/solver"
	"github.com/Baabah/maze/model"
	"fmt"
)

// todo add solution model

func main() {
	var maze = generator.BuildMaze()
	dataChannel := make(chan []model.Position)
	for i := 0; i < 30; i++ {
		go worker(maze, dataChannel)
	}

	var best []model.Position
	for {
		solution := <- dataChannel
		if best == nil {
			best = solution
			fmt.Println(best)
			fmt.Println("Best solution is:", len(best))
			continue
		}
		if len(solution) < len(best) {
			best = solution
			fmt.Println(best)
			fmt.Println("Best solution is:", len(best))
		}
	}
}

func worker(maze model.Maze, dataChannel chan []model.Position) {
	for {
		solution := solver.Solve(maze)
		dataChannel <- solution
	}
}