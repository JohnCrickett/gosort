package main

func mergeSort(lines []string) []string {
	if len(lines) < 2 {
		return lines
	}

	left := mergeSort(lines[:len(lines)/2])
	right := mergeSort(lines[len(lines)/2:])
	return merge(left, right)
}

func merge(left, right []string) []string {
	var merged = []string{}
	var i, j int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		merged = append(merged, left[i])
	}
	for ; j < len(right); j++ {
		merged = append(merged, right[j])
	}
	return merged

}
