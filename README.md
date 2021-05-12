# Rot13 Breaker

This module will attempt to break a message that has been encrypted with a simple letter substitution cipher (such as [ROT13](https://en.wikipedia.org/wiki/ROT13) or [Caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher)). 

Test coverage: **100%**

This code will brute force each possible solution (26), analyse each result using frequency analysis and then return an array ordered on probability. The most likely result will be the first element.

## Install

```
go get github.com/CarlFlo/rot13Breaker
```

## Usage

Each element in the output array contains data about the guess
- Shift: How many characters the input has to be shifted by to get the *original* message back
- Preview: A short preview (up to 20 characters) of the message that can be used for verifying that the shift is correct

```go
input := "Guvf vf n grfg gb frr vs guvf jbexf"

guesses := rot13Breaker.Decrypt(input)

for _, guess := range guesses {
    fmt.Printf("%d - %s\n", guess.Shift, guess.Preview)
}
```
