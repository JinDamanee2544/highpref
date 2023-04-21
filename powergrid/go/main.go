package main

func main() {
	adjList, numNodes := Read()

	solution := solvePowerGrid(numNodes, adjList)

	binString := ""
	for _, b := range solution {
		if b == 1 {
			binString += "1"
		} else {
			binString += "0"
		}
	}
	Write(binString)
}
