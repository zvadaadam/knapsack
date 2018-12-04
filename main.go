package main

import (
	"fmt"
	"time"
	"knapsack/algorithms"
)


func readInstance() (int, int, []algorithms.Item, error) {
	var (
		id int = 0
		numItems int = 0
		capacity int = 0
		items []algorithms.Item
	)

	if _, err := fmt.Scan(&id, &numItems, &capacity); err != nil {
		return id ,0, nil, err
	}

	//fmt.Println("Reading input for instance ", id)

	for i := 0; i < numItems; i++  {
		var (
			weight int
			value int
		)
		if _, err := fmt.Scan(&weight, &value); err != nil {
			return id, 0, nil, err
		}
		item := algorithms.Item{Weight:weight, Value:value, Coef:float64(value)/float64(weight)}
		items = append(items, item)
	}

	return id ,capacity, items, nil
}

func printOutput(id int, numItems int, maxPrice int, config []int) {
	//fmt.Printf("%v", config)

	fmt.Printf("%v %v %v  ", id, numItems, maxPrice)

	for i := 0; i < numItems; i++ {
		isPicked := false
		for j := 0; j < len(config); j++ {
			if config[j] == i {
				isPicked = true
			}
		}

		if isPicked {
			fmt.Printf("%v ", 1)
		} else {
			fmt.Printf("%v ", 0)
		}
	}

	fmt.Println()
}

func relativeError(optimal int, apx int) float64 {
	return float64((optimal - apx))/float64(optimal)
}

func mesureInstance() {

	var instanceTimeBF float64 = 0
	var instanceTimeWeight float64 = 0
	var instanceTimeValue float64 = 0
	var instanceTimeCoef float64 = 0

	var (
		sumErrorWeight float64 = 0
		sumErrorValue float64 = 0
		sumErrorCoef float64 = 0
	)

	for {
		_, capacity, items, error := readInstance()
		if error != nil {
			fmt.Printf("%v\n", instanceTimeBF)
			fmt.Printf("%v\n", (instanceTimeWeight + instanceTimeValue + instanceTimeCoef)/3)
			break
		}

		start := time.Now()
		_, valueBF, _ := algorithms.KnapsackBruteForce(capacity, items, []int{}, 0, 0, 0)
		elapsed := time.Since(start).Seconds()
		instanceTimeBF += elapsed

		start = time.Now()
		valueWeight, _ := algorithms.KnapsackHeuristic(capacity, items, algorithms.WeightSorter(items))
		elapsed = time.Since(start).Seconds()
		instanceTimeWeight += elapsed

		start = time.Now()
		valueValue, _ := algorithms.KnapsackHeuristic(capacity, items, algorithms.ValueSorter(items))
		elapsed = time.Since(start).Seconds()
		instanceTimeValue += elapsed

		start = time.Now()
		valueCoef, _ := algorithms.KnapsackHeuristic(capacity, items, algorithms.CoefSorter(items))
		elapsed = time.Since(start).Seconds()
		instanceTimeCoef += elapsed

		sumErrorWeight += relativeError(valueBF, valueWeight)
		sumErrorValue += relativeError(valueBF, valueValue)
		sumErrorCoef += relativeError(valueBF, valueCoef)

		//fmt.Printf("BF %v , W %v, P %v, C %v\n", valueBF, valueWeight, valueValue, valueCoef)
		//fmt.Printf("E Weight: %v\n", errorWeight)
		//fmt.Printf("E Value: %v\n", errorValue)
		//fmt.Printf("E Coef: %v\n", errorCoef)
	}

	fmt.Printf("%v\n", sumErrorWeight/50)
	fmt.Printf("%v\n", sumErrorValue/50)
	fmt.Printf("%v\n", sumErrorCoef/50)

}

func mesureFPTASAcc() {

	accuracy := []float32{0.001, 0.01, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	accuracyError := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for {
		_, capacity, items, error := readInstance()
		if error != nil {
			fmt.Printf("\n")
			for i := 0; i < len(accuracy); i++ {
				fmt.Printf("%v\n", accuracyError[i]/50)
			}
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

		for i := 0; i < len(accuracy); i++ {
			valueBB, _, _ := algorithms.KnapsackBranchAndBound(capacity, items, 0, 0, 0, make([]int, 0), precalSumValue)
			valueFPTAS, _ := algorithms.KnapsackFPTAS(capacity, items, accuracy[i])

			accuracyError[i] += relativeError(valueBB, valueFPTAS)
		}

	}
}



func main() {

	//mesureInstance()

	mesureFPTASAcc()


	var instanceTimeBB float64 = 0
	var instanceTimeDP float64 = 0
	var instanceTimeDW float64 = 0

	//var sumRelativeErrors float64 = 0

	for {
		_, capacity, items, error := readInstance()
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
