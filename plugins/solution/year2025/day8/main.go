package main

import (
	"cmp"
	"slices"
	"strconv"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
	"github.com/tomsobpl/aoc/pkg/utils/math"
)

type CircuitPoint struct {
	point    utils.Point3D
	circuits []*Circuit
}

func (cp *CircuitPoint) assignToCircuit(c *Circuit) {
	cp.circuits = append(cp.circuits, c)
}

func (cp *CircuitPoint) distanceTo(cp2 *CircuitPoint) float64 {
	return cp.point.DistanceTo(cp2.point)
}

func (cp *CircuitPoint) equals(cp2 *CircuitPoint) bool {
	return cp.point.X == cp2.point.X && cp.point.Y == cp2.point.Y && cp.point.Z == cp2.point.Z
}

func newCircuitPoint(point utils.Point3D) *CircuitPoint {
	return &CircuitPoint{point: point}
}

type Circuit struct {
	points []*CircuitPoint
}

func (c *Circuit) addCircuitPoint(cp *CircuitPoint) {
	if !c.contains(cp) {
		c.points = append(c.points, cp)
	}

	cp.assignToCircuit(c)
}

func (c *Circuit) contains(cp *CircuitPoint) bool {
	return slices.ContainsFunc(c.points, func(cp2 *CircuitPoint) bool {
		return cp.equals(cp2)
	})
}

func (c *Circuit) mergeWith(c2 *Circuit) {
	for cpi := range c2.points {
		c2.points[cpi].circuits = []*Circuit{}
		c.addCircuitPoint(c2.points[cpi])
	}

	c2.points = []*CircuitPoint{}
}

func (c *Circuit) size() int {
	return len(c.points)
}

func newCircuit(points []*CircuitPoint) *Circuit {
	c := &Circuit{}

	for cpi := range points {
		c.addCircuitPoint(points[cpi])
	}

	return c
}

type Pair struct {
	a *CircuitPoint
	b *CircuitPoint
}

func (p Pair) distance() float64 {
	return p.a.distanceTo(p.b)
}

func (p Pair) points() []*CircuitPoint {
	return []*CircuitPoint{p.a, p.b}
}

type Solution struct {
}

func (s Solution) Solve(data core.AocInput) core.AocResult {
	switch data.Part() {
	case core.AocTaskPartOne:
		return s.solvePartOne(data.Payload(), 1000)
	case core.AocTaskPartTwo:
		return s.solvePartTwo(data.Payload())
	}

	return nil
}

func (s Solution) solvePartOne(payload string, limit int) core.AocResult {
	circuits := make([]*Circuit, 0)
	connections := 0
	_, pairs := s.preparePayload(payload)

	for _, pair := range pairs {
		if connections >= limit {
			break
		}

		connections++
		circuits = s.connectBoxesIntoCircuits(circuits, pair)
	}

	slices.SortFunc(circuits, func(a, b *Circuit) int {
		return cmp.Compare(b.size(), a.size())
	})

	result := 1

	for _, circuit := range circuits[:3] {
		result *= circuit.size()
	}

	return core.NewAocResult(strconv.Itoa(result))
}

func (s Solution) connectBoxesIntoCircuits(circuits []*Circuit, pair Pair) []*Circuit {
	if 0 == len(circuits) {
		return append(circuits, newCircuit(pair.points()))
	}

	for ci, circuit := range circuits {
		if circuit.contains(pair.a) && circuit.contains(pair.b) {
			continue
		}

		if circuit.contains(pair.a) && !circuit.contains(pair.b) {
			circuits[ci].addCircuitPoint(pair.b)

			if len(pair.b.circuits) > 1 {
				s.mergeCircuits(pair.b.circuits)
			}

			continue
		}

		if circuit.contains(pair.b) && !circuit.contains(pair.a) {
			circuits[ci].addCircuitPoint(pair.a)

			if len(pair.a.circuits) > 1 {
				s.mergeCircuits(pair.a.circuits)
			}

			continue
		}
	}

	return append(circuits, newCircuit(pair.points()))
}

func (s Solution) mergeCircuits(circuits []*Circuit) {
	if len(circuits) <= 1 {
		return
	}

	for ci := range circuits[1:] {
		circuits[0].mergeWith(circuits[ci+1])
	}
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	circuits := make([]*Circuit, 0)
	maxCircuitSize, pairs := s.preparePayload(payload)

	for _, pair := range pairs {
		circuits = s.connectBoxesIntoCircuits(circuits, pair)

		if maxCircuitSize == max(pair.a.circuits[0].size(), pair.b.circuits[0].size()) {
			return core.NewAocResult(strconv.Itoa(pair.a.point.X * pair.b.point.X))
		}
	}

	return core.NewAocResult(strconv.Itoa(-1))
}

func (s Solution) preparePayload(rawPayload string) (int, []Pair) {
	boxes := make([]*CircuitPoint, 0)
	pairs := make([]Pair, 0)

	for _, line := range utils.ConvertStringToNotEmptyLines(rawPayload) {
		boxes = append(boxes, newCircuitPoint(utils.Point3DFromSlice(utils.ConvertStringToInts(line, ","))))
	}

	for _, box := range math.CartesianProduct(boxes) {
		pairs = append(pairs, Pair{box[0], box[1]})
	}

	slices.SortFunc(pairs, func(a, b Pair) int {
		return cmp.Compare(a.distance(), b.distance())
	})

	return len(boxes), pairs
}

func NewSolution() core.AocSolution {
	return Solution{}
}
