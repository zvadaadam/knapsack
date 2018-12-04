package algorithms


/*
*	Bruteforce method for calculating Knapsack problem
*/
func KnapsackBruteForce(capacity int, items []Item, indexes []int, lastIndex int, sumWeight int, sumValue int) (int, int, []int) {

	var bestWeight = sumWeight
	var bestValue = sumValue
	var bestConfiguration []int = indexes

	if lastIndex == len(items) { return sumWeight, sumValue, indexes }

	for i := lastIndex; i < len(items); i++ {
		weight, value, configuration := KnapsackBruteForce(capacity, items, append(indexes, i), i + 1, sumWeight + items[i].Weight, sumValue + items[i].Value)

		if value > bestValue && weight <= capacity{
			bestValue = value
			bestConfiguration = configuration
		}
	}

	return bestWeight ,bestValue, bestConfiguration
}

