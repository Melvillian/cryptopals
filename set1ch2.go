package main

import (
	"github.com/Melvillian/cryptopals/set1"
	"fmt"
)

func main() {
	str := set1.Xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965");
	fmt.Println(str)
}

