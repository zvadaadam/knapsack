package algorithms

import "sort"

type ValueSorter []Item
func (a ValueSorter) Len() int { return len(a) }
func (a ValueSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ValueSorter) Less(i, j int) bool { return a[i].Value > a[j].Value }


// Weight is sorted descendence
type WeightSorter []Item
func (a WeightSorter) Len() int { return len(a) }
func (a WeightSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a WeightSorter) Less(i, j int) bool { return a[i].Weight < a[j].Weight}


type CoefSorter []Item
func (a CoefSorter) Len() int { return len(a) }
func (a CoefSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a CoefSorter) Less(i, j int) bool { return a[i].Coef > a[j].Coef }


func KnapsackHeuristic(capacity int, items []Item, sorter sort.Interface) (int, []int) {

	sort.Sort(sorter)

	return fillKnapsack(capacity, items)
}


func fillKnapsack(capacity int, items []Item) (int, []int) {
	var (
		sumWeight = 0
		sumValue = 0
		bestConfig []int
	)

	for i := 0; i < len(items); i++ {
		currentWeight := items[i].Weight
		if (sumWeight + currentWeight) < capacity  {
			bestConfig = append(bestConfig, i)
			sumWeight += items[i].Weight
			sumValue += items[i].Value
		}
	}

	return sumValue, bestConfig
}


