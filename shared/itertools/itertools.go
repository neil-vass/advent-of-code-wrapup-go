package itertools

// I'm missing Python's itertools.combinations().
// Made this function by following its pseudocde:
// https://docs.python.org/3/library/itertools.html#itertools.combinations
// The original relies on some Python feautres, making this version harder to follow.
func Combinations[T any](pool []T, r int) [][]T {
	combos := [][]T{}
	n := len(pool)
	if r > n {
		return combos
	}

	indices := make([]int, r)
	for i := range r {
		indices[i] = i
	}

	getCombo := func() []T {
		c := make([]T, r)
		for i, poolIdx := range indices {
			c[i] = pool[poolIdx]
		}
		return c
	}

	combos = append(combos, getCombo())
	for {
		i := r - 1
		didBreak := false
		for ; i >= 0; i-- {
			if indices[i] != i+n-r {
				didBreak = true
				break
			}
		}
		if !didBreak {
			return combos
		}

		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}
		combos = append(combos, getCombo())
	}
}

// I'm missing Python's itertools.product().
// Made this function by following its pseudocde:
// https://docs.python.org/3.7/library/itertools.html#itertools.product
func Product[T any](pools ...[]T) [][]T {
	result := [][]T{{}}
	for _, pool := range pools {
		updatedResult := [][]T{}
		for _, x := range result {
			for _, y := range pool {
				updatedResult = append(updatedResult, append(x, y))
			}
		}
		result = updatedResult
	}
	return result
}
