package utils

type primitives interface {
	~int | ~int32 | ~int16 | ~int8 | ~int64 | ~float32 | ~float64 | ~string | ~byte | ~bool
}

func Includes[T primitives](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}
