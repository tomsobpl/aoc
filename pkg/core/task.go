package core

const (
	AocTaskPartOne int = 1
	AocTaskPartTwo int = 2
)

type AocTask interface {
	Solve(input AocInput) AocResult
}

type aocTask struct {
	solution AocSolution
}

func (task aocTask) Solve(input AocInput) AocResult {
	return task.solution.Solve(input)
}

func NewAocTask(solution AocSolution) AocTask {
	return aocTask{solution: solution}
}
