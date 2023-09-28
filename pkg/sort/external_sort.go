package sort

import (
	"bufio"
	"fmt"
	"github.com/Nav1Cr0ss/algorithms/pkg/fs"
	"os"
	"strconv"
)

const maxMemory = int64(100 * 1024 * 1024)

type ExternalSort struct {
	fs *fs.FS
	ms *MemorySort
}

func NewExternalSort() *ExternalSort {
	return &ExternalSort{
		fs: fs.NewFS(),
		ms: NewMemorySort(),
	}
}

func (s *ExternalSort) Sort(inputFile, outputFile string) error {
	chunkFiles, err := s.createSortedChunks(inputFile, maxMemory)
	if err != nil {
		return err
	}

	defer func(files []string) {
		err = s.fs.DeleteFiles(files)
		if err != nil {

		}
	}(chunkFiles)

	err = s.fs.MergeChunks(chunkFiles, outputFile)
	if err != nil {
		return err
	}

	return nil
}

func (s *ExternalSort) createSortedChunks(inputFile string, maxMemory int64) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error on closing file:", err)
		}
	}(file)

	chunkFiles := []string{}

	scanner := bufio.NewScanner(file)
	buffer := []int{}
	memUsed := int64(0)

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		buffer = append(buffer, num)
		memUsed += int64(len(line) + 1)

		if memUsed >= maxMemory {
			sortedNums := s.ms.MergeSort(buffer)
			chunkFile := fmt.Sprintf("chunk%d.txt", len(chunkFiles))
			err := s.fs.WriteChunk(chunkFile, sortedNums)
			if err != nil {
				return nil, err
			}
			chunkFiles = append(chunkFiles, chunkFile)

			buffer = []int{}
			memUsed = 0
		}
	}

	if len(buffer) > 0 {
		sortedNums := s.ms.MergeSort(buffer)
		chunkFile := fmt.Sprintf("chunk%d.txt", len(chunkFiles))
		err := s.fs.WriteChunk(chunkFile, sortedNums)
		if err != nil {
			return nil, err
		}
		chunkFiles = append(chunkFiles, chunkFile)
	}

	return chunkFiles, nil
}
