package main

import (
	"flag"
	"github.com/Nav1Cr0ss/algorithms/internal/domain/lab_1"
	"github.com/Nav1Cr0ss/algorithms/pkg/sort"
)

func main() {

	var inputFileName string
	var outputFileName string
	var useExternalSort bool

	flag.StringVar(&inputFileName, "input", "input.txt", "Input file name")
	flag.StringVar(&outputFileName, "output", "output.txt", "Output file name")
	flag.BoolVar(&useExternalSort, "external", false, "Use external sort (default is in-memory sort)")

	flag.Parse()

	var sorter lab_1.SortProvider

	switch useExternalSort {
	case true:
		sorter = sort.NewExternalSort()
	default:
		sorter = sort.NewMemorySort()
	}

	lab_1.SortArrayOfInt(inputFileName, outputFileName, sorter)

}
