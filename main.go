package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Please provide directory name")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.Size() == 0 {
			fmt.Println(file.Name())
		}

	}
}
