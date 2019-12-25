package main

import (
	"fmt"
	"math"
	"math/rand"

	"time"
)

type minEl struct {
	weight float64
	J float64
	W float64
	D float64
	alpha []float64
	yFiltered []float64
	distance float64
}

func coordinateSignal(K float64, input map[string]float64) []float64 {
	xPoints := make([]float64, int(K))
	var i float64
	for i = 0; i < 100; i++ {
		xPoints[int(i)] = input["xMin"] + (i * (input["xMax"] - input["xMin"]))/K
	}
	return xPoints
}

func getFiltered (input map[string]float64, noiseFunc []float64) minEl {
	minimal := minEl{
		distance: 1000,
	}
	filtered := make(map[float64]minEl)
	for lambda := 0.0; lambda <= 1; lambda += 0.1 {
		filtered[lambda] = getWeights(input, noiseFunc, lambda)
		fmt.Println("weight: ", lambda)
		fmt.Println("alpha: ", filtered[lambda].alpha)
		fmt.Println("W: ", filtered[lambda].W)
		fmt.Println("D: ", filtered[lambda].D)
		fmt.Println("J: ", filtered[lambda].J)
		fmt.Println("dist: ", filtered[lambda].distance)
		fmt.Println("-------------------------------------------------------------------------------------------------")
		if minimal.distance > filtered[lambda].distance {
			minimal = filtered[lambda]
		}
	}
	fmt.Println("MINIMAL")
	fmt.Println("weight: ", minimal.weight)
	fmt.Println("alpha: ", minimal.alpha)
	fmt.Println("W: ", minimal.W)
	fmt.Println("D: ", minimal.D)
	fmt.Println("J: ", minimal.J)
	fmt.Println("dist: ", minimal.distance)
	return minimal
}

func getWeights(input map[string]float64, noiseFunc []float64, lambda float64) minEl {
	N := math.Log(1 - input["P"])/
		 math.Log(1 - (input["e"]/(input["xMax"] - input["xMin"])))
	yFiltered := make([]float64, 101)
	min := minEl{
		J : 1000.1,
	}
	for i := 0; i <= int(N); i++ {
		alpha := getAlpha(input["r"])
		yFiltered = filteredSignal(alpha, input["K"], input["r"], noiseFunc)
		W := getNoisiness(yFiltered, input["K"])
		D := getDifferences(yFiltered, noiseFunc, input["K"])
		J := getJ(lambda, W, D)
		if min.J > J {
			min.J = J
			min.W = W
			min.D = D
			min.yFiltered = yFiltered
			min.alpha = alpha
			min.weight = lambda
			min.distance = math.Abs(min.W) + math.Abs(min.D)
		}
	}
	return min
}

func filteredSignal(alpha []float64, K float64, r float64, noiseFunc []float64) []float64 {
	M := (r - 1)/2
	yFiltered := make([]float64, int(K) + 1)
	for k := 0; k <= int(M); k++ {
		yFiltered[k] = noiseFunc[k]
		yFiltered[int(K) - k - 1] = noiseFunc[int(K) - k - 1]
	}
	var composition float64
	for k := M; k <= K - M; k++ {
		composition = 1
		for j := k - M; j <= k + M; j++ {
			composition *= math.Pow(noiseFunc[int(j)], alpha[int(j + M - k)])
		}
		yFiltered[int(k)] = composition
	}
	return yFiltered
}

func getNoisiness(filterFunc []float64, K float64) float64 {
	var noisiness float64
	for k := 1; k < int(K); k++ {
		noisiness += math.Abs(filterFunc[k] - filterFunc[k - 1])
	}
	return noisiness
}

func getDifferences(filterFunc []float64, noiseFunc []float64, K float64) float64 {
	var differences float64
	for k := 0; k < int(K); k++ {
		differences += math.Abs(filterFunc[k] - noiseFunc[k])
	}
	return differences/K
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


