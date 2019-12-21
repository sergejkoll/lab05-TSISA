package main

import (
	"math"
)

func main() {
	xPoints := Linspace(0, math.Pi, 101)
	yPoints := MathFuncPoints(xPoints)
	yNoisePoints := noiseMathFuncPoints(xPoints)

	inputData := make(map[string]float64)
	inputData["xMax"] = math.Pi
	inputData["xMin"] = 0
	inputData["weightsNum"] = 10
	inputData["prob"] = 0.95
	inputData["eps"] = 0.01
	inputData["r"] = 3

	yFiltered := getMinWeight(inputData, yNoisePoints)

	plotting(xPoints, yPoints, 101, "sin(x) + 0.5","functions.png",255, 0, 0)
	plotting(xPoints, yNoisePoints, 101, "noise","functions.png", 0, 255, 0)
	plotting(xPoints, yFiltered, 101, "filtered","functions.png", 0, 0, 255)

	inputData["r"] = 5

	yFiltered = getMinWeight(inputData, yNoisePoints)

	plotting(xPoints, yPoints, 101, "sin(x) + 0.5","functions5.png", 255, 0, 0)
	plotting(xPoints, yNoisePoints, 101, "noise","functions5.png", 0, 255, 0)
	plotting(xPoints, yFiltered, 101, "filtered","functions5.png", 0, 0, 255)
}
