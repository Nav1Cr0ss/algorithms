package sort

import "github.com/Nav1Cr0ss/algorithms/pkg/fs"

type MemorySort struct {
	fs *fs.FS
}

func NewMemorySort() *MemorySort {
	return &MemorySort{
		fs: fs.NewFS(),
	}
}

func (s *MemorySort) merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

func (s *MemorySort) MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := s.MergeSort(arr[:mid])
	right := s.MergeSort(arr[mid:])

	return s.merge(left, right)

}

func (s *MemorySort) Sort(inputFile, outputFile string) error {

	nums, err := s.fs.ParseIntArrayFromFile(inputFile)
	if err != nil {
		return err
	}

	sortedNums := s.MergeSort(nums)
	err = s.fs.WriteIntArrayToFile(outputFile, sortedNums)
	if err != nil {
		return err
	}

	return nil
}
