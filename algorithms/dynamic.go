package algorithms

import "math"

func max(a int, b int) int {
	if a > b { return a }
	return b
}

func min(a int, b int) int {
	if a < b { return a }
	return b
}

func KnapsackDynamicWeight(capacity int, items []Item) (int, []int) {

	var matrix = make([][]int, capacity + 1)
	for i := 0; i <= capacity; i++ {
		matrix[i] = make([]int, len(items) + 1)
	}

	// rows - weight
	for w := 0; w <= capacity; w++ {

		// column - items
		for i := 0; i <= len(items); i++ {

			if w == 0 || i == 0 {
				matrix[w][i] = 0
			} else if items[i-1].Weight > w {
				matrix[w][i] = matrix[w][i - 1]
			} else {
				left := items[i-1].Value + matrix[w - items[i-1].Weight][i - 1]
				right := matrix[w][i - 1]
				matrix[w][i] = max(left, right)
			}
		}
	}

	var config []int
	config = backtrack(items, capacity, matrix, capacity, len(items), config)

	return matrix[capacity][len(items)], config
}

func KnapsackDynamicPrice(capacity int, items []Item) (int, []int) {

	var sumPrice int = 0
	for i := 0; i < len(items); i++ {
		sumPrice += items[i].Value
	}

	var matrix = make([][]int, sumPrice + 1)
	for i := 0; i <= sumPrice; i++ {
		matrix[i] = make([]int, len(items) + 1)
	}

	for i := 0; i <= sumPrice; i++ {
		matrix[i][0] = math.MaxInt32
	}

	for i := 0; i <= len(items); i++ {
		matrix[0][i] = 0
	}

	// rows - value
	for i := 1; i <= len(items); i++ {

		// column - items
		for v := 1; v <= sumPrice; v++ {

			weightNotIncluded := matrix[v][i - 1]
			weightIncluded := math.MaxInt32

			if (v >= items[i - 1].Value) {
				w := matrix[v - items[i - 1].Value][i - 1]
				if (w != math.MaxInt32) {
					weightIncluded = w + items[i - 1].Weight
				}
			}

			matrix[v][i] = min(weightNotIncluded, weightIncluded);
		}
	}

	var bestValue int = 0
	//var indexPrice int = 0
	for v := 0; v <= sumPrice; v++ {
		//fmt.Printf("[%v %v]", matrix[v][len(items)], v)
		if matrix[v][len(items)] <= capacity && v > bestValue{
			bestValue = v
			//indexPrice = v
		}
	}

	//var config []int
	//config = backtrackPrice(items, capacity, matrix, indexPrice, len(items), config)

	return bestValue, []int{0}


}

func backtrack(items []Item, capacity int, matrix [][]int, indexWeight int, indexItem int, config []int) []int {
	if indexItem == 0 {
		return config
	}

	if items[indexItem - 1].Weight == 0 || matrix[indexWeight][indexItem] != matrix[indexWeight][indexItem - 1] {
		config = backtrack(items, capacity, matrix, indexWeight - items[indexItem - 1].Weight, indexItem - 1, config)
		config = append(config, 1)
	} else {
		config = backtrack(items, capacity, matrix, indexWeight, indexItem - 1, config)
		config = append(config, 0)
	}

	return config
}

func backtrackPrice(items []Item, capacity int, matrix [][]int, indexPrice int, indexItem int, config []int) []int {
	if indexItem == 0 {
		return config
	}

	if matrix[indexPrice][indexItem] != matrix[indexPrice][indexItem - 1] {
		config = backtrack(items, capacity, matrix, indexPrice - items[indexItem - 1].Value, indexItem - 1, config)
		config = append(config, 1)
	} else {
		config = backtrack(items, capacity, matrix, indexPrice, indexItem - 1, config)
		config = append(config, 0)
	}

	return config
}
