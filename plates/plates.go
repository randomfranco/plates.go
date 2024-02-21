package main

import "sort"

// ComplessiveWeight is a struct containing the weight of the barbell and a map[float64]int of plates, in the map the weight is the key and the value is how many plates you got
type ComplessiveWeight struct {
	bar float64
	// collars weight (full weight of the pairs, ex 5kg)
	collars float64
	availablePlates map[float64]int
}

// sort the keys of the plates map in order to iterate them from bigger to smaller in a greedy fashion
func sortPlates(Plates map[float64]int) []float64{
	keys := make([]float64, 0, len(Plates))
	for k := range Plates {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	return keys
}

// greedy algorithm to select the right amount of weights
func (cw ComplessiveWeight) GreedyPlatesSelector(weight float64) map[float64]int{

	plates := make(map[float64]int)
	sortedPlates := sortPlates(cw.availablePlates)

	// set current weight to bar and only account for plate couples
	currentWeight := cw.bar + cw.collars
	for _, plate := range sortedPlates {
		for cw.availablePlates[plate] > 1 && (currentWeight + plate*2 <= weight) {
			plates[plate]+=2
			cw.availablePlates[plate]-=2
			currentWeight+=plate*2
		}

	}

	return plates
}

// sum of the weights take map[float]int and return float64
func PlatesSum(plates map[float64]int) float64{

	sum := 0.0

	for key, value := range plates{
		sum += key * float64(value)
	}

	return sum
}
