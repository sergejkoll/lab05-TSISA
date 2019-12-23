package main

import (
	"fmt"
	"math"
)

func main() {
	inputData := make(map[string]float64)
	inputData["xMax"] = math.Pi
	inputData["xMin"] = 0
	inputData["K"] = 100
	inputData["L"] = 10
	inputData["P"]= 0.95
	inputData["e"] = 0.01
	inputData["r"] = 3

	xPoints := coordinateSignal(101, inputData)
	yPoints := MathFuncPoints(xPoints)
	yNoisePoints := noiseMathFuncPoints(xPoints)

	filtered := getFiltered(inputData, yNoisePoints)

	plotting(xPoints, yPoints, 100, "sin(x) + 0.5","functions.png",255, 0, 0)
	plotting(xPoints, yNoisePoints, 100, "noise","functions.png", 0, 255, 0)
	plotting(xPoints, filtered.yFiltered, 100, "filtered","functions.png", 0, 0, 255)

	fmt.Println("\n-------------r = 5-------------")

	inputData["r"] = 5

	xPoints = coordinateSignal(101, inputData)
	yPoints = MathFuncPoints(xPoints)
	yNoisePoints = noiseMathFuncPoints(xPoints)

	filtered = getFiltered(inputData, yNoisePoints)

	plotting5(xPoints, yPoints, 100, "sin(x) + 0.5","functions5.png",255, 0, 0)
	plotting5(xPoints, yNoisePoints, 100, "noise","functions5.png", 0, 255, 0)
	plotting5(xPoints, filtered.yFiltered, 100, "filtered","functions5.png", 0, 0, 255)
}
