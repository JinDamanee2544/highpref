package main

import "fmt"

func main() {
	err := StartUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Start Reading ...")
	dataPair := ReadParallel()

	fmt.Println("Start Sorting ...")
	NativeSort(dataPair)

	fmt.Println("Start Writing ...")
	WriteParallel(dataPair)

	fmt.Println("Done")
}
