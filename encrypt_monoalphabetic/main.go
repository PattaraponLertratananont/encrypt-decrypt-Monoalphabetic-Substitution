package main

import (
	"fmt"
	"strings"
)

func main() {
	EncryptText()
}

//EncryptText use for encrypt by algorithm Monoalphabetic Substitution
func EncryptText() {
	alp := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	r := []string{}
	var text string
	var key int
	fmt.Print("Enter Plain text: ")
	fmt.Scan(&text)
	fmt.Print("Enter Key: ")
	fmt.Scan(&key)
	pt := []rune(text)
	for i := range pt {
		for index, v := range alp {
			if v == string(pt[i]) {
				index = (index + key) % cap(alp)
				r = append(r, strings.Join(alp[index:index+1], " "))
			}
		}
	}
	fmt.Print(r)
}
