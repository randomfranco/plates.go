package main

import "sort"


// sort the keys of the plates map in order to iterate them from bigger to smaller in a greedy fashion
func sortPlates(Plates PlateSet_t) []float64{
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
func (cw ComplessiveWeight) GreedyPlatesSelector(weight float64) PlateSet_t{

	plates := make(PlateSet_t)
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
func PlatesSum(plates PlateSet_t) float64{

	sum := 0.0

	for key, value := range plates{
		sum += key * float64(value)
	}

	return sum
}
