package main

import (
	"fmt"
	"knapsack/algorithms"
	"knapsack/helper"
	"time"
)


func main() {

	//mesureInstance()

	helper.MesureFPTASAcc()


	var instanceTimeBB float64 = 0
	var instanceTimeDP float64 = 0
	var instanceTimeDW float64 = 0

	//var sumRelativeErrors float64 = 0

	for {
		_, capacity, items, error := helper.ReadInstance()
		if error != nil {
			fmt.Printf("\nBB: %v\n", instanceTimeBB)
			fmt.Printf("DW: %v\n", instanceTimeDW)
			fmt.Printf("DP: %v\n", instanceTimeDP)
			//fmt.Printf("Relative Errors: %v\n", sumRelativeErrors/50)
			break
		}

		precalSumValue := make([]int, len(items) + 1)
		for i := 0; i < len(items); i++ {
			precalSumValue[i] = 0
			for j := i; j < len(items); j++ {
				precalSumValue[i] += items[j].Value
			}
		}
		precalSumValue[len(items)] = 0

		start := time.Now()
		algorithms.KnapsackBranchAndBound(capacity, items, 0, 0, 0, make([]int, 0), precalSumValue)
		elapsed := time.Since(start).Seconds()
		instanceTimeBB += elapsed

		start = time.Now()
		algorithms.KnapsackDynamicPrice(capacity, items)
		elapsed = time.Since(start).Seconds()
		instanceTimeDP += elapsed

		start = time.Now()
		algorithms.KnapsackDynamicWeight(capacity, items)
		elapsed = time.Since(start).Seconds()
		instanceTimeDW += elapsed

		//valueAPX, _ := knapsackFPTAS(capacity, items, 1)
		//sumRelativeErrors += relativeError(valueOPT, valueAPX)

		//fmt.Printf("%v\n", relativeError(valueOPT, valueAPX))


		//value, config := knapsackDynamicPrice(capacity, items)
		//
		//fmt.Printf("%v %v %v ", id, len(items), value)
		//for i := 0; i < len(items); i++ {
		//	fmt.Printf(" %b", config[i])
		//}
		//fmt.Printf("\n")

	}

	return
}
