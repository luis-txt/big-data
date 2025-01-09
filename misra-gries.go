package main

func misraGries(stream []float64, k int) int {
	var candidate float64
	count := 0

	// Identify potential candidate
	for _, value := range stream {
		if count == 0 {
			candidate = value
			count = 1
		} else if value == candidate {
			count++
		} else {
			count--
		}
	}

	// Verify candidate
	count = 0
	for _, value := range stream {
		if value == candidate {
			count++
		}
	}
	if count > len(stream)/k {
		return int(candidate)
	}
	return -1
}

func misraGriesExtended(stream []float64, k int) map[float64]int {
	candidateCounts := make(map[float64]int)

	// Identify potential candidates
	for _, value := range stream {
		if _, exists := candidateCounts[value]; exists {
			candidateCounts[value]++
		} else if len(candidateCounts) < k-1 {
			candidateCounts[value] = 1
		} else {
			for key := range candidateCounts {
				candidateCounts[key]--
				if candidateCounts[key] == 0 {
					delete(candidateCounts, key)
				}
			}
		}
	}

	// Count actual occurrences
	for key := range candidateCounts {
		candidateCounts[key] = 0
	}
	for _, value := range stream {
		if _, exists := candidateCounts[value]; exists {
			candidateCounts[value]++
		}
	}

	// Filter out invalid candidates
	threshold := len(stream) / k
	for key, count := range candidateCounts {
		if count <= threshold {
			delete(candidateCounts, key)
		}
	}
	return candidateCounts
}
