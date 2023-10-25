package main

import (
	"flag"
	"github.com/Nav1Cr0ss/algorithms/pkg/fs"

	//fs2 "github.com/Nav1Cr0ss/algorithms/pkg/fs"
	"github.com/Nav1Cr0ss/algorithms/pkg/utilz"
	"time"
)

func main() {
	var fileName string
	var numsToGenerate int

	flag.StringVar(&fileName, "file_name", "input.txt", "file name")
	flag.IntVar(&numsToGenerate, "output", 100000, "Nums to generate")

	files := fs.NewFS()
	start := time.Now()
	err := files.CreateFileWithArrOfInt(fileName, numsToGenerate)
	if err != nil {
		return
	}
	utilz.MeasureTime(start, "Generating array")
}
