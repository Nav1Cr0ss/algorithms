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

func (s *ExternalSort) saveSortedChunk(nums *[]int, fileNum int) (string, error) {
	chunkFile := fmt.Sprintf("chunk%d.txt", fileNum)
	sortedNums := s.ms.MergeSort(*nums)
	err := s.fs.WriteChunk(chunkFile, sortedNums)
	if err != nil {
		return "", err
	}
	return chunkFile, nil
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

	var chunkFiles []string

	scanner := bufio.NewScanner(file)
	buffer := make([]int, 0, 100)
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
			chunkFile, err := s.saveSortedChunk(&buffer, len(chunkFiles))
			if err != nil {
				return nil, err
			}
			chunkFiles = append(chunkFiles, chunkFile)

			buffer = []int{}
			memUsed = 0
		}
	}

	if len(buffer) > 0 {
		chunkFile, err := s.saveSortedChunk(&buffer, len(chunkFiles))
		if err != nil {
			return nil, err
		}
		chunkFiles = append(chunkFiles, chunkFile)
	}

	return chunkFiles, nil
}
