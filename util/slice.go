package util

func SliceContains[T comparable](slice []T, item T) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func SliceIndexOf[T comparable](slice []T, item T) int {
	i := 0
	for _, a := range slice {
		if a == item {
			return i
		}
		i++
	}
	return -1
}

func SliceRemove[T comparable](slice []T, item T) []T {
	index := SliceIndexOf(slice, item)
	if index == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func SliceUnique[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func SliceReverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func SliceMin(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	minValue := slice[0]
	for _, v := range slice[1:] {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}

func SliceMax(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	maxValue := slice[0]
	for _, v := range slice[1:] {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func SliceLastItems[T any](slice []T, count int) []T {
	if len(slice) < count {
		return slice
	}
	return slice[len(slice)-count:]
}
