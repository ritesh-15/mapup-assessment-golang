package utils

import (
	"sort"
	"sync"
)

func SequentialSort(arrays [][]int) [][]int {
	for index := range arrays {
		sort.Ints(arrays[index])
	}
	return arrays
}

func ConcurrentSort(arrays [][]int) [][]int {
	var wg sync.WaitGroup
	channel := make(chan []int, len(arrays))

	for _, array := range arrays {
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			sort.Ints(nums)
			channel <- nums
		}(array)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	sortedArrays := make([][]int, len(arrays))

	for i := 0; i < len(arrays); i++ {
		sortedArrays[i] = <-channel
	}

	return sortedArrays
}
