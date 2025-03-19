package main

import "math/rand/v2"

func randomSort(lines []string) {
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
}
