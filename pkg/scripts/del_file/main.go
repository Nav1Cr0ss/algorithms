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

	flag.StringVar(&fileName, "file_name", "input.txt", "file name")

	files := fs.NewFS()
	start := time.Now()
	err := files.DeleteFile(fileName)
	if err != nil {
		return
	}
	utilz.MeasureTime(start, "Deleting")
}
