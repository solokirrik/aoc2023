package main

import "math"

func NewVector(p1, p2 point) Vector {
	return Vector{float64(p2.x) - float64(p1.x), float64(p2.y) - float64(p1.y)}
}

// Vector represents a 2D vector with components x and y.
type Vector struct {
	X, Y float64
}

// Magnitude calculates the magnitude of a vector.
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// DotProduct calculates the dot product of two vectors.
func DotProduct(v, u Vector) float64 {
	return v.X*u.X + v.Y*u.Y
}

// AngleBetweenVectors calculates the angle in radians between two vectors.
func AngleBetweenVectors(v, u Vector) float64 {
	dotProduct := DotProduct(v, u)
	magnitudeV := v.Magnitude()
	magnitudeU := u.Magnitude()

	// Avoid division by zero
	if magnitudeV == 0 || magnitudeU == 0 {
		return 0
	}

	cosTheta := dotProduct / (magnitudeV * magnitudeU)
	return math.Acos(cosTheta)
}
