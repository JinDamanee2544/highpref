package main

import (
	"bufio"
	"fmt"
	"os"
)

var config Config

func readConfig() {
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	if inputFileName == "" {
		panic("No input file specified")
	}
	if outputFileName == "" {
		panic("No output file specified")
	}

	config.inputFileName = inputFileName
	config.outputFileName = outputFileName
}

func Read() ([][]int, int) {
	readConfig()

	inputFile, err := os.Open(config.inputFileName)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Read number of nodes and edges from input file
	var numNodes, numEdges int
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &numNodes)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &numEdges)

	fmt.Println("Number of nodes:", numNodes)
	fmt.Println("Number of edges:", numEdges)

	// Build adjacency list from input file
	adjList := make([][]int, numNodes)
	for i := 0; i < numEdges; i++ {
		scanner.Scan()
		var src, dest int
		fmt.Sscan(scanner.Text(), &src, &dest)

		// fmt.Println("Edge:", src, dest)

		adjList[src] = append(adjList[src], dest)
		adjList[dest] = append(adjList[dest], src)

	}
	fmt.Println(adjList)

	return adjList, numNodes
}
