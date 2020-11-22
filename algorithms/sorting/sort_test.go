package sorting

import (
	"fmt"
	"github.com/evleria/algorithms-in-go/utils"
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

type Sorting int

const (
	NotSorted Sorting = iota
	Ascending
	Descending
)

func TestSort(t *testing.T) {
	testCases := []struct {
		from   int
		to     int
		len    int
		sorted Sorting
	}{
		{from: 0, to: 0, len: 0, sorted: NotSorted},            // empty
		{from: 0, to: 20, len: 1000, sorted: NotSorted},        // duplicates
		{from: -1000, to: 1000, len: 1000, sorted: NotSorted},  // ordinary
		{from: 0, to: 0, len: 50, sorted: NotSorted},           // all zeroes
		{from: -1000, to: 1000, len: 1000, sorted: Ascending},  // sorted asc
		{from: -1000, to: 1000, len: 1000, sorted: Descending}, // sorted desc
	}

	sortings := []func([]int){
		bubbleSort,
		insertionSort,
		selectionSort,
		mergeSort,
	}

	for _, sorting := range sortings {
		for _, testCase := range testCases {

			data := utils.GenerateInts(testCase.from, testCase.to, testCase.len)

			switch testCase.sorted {
			case Ascending:
				sort.Sort(sort.IntSlice(data))

			case Descending:
				sort.Sort(sort.Reverse(sort.IntSlice(data)))
			}

			sorting(data)

			assert.Check(t, utils.Sorted(data), fmt.Sprintf("%+v", testCase))
		}
	}
}
