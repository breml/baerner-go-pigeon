package main

import (
	"fmt"
	"os"

	"github.com/breml/baerner-go-pigeon/wordcount"
)

func main() {
	res, err := wordcount.ParseReader("stdin", os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wordCount, _ := res.(int)

	fmt.Printf("%d words found\n", wordCount)
}
