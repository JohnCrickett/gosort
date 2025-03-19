package main

func quickSort(lines []string, low, high int) {
	if low < high {
		p := partition(lines, low, high)
		quickSort(lines, low, p-1)
		quickSort(lines, p+1, high)
	}
}

func partition(lines []string, low, high int) int {
	pivot := lines[high]
	i := low

	for j := low; j < high; j++ {
		if lines[j] < pivot {
			lines[i], lines[j] = lines[j], lines[i]
			i += 1
		}
	}
	lines[high], lines[i] = lines[i], lines[high]
	return i
}
