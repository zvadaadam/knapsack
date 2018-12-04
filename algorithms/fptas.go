package algorithms


func KnapsackFPTAS(capacity int, items []Item, accuracy float32) (int, []int) {

	var maxValue int = 0
	for i := 0; i < len(items); i++ {
		if items[i].Value > maxValue {
			maxValue = items[i].Value
		}
	}

	var factor float32
	if accuracy != 1 {
		epsilon := 1 - accuracy
		factor = epsilon * (float32(maxValue)/float32(len(items)))
	} else {
		factor = 1
	}

	for i := 0; i < len(items); i++ {
		//items[i].factorValue = int(float32(items[i].value)/factor)
		items[i].Value = int(float32(items[i].Value)/factor)
	}

	value, config := KnapsackDynamicPrice(capacity, items)

	return int(float32(value) * factor), config
}


