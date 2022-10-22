package main

func mergeSort2(dataset []int) []int {
	if len(dataset) <= 1 {
		return dataset
	}

	mid := len(dataset) / 2
	leftarr := dataset[:mid]
	rightarr := dataset[mid:]
	merged := make([]int, len(dataset))
	// recursively break down the arrays
	leftarrMrg := mergeSort2(leftarr)
	rightarrMrg := mergeSort2(rightarr)

	// perform the merging

	i := 0 // index into the left array
	j := 0 // index into the right array
	k := 0 // index into the merged array

	// while both arrays have content

	for i < len(leftarrMrg) && j < len(rightarrMrg) {
		if leftarrMrg[i] < rightarrMrg[j] {
			merged[k] = leftarrMrg[i]
			i++
			k++
		} else {
			merged[k] = rightarrMrg[j]
			j += 1
			k += 1
		}
	}

	// if the left array still has values, add them

	for i < len(leftarrMrg) {
		merged[k] = leftarrMrg[i]
		i += 1
		k += 1
	}

	// if the right array still has values, add them

	for j < len(rightarrMrg) {
		merged[k] = rightarrMrg[j]
		j += 1
		k += 1
	}

	return merged
}
