package set

type myMap[K comparable, V any] map[K]V

// DiffArray 求两个切片的差集
func DiffArray[T comparable](a []T, b []T) []T {
	var diffArray []T
	var temp myMap[T, struct{}] = map[T]struct{}{}
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}
	return diffArray
}

// IntersectArray 求两个切片的交集
func IntersectArray[T comparable](a []T, b []T) []T {
	var inter []T
	var mp myMap[T, bool] = map[T]bool{}

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}
	return inter
}
