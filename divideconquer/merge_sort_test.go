package divcon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	A := []int{1, 4, 5, 8}
	B := []int{2, 3, 6, 7}

	C := merge(A, B)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, expected, C)
}

func TestDivide(t *testing.T) {
	A := []int{1, 3, 10, 4, 2, 6, 9, 8}
	result := divide(A)
	assert.Equal(t, 4, len(result))
	t.Log(result)

	B := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result = divide(B)
	assert.Equal(t, 5, len(result))
	t.Log(result)
}

func TestSort(t *testing.T) {
	A := []int{2, 1}
	sortTuple(A)
	expected := []int{1, 2}
	assert.Equal(t, expected, A)
}

func TestMergeSort(t *testing.T) {
	inp := []int{5, 4, 1, 8, 7, 2, 9, 6, 3}
	result := MergeSort(inp)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, result)
	t.Log(result)
}
