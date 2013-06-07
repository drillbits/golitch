package main

import (
	"./golitch"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var (
		inFilename  string
		outFilename string
		num         int
	)
	flag.StringVar(&inFilename, "f", "", "Input filename.")
	flag.StringVar(&outFilename, "o", "glitch.jpg", "Output filename.")
	flag.IntVar(&num, "n", 20, "Number of glitch.")
	flag.Parse()

	f, err := os.Open(inFilename)
	if err != nil {
		fmt.Println(f, err)
		return
	}
	defer f.Close()

	b, err := golitch.Glitch(f, num)
	if err != nil {
		fmt.Println(b, err)
		return
	}

	err = ioutil.WriteFile(outFilename, b, 0644)
	if err != nil {
		fmt.Println(b, err)
		return
	}
}
