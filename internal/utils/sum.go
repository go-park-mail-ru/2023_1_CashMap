package utils

func Sum[T Numbers](numbers []T) T {
	var out T
	for _, num := range numbers {
		out += num
	}

	return out
}

func Average[T Numbers](numbers []T) float64 {
	return float64(Sum(numbers)) / float64(len(numbers))
}
