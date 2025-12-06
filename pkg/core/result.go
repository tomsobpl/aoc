package core

type AocResult interface {
	Result() string
}

type aocResult struct {
	result string
}

func (r aocResult) Result() string { return r.result }

func NewAocResult(result string) AocResult {
	return aocResult{result: result}
}
