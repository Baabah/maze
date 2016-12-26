package generator

import "fmt"

func BuildMaze() Maze {
	maze := Maze{[][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', ' ', '#', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', '#'},
		{'#', ' ', '#', '#', '#', '#', ' ', '#', ' ', ' ', '#', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', '#', ' ', '#', '#', ' ', '#', ' ', '#'},
		{'#', ' ', '#', '#', ' ', '#', ' ', ' ', '#', ' ', ' ', ' ', '#'},
		{'#', '#', '#', '#', '*', '#', '#', '#', '#', '#', '#', '#', '#'},
	}}
	return maze
}

func PrintMaze(maze Maze) {
	for _, row := range maze.points {
		for index, element := range row {
			if (index == 0) {
				fmt.Println()
			}
			fmt.Print(fmt.Sprintf("%c", element))
		}
	}
}