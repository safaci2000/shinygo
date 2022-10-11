package fuzz

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Reverse a string.
func Reverse(s string) (string, error) {
	// b := []byte(s) //Bug
	//Patch
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}

	fmt.Printf("input: %q\n", s)
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

/*
func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}

*/
