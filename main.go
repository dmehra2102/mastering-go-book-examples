package main

import (
	"fmt"

	hashtable "example.com/mastering-go/hash-table"
)

func main() {
	var r byte = 'A'                         // ASCII character
	var r2 rune = '世'                        // Chinese character
	var r3 rune = 0x4E16                     // Unicode code point for '世'
	var r4 rune = '😊'                        // Emoji
	fmt.Printf("%c %c %c %c", r, r2, r3, r4) // Output: A 世 世 😊
	// UsingSTDOutput()
	// UsingSTDInput()
	// GoEnvironmentExample()
	// LoopsExample()
	// SliceExamples()
	// CalculatePi()
	// store.UseStore()
	// binarytree.BinaryTreeUsage()
	hashtable.HashTableExample()
}
