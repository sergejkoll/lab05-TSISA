package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

type minEl struct {
	weight float64
	J float64
	W float64
	D float64
	alpha []float64
	distance float64
}

func getMinWeight(inputData map[string]float64, noiseFunc []float64) []float64 {
	N := math.Log(1 - inputData["prob"])/
		math.Log(1 - (inputData["eps"]/(inputData["xMax"] - inputData["xMin"])))
	M := (inputData["r"] - 1) / 2
	yFilteredPoints := make(map[float64][]float64)
	minElSlice := make([]minEl, 0)
	for weight := 0.0; weight <= 1.0; weight += 0.1 {
		var Jmin float64
		Jmin = 1000
		var noisinessMin float64
		var differencesMin float64
		var alphaMin []float64
		for i := 0; i < int(N); i++ {
			filterFunc := make([]float64, 100)
			alpha := getAlpha(inputData["r"])
			for k := M; k < 100 - M; k++ {
				var multiplication float64
				multiplication = 1
				for j := k - M; j < k + M; j++ {
					multiplication *= math.Pow(noiseFunc[int(j)], alpha[int(j + M + 1 - k)])
				}
				filterFunc[int(k)] = multiplication
			}
			noisiness := getNoisiness(filterFunc)
			differences := getDifferences(filterFunc, noiseFunc)
			J := getJ(weight, noisiness, differences)
			if J < Jmin {
				Jmin = J
				noisinessMin = noisiness
				differencesMin = differences
				alphaMin = alpha
				yFilteredPoints[weight] = filterFunc
			}
		}
		distance := math.Abs(noisinessMin) + math.Abs(differencesMin)
		minElSlice = append(minElSlice, minEl{weight, Jmin, noisinessMin, differencesMin, alphaMin, distance})
	}
	sort.Slice(minElSlice, func(i, j int) bool {
		return minElSlice[i].distance < minElSlice[j].distance
	})

	return yFilteredPoints[minElSlice[0].weight]
}

func getNoisiness(filterFunc []float64) float64 {
	var sum float64
	for k := 1; k < 100; k++ {
		sum += math.Abs(filterFunc[k] - filterFunc[k - 1])
	}
	return sum
}

func getDifferences(filterFunc []float64, noiseFunc []float64) float64 {
	var differences float64
	for k := 0; k < 100; k++ {
		differences += math.Abs(filterFunc[k] - noiseFunc[k])
	}
	return differences/100
}

func getJ(weight float64, noisiness float64, differences float64) float64 {
	return weight * noisiness + (1 - weight) * differences
}

func randInInterval(max float64, min float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64() * (max - min)
}

func getAlpha(r float64) []float64 {
	alpha := make([]float64, int(r))
	M := int((r - 1)/2)
	rand.Seed(time.Now().UnixNano())
	// rand [0;1)
	alpha[M] = rand.Float64()
	sumAll := alpha[M]
	for i := M - 1; i > 0; i--{
		var sum float64
		for k := i + 1; k < int(r) - i - 1; k++ {
			sum += alpha[k]
		}
		alpha[i] = 0.5 * randInInterval(0, 1 - sum)
		alpha[int(r) - i - 1] = alpha[i]
		sumAll += 2 * alpha[i]
	}
	alpha[0] = 0.5 * (1 - sumAll)
	alpha[int(r) - 1] = alpha[0]
	return alpha
}


