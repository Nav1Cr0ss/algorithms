package lab_1

import (
	"fmt"
	"github.com/Nav1Cr0ss/algorithms/pkg/utilz"
	"os"
	"time"
)

type SortProvider interface {
	Sort(inputFile, outputFile string) error
}

func Run(inputFile, outputFile string, sp SortProvider) {
	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File size: %.2f mb\n", float64(fileInfo.Size())/(1024*1024))

	start := time.Now()

	err = sp.Sort(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error on sorting file")
	}

	utilz.MeasureTime(start, "Sorting array")
}
