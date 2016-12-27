package main

import (
	"github.com/Baabah/maze/generator"
	"github.com/Baabah/maze/solver"
	"github.com/Baabah/maze/model"
	"fmt"
)

func main() {
	var maze = generator.BuildMaze()
	var solutions [][]model.Position
	for i := 0; i < 5000; i++ {
		solutions = append(solutions, solver.Solve(maze))
	}
	var best []model.Position
	for _, solution := range solutions {
		if best == nil {
			best = solution
			continue
		}
		if len(solution) < len(best) {
			best = solution
		}
	}
	fmt.Println(best)
	fmt.Println("Best solution is:", len(best))

}