package rot13Breaker

import (
	"bytes"
	"math"
	"strings"
)

type attempt struct {
	entropy float64
	shift   int
	preview string
}

var (
	// a-z frequency in the english language. Numbers from: https://www3.nd.edu/~busiforc/handouts/cryptography/Letter%20Frequencies.html
	letterFrequency = [26]float64{0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, 0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, 0.00978, 0.02360, 0.00150, 0.01974, 0.00074}
)

// Decrypt returns a struct slice containing its top n guesses
func Decrypt(message string) [26]attempt {

	attemps := [26]attempt{}

	message = strings.ToLower(message)

	// Brute force approach. All 26 combinations will be tried
	for i := 0; i < 26; i++ {
		attemps[i] = calculateScore(i, &message)
	}

	insertionSort(&attemps)

	// For debugging
	/*
		for _, v := range attemps {
			fmt.Printf("%v %d %s\n", v.entropy, v.shift, v.preview)
		}
	*/

	// Returns the top n guesses
	return attemps
}

func calculateScore(shift int, input *string) attempt {

	var buffer bytes.Buffer

	for _, e := range *input {

		// This will persevere symbols such as spaces
		char := e

		if e >= 'a' && e <= 'z' {
			char = rune(((int(e) - 97 + shift) % 26) + 97)
		}

		buffer.WriteRune(char)
	}
	// With the input now shifted by n so can we analyse it using letter frequency data

	// Offers a preview for
	preview := buffer.String()[:int(math.Min(float64(len(*input)-1), 20))]

	sum := float64(0)
	skipped := 0

	for _, e := range buffer.String() {
		if e >= 'a' && e <= 'z' {
			sum -= math.Log(letterFrequency[e-97])
		} else {
			skipped += 1
		}
	}

	// Calculate the entropy. The lower the better
	entropy := sum / math.Log(2) / float64((len(*input) - skipped))

	return attempt{shift: shift, entropy: entropy, preview: preview}
}

func insertionSort(slice *[26]attempt) {

	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 {
			if slice[j-1].entropy > slice[j].entropy {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j = j - 1
		}
	}
}
