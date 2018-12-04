package algorithms

/*
 *	Bruteforce method with using branch and bound approach for calculating Knapsack problem
 */
func KnapsackBranchAndBound(capacity int, items []Item, weightSum int, valueSum int, index int, config []int, precalSumValues []int) (int, int, []int) {

	if index == len(items) || weightSum > capacity {
		return valueSum, weightSum, config
	}

	//fmt.Println(config)
	leftValue, leftWeight, leftConfig := KnapsackBranchAndBound(capacity, items, weightSum + items[index].Weight, valueSum + items[index].Value, index + 1, append(config, 1), precalSumValues)

	var (
		rightValue = valueSum
		rightWeight = weightSum
		rightConfig = config
	)

	if valueSum + precalSumValues[index + 1] >= leftValue {
		rightValue, rightWeight, rightConfig = KnapsackBranchAndBound(capacity, items, weightSum, valueSum, index + 1, append(config, 0), precalSumValues)
	}

	if leftWeight <= capacity && leftValue > rightValue {
		return leftValue, leftWeight, leftConfig
	} else if rightWeight <= capacity {
		return rightValue, rightWeight, rightConfig
	}

	return 0, 0, config
}
