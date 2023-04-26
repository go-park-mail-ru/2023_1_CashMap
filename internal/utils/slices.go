package utils

func ShiftRight[T any](slice []T, startIndex int, shift int) {
	for i := len(slice) - 1 - shift; i >= startIndex; i-- {
		slice[i+shift] = slice[i]
	}

	var lastEmptyIndex int
	if startIndex+shift < len(slice) {
		lastEmptyIndex = startIndex + shift - 1
	} else {
		lastEmptyIndex = len(slice) - 1
	}
	var empty T
	for i := startIndex; i <= lastEmptyIndex; i++ {
		slice[i] = empty
	}
}

func SliceMax[T Numbers | string](slice []T) T {
	var maxValue T
	for _, value := range slice {
		if maxValue <= value {
			maxValue = value
		}
	}

	return maxValue
}

func SliceMin[T Numbers](slice []T) T {
	var minValue T = slice[0]
	for _, value := range slice {
		if minValue >= value {
			minValue = value
		}
	}

	return minValue
}

func GetMaxLength(slice ...string) int {
	var maxValue int
	for _, value := range slice {
		if maxValue < len(value) {
			maxValue = len(value)
		}
	}

	return maxValue
}
