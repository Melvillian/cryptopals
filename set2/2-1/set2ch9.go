package main

import (
	"github.com/Melvillian/cryptopals/lib"
	"fmt"
)

// Implement PKCS#7 padding
func main() {
	data := []byte("YELLOW SUBMARINE")
	paddedData := lib.PadWithPKCS7(data, 20)
	fmt.Println(paddedData)
}
