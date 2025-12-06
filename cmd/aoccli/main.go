package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"plugin"

	"github.com/tomsobpl/aoc/pkg/core"
)

func loadPlugin(input core.AocInput) func() core.AocSolution {
	p, err := plugin.Open(fmt.Sprintf("%s/year%d-day%d-plugin.so", pluginsPath, input.Year(), input.Day()))
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("NewSolution")
	if err != nil {
		log.Fatal(err)
	}

	solution, ok := f.(func() core.AocSolution)
	if !ok {
		log.Fatal("unexpected type from module symbol")
	}

	return solution
}

var pluginsPath string

func init() {
	pluginsPath = os.Getenv("AOC_PLUGINS_PATH")
}

func main() {
	solveCommand := flag.NewFlagSet("solve", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("TODO: print usage")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "solve":
		solveCommand.Parse(os.Args[2:])
		args := solveCommand.Args()

		if len(args) < 1 {
			fmt.Println("Input file path is required")
			os.Exit(1)
		}

		yamlBytes, err := os.ReadFile(args[0])

		if err != nil {
			log.Fatal(err)
		}
		input, err := core.NewAocInputFromYaml(yamlBytes)

		var NewSolution func() core.AocSolution = loadPlugin(input)
		task := core.NewAocTask(NewSolution())

		result := task.Solve(input)
		fmt.Printf("Result: %v\n", result)
	default:
		fmt.Println("TODO: print usage")
		os.Exit(1)
	}
}
