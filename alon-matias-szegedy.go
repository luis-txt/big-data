package main

import "math/rand"

func secondFrequencyMoment(stream []float64) float64 {
	if len(stream) == 0 {
		return 0.0
	}
	frequencies := make(map[float64]int)

	for _, value := range stream {
		if _, exists := frequencies[value]; exists {
			frequencies[value]++
		} else {
			frequencies[value] = 1
		}
	}

	f2 := 0.0
	for _, frequency := range frequencies {
		f2 += float64(frequency * frequency)
	}
	return f2
}

func alonMatiasSzegedy(stream []float64, k int) float64 {
	if len(stream) == 0 || k <= 0 {
		return 0.0
	}

	estimates := make([]float64, k)
	n := len(stream)

	for i := 0; i < k; i++ {
		randomIndex := rand.Intn(n)
		sampledValue := stream[randomIndex]

		frequency := 0
		for _, value := range stream {
			if value == sampledValue {
				frequency++
			}
		}

		// Estimate F2 for sample
		estimates[i] = float64(frequency*frequency) * float64(n)
	}

	// Average estimates
	total := 0.0
	for _, estimate := range estimates {
		total += estimate
	}
	return total / float64(k)
}
