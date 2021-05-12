# Rot13 Breaker

This module will attempt to break a message that has been encrypted with a simple letter substitution cipher (such as [ROT13](https://en.wikipedia.org/wiki/ROT13) or [Caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher)). 

Test coverage: **100%**

This code will brute force each solution (26), analyse each result using frequency analysis and then return an ordered array with the guesses where the first element is the best guess.

## Install

```
go get github.com/CarlFlo/rot13Breaker
```

## Usage

Each element in the output array contains data about the guess
- Entropy: Lower the better, the chance that this is the correct answer
- Shift: How many characters the message was shifted by
- Preview: A short preview (up to 20 characters) of the message that can be used for verifying that the shift is correct

```go
input := "Guvf vf n grfg gb frr vs guvf jbexf"

guesses := rot13Breaker.Decrypt(input)

for _, guess := range guesses {
    fmt.Printf("%f - %d - %s\n", guess.entropy, guess.shift, guess.preview)
}
```
