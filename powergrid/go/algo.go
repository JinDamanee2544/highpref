package main

import "math"

func solvePowerGrid(numCities int, adjList [][]int) []int {
	solution := make([]int, numCities)
	bestSolution := make([]int, numCities)
	minCost := math.MaxInt32
	currentCost := 0
	currentDepth := 0

	dfs(&adjList, &solution, &bestSolution, &minCost, currentCost, currentDepth, 0)

	return bestSolution
}

func dfs(adjList *[][]int, solution *[]int, bestSolution *[]int, minCost *int, currentCost int, currentDepth int, currentCity int) {
	if currentCost >= *minCost {
		return
	}

	if currentDepth == len(*solution) {
		*bestSolution = make([]int, len(*solution))
		copy(*bestSolution, *solution)
		*minCost = currentCost
		return
	}

	for i := 0; i < 2; i++ {
		(*solution)[currentCity] = i
		if i == 1 {
			currentCost++
			for _, connectedCity := range (*adjList)[currentCity] {
				(*solution)[connectedCity] = i
			}
		}

		dfs(adjList, solution, bestSolution, minCost, currentCost, currentDepth+1, currentCity+1)

		if i == 1 {
			currentCost--
			for _, connectedCity := range (*adjList)[currentCity] {
				(*solution)[connectedCity] = 0
			}
		}
		(*solution)[currentCity] = 0
	}
}
