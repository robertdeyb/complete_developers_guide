package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, f)
	} else {
		fmt.Println("Please input another argument for file")
	}
}
