package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization
	els := []int{8, 9, 7, 4, 6}

	// Execution
	els = BubbleSort(els)

	// Validation
	assert.NotNil(t, els)
	assert.EqualValues(t, len(els), 5, "Length of elements should be 5")
	assert.EqualValues(t, els, []int{4, 6, 7, 8, 9})
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
