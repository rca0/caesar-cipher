package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func cipher(t string, d int) string {
	shift, offset := rune(3), rune(26)
	runes := []rune(t)

	for i, c := range runes {
		switch d {
		case -1: // encoding
			if c >= 'a'+shift && c <= 'z' || c >= 'A' && c <= 'Z' {
				c = c - shift
			} else if c >= 'a' && c < 'a'+shift || c >= 'A' && c < 'A'+shift {
				c = c - shift + offset
			}
		case +1: // decoding
			if c >= 'a' && c <= 'z'-shift || c >= 'A' && c <= 'Z'-shift {
				c = c + shift
			} else if c > 'z'-shift && c <= 'z' || c > 'Z'-shift && c <= 'Z' {
				c = c + shift - offset
			}
		}

		runes[i] = c
	}

	return string(runes)
}

func encode(t string) string  { return cipher(t, -1) }
func decoded(t string) string { return cipher(t, +1) }

func main() {
	fmt.Printf("Ceaser cipher")
	if len(os.Args) < 2 {
		log.Fatalln("You must pass the message to encode/decode")
	}
	encoded := encode(strings.Join(os.Args[1:], " "))
	fmt.Println("Encoded: " + encoded)
	decoded := decoded(encoded)
	fmt.Println("Decoded: " + decoded)
}
