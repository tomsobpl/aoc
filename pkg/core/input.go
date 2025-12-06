package core

import "gopkg.in/yaml.v3"

type AocInput interface {
	Year() int
	Day() int
	Part() int
	Payload() string
}

type aocInput struct {
	Year_    int    `yaml:"year"`
	Day_     int    `yaml:"day"`
	Part_    int    `yaml:"part"`
	Payload_ string `yaml:"payload"`
}

func (a aocInput) Year() int {
	return a.Year_
}

func (a aocInput) Day() int {
	return a.Day_
}

func (a aocInput) Part() int {
	return a.Part_
}

func (a aocInput) Payload() string {
	return a.Payload_
}

func NewAocInput(year int, day int, part int, payload string) AocInput {
	return aocInput{
		Year_:    year,
		Day_:     day,
		Part_:    part,
		Payload_: payload,
	}
}

func NewAocInputFromYaml(raw []byte) (AocInput, error) {
	var input aocInput
	err := yaml.Unmarshal(raw, &input)
	return input, err
}
