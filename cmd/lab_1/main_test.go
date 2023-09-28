package main

import (
	"bufio"
	"fmt"
	"github.com/Nav1Cr0ss/algorithms/internal/domain/lab_1"
	"github.com/Nav1Cr0ss/algorithms/pkg/fs"
	"github.com/Nav1Cr0ss/algorithms/pkg/sort"
	"os"
	"strconv"
	"testing"
)

func TestSorting(t *testing.T) {
	tests := []struct {
		name           string
		sortProvider   lab_1.SortProvider
		inputFileName  string
		outputFileName string
		totalNum       int
	}{
		{
			name:           "MemorySort",
			sortProvider:   sort.NewMemorySort(),
			inputFileName:  "test_memory_input.txt",
			outputFileName: "test_memory_output.txt",
			totalNum:       1000000,
		},
		{
			name:           "ExternalSort",
			sortProvider:   sort.NewExternalSort(),
			inputFileName:  "test_external_input.txt",
			outputFileName: "test_external_output.txt",
			totalNum:       10000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fileSystem := fs.NewFS()
			if err := generateInputFile(fileSystem, tt.inputFileName, tt.totalNum); err != nil {
				panic(err)
			}
			defer func() {
				if err := deleteOutputFile(fileSystem, []string{tt.inputFileName, tt.outputFileName}); err != nil {
					panic(err)
				}
			}()

			lab_1.SortArrayOfInt(tt.inputFileName, tt.outputFileName, tt.sortProvider)

			outputFile, err := os.Open(tt.outputFileName)
			if err != nil {
				t.Fatalf("Failed to open %s: %v", tt.outputFileName, err)
			}
			defer func(outputFile *os.File) {
				err := outputFile.Close()
				if err != nil {
					fmt.Println("Error on closing file:", err)
				}
			}(outputFile)

			scanner := bufio.NewScanner(outputFile)

			numbers := make([]int, 100)
			for i := 0; i < 100; i++ {
				if !scanner.Scan() {
					t.Fatalf("Failed to read 100 numbers from %s", tt.outputFileName)
				}

				num, err := strconv.Atoi(scanner.Text())
				if err != nil {
					t.Fatalf("Failed to parse number from %s: %v", tt.outputFileName, err)
				}

				numbers[i] = num
			}

			if !isSorted(numbers) {
				t.Fatalf("The numbers in %s are not sorted", tt.outputFileName)
			}

		})
	}
}

func isSorted(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return false
		}
	}

	return true
}

func generateInputFile(fs *fs.FS, fileName string, totalNum int) error {

	err := fs.CreateFileWithArrOfInt(fileName, totalNum)
	if err != nil {
		return err
	}

	return nil
}

func deleteOutputFile(fs *fs.FS, fileNames []string) error {
	err := fs.DeleteFiles(fileNames)
	if err != nil {
		return err
	}

	return nil
}
