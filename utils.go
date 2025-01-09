package main

import (
	"math/rand"
)

func generateRandomData(count int, min, max float64) []float64 {
	data := make([]float64, count)

	for i := 0; i < count; i++ {
		data[i] = min + rand.Float64()*(max-min)
	}
	return data
}
