package helper

import (
	"fmt"
	"knapsack/algorithms"
)

func ReadInstance() (int, int, []algorithms.Item, error) {
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

func PrintOutput(id int, numItems int, maxPrice int, config []int) {
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
