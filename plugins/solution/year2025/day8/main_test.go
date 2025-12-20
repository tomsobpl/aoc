package main

import (
	_ "embed"
	"testing"
)

//go:embed main_testdata.txt
var inputData string

func TestPartOneSolution(t *testing.T) {
	want := "40"
	have := Solution{}.solvePartOne(loadInput(), 10)

	if have.Result() != want {
		t.Errorf("Expected %s, got %s", want, have.Result())
	}
}

func TestPartTwoSolution(t *testing.T) {
	want := "25272"
	have := Solution{}.solvePartTwo(loadInput())

	if have.Result() != want {
		t.Errorf("Expected %s, got %s", want, have.Result())
	}
}

func loadInput() string {
	return inputData
}
