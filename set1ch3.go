package main

import (
	"github.com/Melvillian/cryptopals/set1"
	"fmt"
)

func main() {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	probableHexByte := set1.TryAllHexCharDecryptions(hexString)
	fmt.Println(probableHexByte)
}

