package math

func CartesianProduct[T any](items []T) [][2]T {
	length := len(items)

	if length < 2 {
		return nil
	}

	capacity := length * (length - 1) / 2
	result := make([][2]T, 0, capacity)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			result = append(result, [2]T{items[i], items[j]})
		}
	}

	return result
}
