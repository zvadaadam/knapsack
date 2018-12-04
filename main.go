package main

import (
	"flag"
	"fmt"
	"knapsack/helper"
)


func main() {

	algorithmArg := flag.String("algorithm", "bf", "name of the algorithm")

	flag.Parse()

	algorithmName := helper.AlgorithmName(*algorithmArg)

	switch algorithmName {

	case helper.Bruteforce:
		fmt.Println("Performing Bruteforce...")
		helper.MesureBruteForce()

	case helper.BranchAndBound:
		fmt.Println("Performing Branch&Bound...")
		helper.MesureBranchBound()

	case helper.Heuristic:
		fmt.Println("Performing Heuristic...")
		helper.MesureHeuristic()

	case helper.DynamicPrice:
		fmt.Println("Performing Dynamic Price...")
		helper.MesureDynamicPrice()

	case helper.DynamicWeight:
		fmt.Println("Performing Dynamic Weight...")
		helper.MesureDynamicWeight()

	case helper.FPTAS:
		fmt.Println("Performing FPTAS...")
		helper.MesureFPTAS(0.5)

	default:
		fmt.Println("Not supported algorithm name...")
	}

	return
}
