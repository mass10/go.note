package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// シードを初期化
	rand.Seed(time.Now().UnixNano())

LOOP:

	for ;; {

		for ii := range []int{0, 1, 2, 3, 4} {

			fmt.Printf("[TRACE] LOOP: (%v)\n", ii)

			value := rand.Intn(256)

			if 250 <= value {
				fmt.Printf("[TRACE] FOUND: (%v)\n", value)
				break LOOP
			}
		}
	}
}
