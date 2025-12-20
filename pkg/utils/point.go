package utils

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p Point) DistanceTo(other Point) float32 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	return float32(math.Sqrt(dx*dx + dy*dy))
}

type Point3D struct {
	X int
	Y int
	Z int
}

func (p Point3D) DistanceTo(other Point3D) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	dz := float64(p.Z - other.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (p Point3D) Slug() string {
	return fmt.Sprintf("%d-%d-%d", p.X, p.Y, p.Z)
}

func Point3DFromSlice(input []int) Point3D {
	return Point3D{input[0], input[1], input[2]}
}
