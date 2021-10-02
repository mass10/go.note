package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// シードを初期化
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Printf("[%v]\n", rand.Intn(256))
	}
}
