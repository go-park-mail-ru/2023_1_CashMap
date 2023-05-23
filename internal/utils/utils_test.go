package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_SlicesShift(t *testing.T) {
	tests := []struct {
		inputSlice    []int
		expectedSlice []int
		shift         int
		startIndex    int
	}{
		{
			shift:         2,
			startIndex:    0,
			inputSlice:    []int{0, 1, 2, 3, 4, 5},
			expectedSlice: []int{0, 0, 0, 1, 2, 3},
		},

		{
			shift:         1000,
			startIndex:    0,
			inputSlice:    []int{0, 1, 2, 3, 4, 5},
			expectedSlice: []int{0, 0, 0, 0, 0, 0},
		},

		{
			shift:         2,
			startIndex:    100,
			inputSlice:    []int{0, 1, 2, 3, 4, 5},
			expectedSlice: []int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run("slice shift test", func(t *testing.T) {
			test := test
			t.Run("slice shift test", func(t *testing.T) {
				ShiftRight(test.inputSlice, test.startIndex, test.shift)

				assert.Equal(t, test.inputSlice, test.expectedSlice)
			})
		})
	}
}

func TestUtils_SlicesMax(t *testing.T) {
	tests := []struct {
		inputSlice []int
		expected   int
	}{
		{
			inputSlice: []int{0, 1, 2, 3, 4, 5},
			expected:   5,
		},
	}

	for _, test := range tests {
		t.Run("slice max test", func(t *testing.T) {
			test := test
			t.Run("slice max test", func(t *testing.T) {
				max := SliceMax(test.inputSlice)

				assert.Equal(t, test.expected, max)
			})
		})
	}
}

func TestUtils_SlicesMin(t *testing.T) {
	tests := []struct {
		inputSlice []int
		expected   int
	}{
		{
			inputSlice: []int{0, 1, 2, 3, 4, 5},
			expected:   0,
		},
	}

	for _, test := range tests {
		t.Run("slice min test", func(t *testing.T) {
			test := test
			t.Run("slice min test", func(t *testing.T) {
				min := SliceMin(test.inputSlice)

				assert.Equal(t, test.expected, min)
			})
		})
	}
}

func TestUtils_Reverse(t *testing.T) {
	tests := []struct {
		inputSlice    []int
		expectedSlice []int
	}{
		{
			inputSlice:    []int{0, 1, 2, 3, 4, 5},
			expectedSlice: []int{5, 4, 3, 2, 1, 0},
		},
	}

	for _, test := range tests {
		t.Run("slice reverse test", func(t *testing.T) {
			test := test
			t.Run("slice reverse test", func(t *testing.T) {
				Reverse(test.inputSlice)

				assert.Equal(t, test.expectedSlice, test.inputSlice)
			})
		})
	}
}

func TestUtils_GetMaxLength(t *testing.T) {
	tests := []struct {
		inputSlices []string
		expected    int
	}{
		{
			inputSlices: []string{
				"qwerty",
				"qwertyui",
				"qweew",
				"qwas",
				"qwqwwewqeqw",
				"q",
			},

			expected: 11,
		},
	}

	for _, test := range tests {
		t.Run("slice get max length test", func(t *testing.T) {
			test := test
			t.Run("slice get max length test", func(t *testing.T) {
				max := GetMaxLength(test.inputSlices...)

				assert.Equal(t, test.expected, max)
			})
		})
	}
}
