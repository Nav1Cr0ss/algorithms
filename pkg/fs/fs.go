package fs

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/Nav1Cr0ss/algorithms/pkg/data_types"
	"math/rand"
	"os"
	"strconv"
)

type FS struct{}

func NewFS() *FS {
	return &FS{}
}

func (f *FS) DeleteFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}

func (f *FS) ParseIntArrayFromFile(filename string) ([]int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error on closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	var nums []int

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return nums, err
	}

	return nums, nil
}

func (f *FS) WriteIntArrayToFile(filename string, arr []int) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("Error on closing file:", err)
		}
	}(outputFile)

	writer := bufio.NewWriter(outputFile)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			fmt.Println("Error on flushing file:", err)
		}
	}(writer)

	for _, num := range arr {
		_, err := writer.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}

	return nil
}

func (f *FS) WriteChunk(filename string, data []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error on closing file:", err)
		}
	}(file)

	writer := bufio.NewWriter(file)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			fmt.Println("Error on flushing file:", err)
		}
	}(writer)

	for _, num := range data {
		_, err := writer.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *FS) MergeChunks(chunkFiles []string, outputFile string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
			fmt.Println("Error on closing file:", err)
		}
	}(output)

	pq := make(data_types.PriorityQueue, len(chunkFiles))

	for i, chunkFile := range chunkFiles {
		chunkReader, err := os.Open(chunkFile)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(chunkReader)
		if scanner.Scan() {
			item := &data_types.Item{
				Value:   scanner.Text(),
				Scanner: scanner,
			}
			pq[i] = item
		}

		if scanner.Err() != nil {
			return scanner.Err()
		}
	}

	heap.Init(&pq)
	writer := bufio.NewWriter(output)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			fmt.Println("Error on flushing file:", err)
		}
	}(writer)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*data_types.Item)
		_, err := fmt.Fprintln(writer, item.Value)
		if err != nil {
			return err
		}

		if item.Scanner.Scan() {
			item.Value = item.Scanner.Text()
			heap.Push(&pq, item)
		}
	}

	return nil
}

func (f *FS) DeleteFiles(fileNames []string) error {
	for _, name := range fileNames {
		err := f.DeleteFile(name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *FS) CreateFileWithArrOfInt(filename string, totalNumbers int) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer func(outputFile *os.File) {
		err = outputFile.Close()
		if err != nil {
			panic("panic on closing file")
		}
	}(outputFile)

	writer := bufio.NewWriter(outputFile)
	defer func(writer *bufio.Writer) {
		err = writer.Flush()
		if err != nil {
			panic("panic on flushing file")
		}
	}(writer)

	for i := 0; i < totalNumbers; i++ {
		randomNum := rand.Intn(1000000)
		_, err := writer.WriteString(strconv.Itoa(randomNum) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}

	return nil

}
