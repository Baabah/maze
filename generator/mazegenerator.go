package generator

import "github.com/Baabah/maze/model"

func BuildMaze() model.Maze {
	maze := model.Maze{[][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', ' ', '#', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', '#'},
		{'#', ' ', '#', '#', '#', '#', ' ', '#', ' ', ' ', '#', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', '#', ' ', '#', '#', ' ', '#', ' ', '#'},
		{'#', ' ', '#', '#', ' ', '#', ' ', ' ', '#', ' ', ' ', ' ', '#'},
		{'#', '#', '#', '#', '*', '#', '#', '#', '#', '#', '#', '#', '#'},
	}}
	return maze
}