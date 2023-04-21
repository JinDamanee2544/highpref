package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
)

var info OsInfo

func Init(inputFileName string, outputFileName string, cpuCore uint8) {
	info.InputFileName = inputFileName
	info.OutputFileName = outputFileName
	info.CPUCore = cpuCore
}

func StartUp() error {
	fmt.Println("Args \t\t\t: ", os.Args)
	fmt.Println("Number of arguments \t: ", len(os.Args))

	if len(os.Args) != 3 {
		fmt.Println("Usage: prog <input file> <output file>")
		return errors.New("Usage: prog <input file> <output file>")
	}
	// os.Args[0] is always prog
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	cpu := flag.Int("cpu", runtime.NumCPU(), "")

	flag.Parse()

	if inputFile == "" {
		return errors.New("Input file name is required (-input)")
	}
	if outputFile == "" {
		return errors.New("Output file name is required (-output)")
	}

	Init(inputFile, outputFile, uint8(*cpu))

	fmt.Println("Your logical cores \t: ", *cpu)
	fmt.Println("Input file \t\t: ", inputFile)
	fmt.Println("Output file \t\t: ", outputFile)

	return nil
}
