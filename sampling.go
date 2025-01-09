package main

import (
	"math"
	"math/rand"
)

func bernoulliSampling(stream []float64, q float64) []float64 {
	sample := make([]float64, 0, int(float64(len(stream))*q))

	for _, value := range stream {
		if rand.Float64() <= q {
			sample = append(sample, value)
		}
	}
	return sample
}

func improvedBernoulliSampling(stream []float64, q float64) []float64 {
	sample := make([]float64, 0, int(float64(len(stream))*q))

	i := 0

	for i < len(stream) {
		skips := int(math.Log(rand.Float64()) / math.Log(1-q))
		i += skips
		if i < len(stream) {
			sample = append(sample, stream[i])
			i++
		}
	}
	return sample
}

func reservoirSampling(stream []float64, k int) []float64 {
	reservoir := make([]float64, 0, k)

	// Fill reservooir initially
	for i := 0; i < k && i < len(stream); i++ {
		reservoir = append(reservoir, stream[i])
	}

	// Sample
	for i := k; i < len(stream); i++ {
		j := rand.Intn(i + 1)
		if j < k {
			reservoir[j] = stream[i]
		}
	}
	return reservoir
}
