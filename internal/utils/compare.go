package utils

type Numbers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Numbers](first T, second T) T {
	if first >= second {
		return first
	}

	return second
}

func Min[T Numbers](first T, second T) T {
	if first <= second {
		return first
	}

	return second
}

func MinLength(str1 string, str2 string) int {
	if len(str1) <= len(str2) {
		return len(str1)
	}

	return len(str2)
}
