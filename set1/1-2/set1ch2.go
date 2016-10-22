package main

import (
	"github.com/Melvillian/cryptopals/lib"
)

func main() {
	str := lib.Xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	lib.PrintComparisons("746865206b696420646f6e277420706c6179", str)
}
