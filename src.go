package main

import (
	"math"
	"math/rand"
	"time"
)

func MathFunc(x float64) float64 {
	return math.Sin(x) + 0.5
}

func MathFuncPoints(xPoints []float64) []float64 {
	var yPoints []float64
	for _, point := range xPoints {
		yPoints = append(yPoints, MathFunc(point))
	}
	return yPoints
}

func noiseMathFunc(x float64) float64 {
	rand.Seed(time.Now().UnixNano())
	min := -(0.5 / 2)
	max := 0.5 / 2
	r := min + rand.Float64() * (max - min)
	return math.Sin(x) + 0.5 + r
}

func noiseMathFuncPoints(xPoints []float64) (yPoints []float64) {
	for _, point := range xPoints {
		yPoints = append(yPoints, noiseMathFunc(point))
	}
	return yPoints
}
