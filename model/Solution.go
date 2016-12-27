package model

type Solution struct {
	Steps []Position
}

func (solution *Solution) Contains(position Position) bool {
	for _, step := range solution.Steps {
		if step == position {
			return true
		}
	}
	return false
}

func (solution *Solution) RemoveAfter(position Position) {
	var stepsWithoutDuplicates []Position
	for _, step := range solution.Steps {
		if step == position {
			return
		}
		stepsWithoutDuplicates = append(stepsWithoutDuplicates, step)
	}
}
