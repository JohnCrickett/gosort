package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
)

func main() {
	unique := flag.Bool("u", false, "Unique keys.")
	qsort := flag.Bool("qsort", false, "use quick sort")
	msort := flag.Bool("mergesort", false, "use merge sort")
	rSort := flag.Bool("random-sort", false, "use random sort")
	flag.Parse()
	filename := flag.CommandLine.Args()[0]

	sorts := []*bool{qsort, msort, rSort}
	var numSorts int
	for _, sort := range sorts {
		if *sort {
			numSorts++
		}
	}

	if numSorts > 1 {
		fmt.Println("You cannot use multiple sort algorithms together.")
		os.Exit(1)
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if *qsort {
		if *unique {
			fmt.Println("Cannot use quicksort with unique")
			os.Exit(1)
		}
		quickSort(lines, 0, len(lines)-1)
	} else if *msort {
		lines = mergeSort(lines)
	} else if *rSort {
		randomSort(lines)
	} else {
		sort.Strings(lines)
		if *unique {
			lines = slices.Compact(lines)
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
