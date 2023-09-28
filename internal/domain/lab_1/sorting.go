package lab_1

import (
	"fmt"
	"github.com/Nav1Cr0ss/algorithms/pkg/utilz"
	"time"
)

type SortProvider interface {
	Sort(inputFile, outputFile string) error
}

func SortArrayOfInt(inputFile, outputFile string, sp SortProvider) {
	start := time.Now()

	err := sp.Sort(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error on sorting file")
	}

	utilz.MeasureTime(start, "Sorting array")
}
