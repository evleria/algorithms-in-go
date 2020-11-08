package sorting

import (
	is "github.com/evleria/algorithms-in-go/utils"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSort(t *testing.T) {
	testCases := [][]int{
		{1, 3, 10, 8, 6, -3},
		{6, 1, 40, 5, -10, 0},
		{10, 3, 10, 10, 2, -1},
	}

	sortings := []func([]int){
		bubbleSort,
		insertionSort,
		selectionSort,
		mergeSort,
	}

	for _, sorting := range sortings {
		for _, testCaseTemplate := range testCases {
			testCase := make([]int, len(testCaseTemplate))
			copy(testCase, testCaseTemplate)

			sorting(testCase)

			assert.Check(t, is.Sorted(testCase), testCaseTemplate)
		}
	}
}
