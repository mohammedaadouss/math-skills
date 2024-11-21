package main

import (
	"fmt"
	"os"

	dvma "dvma/internal"
)

func main() {
	fileName := os.Args[1]
	if len(os.Args) != 2 {
		fmt.Println("Check your arguments!")
		return
	} else {
		if !dvma.NameChecker(fileName) {
			fmt.Println("Check the name of the file you are entering!")
		} else {
			dvma.DataSplitter(fileName)
		}
	}
}
