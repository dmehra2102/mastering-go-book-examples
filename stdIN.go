package main

import (
	"bufio"
	"fmt"
	"os"
)

func UsingSTDInput() {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

}
