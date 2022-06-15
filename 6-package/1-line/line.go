package main

import "math"

// if you init with uppercase will be public to another package
// if you init with lowercase will be private in the package

//Point - something public
//point or _Point - something private

// point represents a coordinate in the Cartesian plane
type Point struct {
	x float64
	y float64
}

func side(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancy is responsible for calculate the distancy between two points
func Distancy(a, b Point) float64 {
	cx, cy := side(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
