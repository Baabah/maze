package solver

import (
	"fmt"
	"errors"
	"math/rand"
	"github.com/Baabah/maze/model"
	"time"
)

func Solve(maze model.Maze) (model.Solution, error) {
	startPos, errorResponse := findStart(maze)
	if errorResponse != nil {
		fmt.Println(errorResponse)
		return model.Solution{}, errorResponse
	}

	rand.Seed(time.Now().UTC().UnixNano())
	var steps []model.Position
	for {
		possibleNewPositions := findPathOptions(maze, startPos)
		newPos := possibleNewPositions[rand.Intn(len(possibleNewPositions))]
		steps = append(steps, newPos)
		if maze.IsEnd(newPos) == true {
			solution := model.Solution{steps}
			solutionWithoutDuplicates := removeDuplicates(solution)
			return solutionWithoutDuplicates, nil
		}
		startPos = newPos
	}
}

func findPathOptions(maze model.Maze, startPos model.Position) []model.Position {
	var possibleNewPositions []model.Position
	possiblePositions := []model.Position{
		{startPos.X, startPos.Y - 1}, // 1 up
		{startPos.X, startPos.Y + 1}, // 1 down
		{startPos.X - 1, startPos.Y}, // 1 left
		{startPos.X + 1, startPos.Y}, // 1 right
	}

	for _, possiblePosition := range possiblePositions {
		if maze.Contains(possiblePosition) {
			if maze.GetValueByPosition(possiblePosition) == ' ' {
				possibleNewPositions = append(possibleNewPositions, possiblePosition)
			}
		}
	}
	return possibleNewPositions
}

func findStart(maze model.Maze) (model.Position, error) {
	for y, row := range maze.Points {
		for x, element := range row {
			if element == '*' {
				position := model.Position{x, y}
				return position, nil;
			}
		}
	}
	return model.Position{}, errors.New("No start found")
}

func removeDuplicates(solution model.Solution) model.Solution {
	var solutionWithoutDuplicates model.Solution
	for _, position := range solution.Steps {
		if solutionWithoutDuplicates.Contains(position) {
			solutionWithoutDuplicates.RemoveAfter(position)
			continue
		}
		solutionWithoutDuplicates.Steps = append(solutionWithoutDuplicates.Steps, position)
	}
	return solutionWithoutDuplicates
}