package common

func First[T, U any](val T, _ U) T {
	return val
}
