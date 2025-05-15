package utils

func InList[T comparable](list []T, item T) bool {
	for _, t := range list {
		if t == item {
			return true
		}
	}
	return false
}
