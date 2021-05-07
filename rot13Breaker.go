package main

import (
	"fmt"
	"strings"
)

type character struct {
	char  rune
	count int
}

var (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	// Most frequent characters in the english language
	mfc = [6]rune{'e', 't', 'a', 'o', 'i', 'n'}
)

func main() {
	// THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG
	Decrypt("QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD")
}

//
func Decrypt(message string) {

	message = strings.ToLower(message)

	// Each index represents a corresponding letter in the alphabet i.e. a==0 and z==25.
	var characters [26]character

	// Initialize the slice
	for i := range characters {
		characters[i].char = rune(alphabet[i])
	}

	for _, char := range message {
		if char >= 'a' && char <= 'z' {
			characters[char-97].count++
		}
	}

	// Orders the chars in decending order in the array
	insertionSort(&characters)

	fmt.Printf("%v\n", characters)

	// The message is broken using frequence analysis
	for i, c := range mfc {
		char := characters[0+i].char

		if c == char {

		}

	}

}

func insertionSort(slice *[26]character) {

	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 {
			if slice[j-1].count < slice[j].count {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j = j - 1
		}
	}
}
