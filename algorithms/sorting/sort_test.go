package sorting

import (
	"fmt"
	"github.com/evleria/algorithms-in-go/utils"
	"gotest.tools/v3/assert"
	"sort"
	"testing"
)

type Sorting int

const (
	Ascending Sorting = iota + 1
	Descending
)

func TestSort(t *testing.T) {
	testCases := []struct {
		from   int
		to     int
		len    int
		sorted Sorting
	}{
		{from: 0, to: 0, len: 0},                               // empty
		{from: 0, to: 20, len: 1000},                           // duplicates
		{from: -1000, to: 1000, len: 1000},                     // ordinary
		{from: 0, to: 0, len: 50},                              // all zeroes
		{from: -1000, to: 1000, len: 1000, sorted: Ascending},  // sorted asc
		{from: -1000, to: 1000, len: 1000, sorted: Descending}, // sorted desc
	}

	sortings := []func([]int){
		bubbleSort,
		insertionSort,
		selectionSort,
		mergeSort,
		quickSort,
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
