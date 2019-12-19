package main

import (
	"fmt"
	"strings"
)

func main() {
	DecryptText()
}

// DecryptText use for decrypt cipher text by Monoalphabetic Substitution
func DecryptText() {
	alp := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	r := []string{}
	fmt.Print("Enter Cipher text: ")
	var text string
	fmt.Scan(&text)
	fmt.Print("Enter Key: ")
	var key int
	fmt.Scan(&key)
	pt := []rune(text)
	for i := range pt {
		for index, v := range alp {
			if v == string(pt[i]) {
				if index-key <= 0 {
					index = cap(alp) - (key - index)
				} else {
					index = (index % cap(alp)) - key
				}
				r = append(r, strings.Join(alp[index:index+1], " "))
			}
		}
	}
	fmt.Print(r)
}
