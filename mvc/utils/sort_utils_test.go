package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}
	BubbleSort(els)

	assert.NotNil(t, els, "els should not be nil")
	assert.EqualValues(t, 5, len(els), "els should have 5 elements")
	assert.EqualValues(t, 5, els[0], "els[0] should be 5")
	assert.EqualValues(t, 6, els[1], "els[1] should be 6")
	assert.EqualValues(t, 7, els[2], "els[2] should be 7")
	assert.EqualValues(t, 8, els[3], "els[3] should be 8")
	assert.EqualValues(t, 9, els[4], "els[4] should be 9")
}

func getElements(n int) []int{
	result := make([]int, n)
	i := 0
	for j:= n-1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.NotNil(t, els, "elements should be returned")
	assert.EqualValues(t, 5, len(els), "elements should have 5 elements")
}

func BenchmarkBubbleSort(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}