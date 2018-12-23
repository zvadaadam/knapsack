package algorithms

import (
	"fmt"
	"math/rand"
	"sort"
)

func Evolution(capacity int, items []Item, numGenerations int, populationSize int, tournamentSize int, max int) (int, []bool) {

	populations := initPopulation(len(items), populationSize)

	for i := 0; i < numGenerations; i++ {

		individualA := selection(populations, items, capacity, tournamentSize)
		individualB := selection(populations, items, capacity, tournamentSize)

		child := crossover(individualA, individualB)
		child = mutation(child)

		populations = removeWeakest(populations, items, capacity)

		populations = append(populations, child)

		//fmt.Printf("Generation: %v, Fitness: %v, GlobalMax: %v\n", i , fitness(child, items, capacity), max)
	}

	populations = sortByFitness(populations, items, capacity)
	fmt.Printf("Fitness: %v, GlobalMax: %v\n", fitness(populations[0], items, capacity), max)

	return fitness(populations[0], items, capacity), populations[0]
}

func removeWeakest(population [][]bool, items []Item, capacity int) [][]bool {

	population = sortByFitness(population, items, capacity)

	population = population[:len(population)-1]

	return population
}

func initPopulation(numItems int, populationSize int) [][]bool {


	populations := make([][]bool, populationSize)
	for i := 0; i < populationSize; i++ {
		populations[i] = make([]bool, numItems)

		for j := 0; j < numItems; j++ {
			populations[i][j] = rand.Intn(2) == 0
		}
	}

	return populations
}

func fitness(individual []bool, items []Item, capacity int) int {

	sumValue := 0
	sumWeight := 0
	for i := 0; i < len(individual); i++ {
		if individual[i] {
			sumValue += items[i].Value
			sumWeight += items[i].Weight
		}
	}

	if capacity <= sumWeight {
		sumValue = 0
	}

	return sumValue
}

func sortByFitness(populations [][]bool, items []Item, capacity int) [][]bool {

	sort.Slice(populations, func(left int, right int) bool {
		leftIndividualFitness := fitness(populations[left], items, capacity)
		rightIndividualFitness := fitness(populations[right], items, capacity)

		return leftIndividualFitness > rightIndividualFitness
	})

	return populations
}

func selection(populations [][]bool, items []Item, capacity int, tournamentSize int) []bool {


	tournament := make([][]bool, 0)
	for i := 0; i < tournamentSize; i++ {
		rndIndex := rand.Intn(len(populations))
		tournament = append(tournament, populations[rndIndex])
	}

	tournament = sortByFitness(tournament, items, capacity)

	return tournament[0]
}

func crossover(parentA []bool, parentB []bool) []bool {

	pivot := rand.Intn(len(parentA))

	child := make([]bool, len(parentA))
	for i := 0; i < len(parentA); i++ {
		if i < pivot {
			child[i] = parentA[i]
		} else {
			child[i] = parentB[i]
		}
	}

	return child
}

func mutation(individual []bool) []bool {

	if rand.Intn(10) > 7 {
		rndIndex := rand.Intn(len(individual))
		individual[rndIndex] = !individual[rndIndex]

		rndIndex = rand.Intn(len(individual))
		individual[rndIndex] = !individual[rndIndex]
	}

	return individual
}