package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func WriteMono(data []KeyPair) {
	file, err := os.Create(info.OutputFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bw := bufio.NewWriter(file)

	for _, kv := range data {
		bw.Write([]byte(fmt.Sprintf("std-%v: %v\n", kv.Key, kv.Value)))
	}

	err = bw.Flush()
	if err != nil {
		panic(err)
	}
}

func cuttingData(data []KeyPair) [][]KeyPair {
	dividedData := make([][]KeyPair, info.CPUCore)
	chuckSize := len(data) / int(info.CPUCore)
	for i := 0; i < int(info.CPUCore); i += chuckSize {
		end := i + chuckSize
		if end > len(data) {
			end = len(data)
		}
		dividedData = append(dividedData, data[i:end])
	}
	return dividedData
}

func WriteParallel(data []KeyPair) {
	file, err := os.Create(info.OutputFileName)
	if err != nil {
		panic(err)
	}
	file.Close()

	dividedData := cuttingData(data)

	wgPrintBuff := sync.WaitGroup{}
	wgWrite := sync.WaitGroup{}

	blockSpace := make([]int, info.CPUCore)

	for i := 0; i < int(info.CPUCore); i++ {
		wgPrintBuff.Add(1)
		go func(i int) {
			for _, kv := range dividedData[i] {
				blockSpace[i] += int(len([]byte(fmt.Sprintf("std-%v: %v\n", kv.Key, kv.Value))))
			}
			wgPrintBuff.Done()
		}(i)
	}

	wgPrintBuff.Wait()

	var offset int
	for i := 0; i < int(info.CPUCore); i++ {
		wgWrite.Add(1)
		go func(offset int, data []KeyPair) {
			writeOffset(offset, data)
			wgWrite.Done()
		}(offset, dividedData[i])
		offset += blockSpace[i]
	}

	wgWrite.Wait()
}

func writeOffset(offset int, data []KeyPair) {
	file, err := os.OpenFile(info.OutputFileName, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Seek(int64(offset), 0)

	writer := bufio.NewWriter(file)

	for _, kv := range data {
		writer.WriteString(fmt.Sprintf("std-%v: %v\n", kv.Key, kv.Value))
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
