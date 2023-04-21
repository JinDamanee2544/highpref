package main

import (
	"bufio"
	"os"
)

func Write(binString string) {
	outputFile, err := os.Create(config.outputFileName)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	bw := bufio.NewWriter(outputFile)
	bw.WriteString(binString)

	err = bw.Flush()
	if err != nil {
		panic(err)
	}
}
