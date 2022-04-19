package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "あいうえお"

	// パスワード生成
	bc, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println(string(bc))
}
