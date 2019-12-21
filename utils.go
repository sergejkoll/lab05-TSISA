package main

import "github.com/pkg/errors"

func Linspace(start, end float64, num int) []float64 {
	if num < 0 {
		panic(errors.Errorf("number of samples, %d, must be non-negative.", num))
	}
	result := make([]float64, num)
	step := (end - start) / float64(num-1)
	for i := range result {
		result[i] = start + float64(i)*step
	}
	return result
}
