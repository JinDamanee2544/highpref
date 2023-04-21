package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var chuckSize int64

func ReadMono() []KeyPair {
	file, err := os.Open(info.InputFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data []KeyPair

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if line[0] == 's' {
			group := strings.Split(line, ": ")
			index := group[0][4:]
			val := group[1][:len(group[1])-1]
			key, err := strconv.ParseFloat(index, 64)
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, KeyPair{float64(key), float64(value)})
		}
	}
	return data
}

func ReadParallel() []KeyPair {
	file, err := os.Open(info.InputFileName)
	if err != nil {
		panic(err)
	}
	fileStat, err := file.Stat()
	file.Close()

	var wg sync.WaitGroup

	dataBlock := make([][]KeyPair, info.CPUCore)
	chuckSize = fileStat.Size() / int64(info.CPUCore)
	wg.Add(int(info.CPUCore))

	var offset int64 = 0
	for i := 0; i < int(info.CPUCore); i++ {
		go func(offset int64, i int) {
			dataBlock[i] = read(offset)
			wg.Done()
		}(offset, i)
		offset += int64(chuckSize)
	}
	wg.Wait()

	var data []KeyPair

	for i := 0; i < int(info.CPUCore); i++ {
		data = append(data, dataBlock[i]...)
	}

	return data
}

func read(offset int64) []KeyPair {
	var data []KeyPair

	file, err := os.Open(info.InputFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Seek(offset, 0); err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)

	var current int64 = 0
	for current < chuckSize {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if line[0] == 's' {
			group := strings.Split(line, ": ")
			index := group[0][4:]
			val := group[1][:len(group[1])-1]
			key, err := strconv.ParseFloat(index, 64)
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, KeyPair{float64(key), float64(value)})
		}
		current += int64(len([]byte(line)))
	}

	return data
}
