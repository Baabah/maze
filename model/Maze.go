package model

import "fmt"

type Maze struct {
	Points [][]rune
}

func (maze *Maze) GetValueByPosition(position Position) rune {
	return maze.Points[position.Y][position.X]
}

func (maze *Maze) Contains(position Position) bool {
	for y, row := range maze.Points {
		if y == position.Y {
			for x := range row {
				if x == position.X {
					return true
				}
			}
		}
	}
	return false
}

func (maze *Maze) IsEnd(position Position) bool {
	if (position.Y == len(maze.Points)) {
		return true
	}
	if position.Y == 0 {
		return true
	}
	if (position.X == len(maze.Points[0])) {
		return true
	}
	if position.X == 0 {
		return true
	}
	return false
}

func (maze *Maze) Print() {
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
