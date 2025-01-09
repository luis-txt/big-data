package main

import (
	"math"
)

func naiveSum(stream []float64) float64 {
	sum := 0.0
	for _, value := range stream {
		sum += value
	}
	return sum
}

func kahanSum(stream []float64) float64 {
	sum := 0.0
	c := 0.0
	for _, value := range stream {
		y := value - c
		t := sum + y
		c = (t - sum) - y
		sum = t
	}
	return sum
}

func naiveMean(stream []float64) float64 {
	mean := 0.0
	n := 1
	for _, value := range stream {
		mean = float64((mean*float64(n))+value) / float64(n+1)
		n++
	}
	return mean
}

func kahanMean(stream []float64) float64 {
	mean := 0.0
	n := 1
	c := 0.0
	for _, value := range stream {
		sum := mean * float64(n)
		y := value - c
		t := sum + y
		c = (t - sum) - y
		mean = float64(t) / float64(n+1)
		n++
	}
	return mean
}

func naiveVariance(numbers []float64) float64 {
	sum := 0.0
	sumOfSquares := 0.0
	n := 0

	for _, value := range numbers {
		sumOfSquares += math.Pow(value, 2)
		sum += value
		n++
	}
	return (float64(sumOfSquares) - (float64(math.Pow(sum, 2)) / float64(n))) / float64(n)
}

func improvedVariance(numbers []float64) float64 {
	s := 0.0
	mean := 0.0
	n := 1

	for _, value := range numbers {
		newMean := (float64(n)*mean + value) / float64(n+1)
		s += (value - mean) * (value - newMean)
		n++
		mean = newMean
	}
	n--
	return float64(s) / float64(n)
}
