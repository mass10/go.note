package main

import (
	"fmt"
	"regexp"
)

func 数字文字列の検査(s string) {

	b := []byte(s)
	ok, _ := regexp.Match("^[0-9]+$", b)
	if ok {
		fmt.Println("\"" + s + "\" は数字です。")
	} else {
		fmt.Println("\"" + s + "\" は数字ではありません。")
	}
}

func 正規表現のテスト() {

	数字文字列の検査("")
	数字文字列の検査("0")
	数字文字列の検査("00")
	数字文字列の検査(" 00")
	数字文字列の検査("00 ")
	数字文字列の検査("789213-20138123-181723")
	数字文字列の検査("-")
	数字文字列の検査("00000")
	数字文字列の検査("999999999999")
	数字文字列の検査("99999.9999999")
	数字文字列の検査("872137126\n00113")
	数字文字列の検査("098ASy317ASyuncl023kjHSA")
	数字文字列の検査("0A0")
	数字文字列の検査("0A9")
	数字文字列の検査("abc")
}

func diagnoseCommandString(s string) {

	// b := []byte(s)
	// pattern := regexp.MustCompile("^?flags[\"]?[\\:a-zA-Z0-9_ ]+[\"]?$").Split(s, -1)
	// pattern := regexp.MustCompile("^?flags([\\:a-zA-Z0-9_ ]+)$").FindSubmatch([]byte(s))
	// pattern := regexp.MustCompile("^?flags([\\:a-zA-Z0-9_ ]+)$").FindSubmatch([]byte(s))
	pattern := regexp.MustCompile(`(?ms)\A(["]?[\\:a-zA-Z0-9+-\. \(\)]+["]?)[ ]?(["]?[\\:a-zA-Z0-9+-\. \(\)]*["]?)\z`).FindAll([]byte(s), -1)
	// pattern := regexp.MustCompile("^?flags([\"]?[\\:a-zA-Z0-9_ ]+[\"]?)$").FindSubmatch([]byte(s))
	if len(pattern) > 0 {
		fmt.Printf("(%d RESULTS)\n", len(pattern))
		for _, e := range pattern {
			fmt.Printf("[%s]\n", e)
		}
	} else {
		fmt.Printf("(NO RESULTS)\n")
	}
}

func main() {
	if false {
		正規表現のテスト()
	}

	diagnoseCommandString("999")
	diagnoseCommandString("abc")
	diagnoseCommandString("AhjSy7819")
	diagnoseCommandString("Ah jS y7819")
	diagnoseCommandString("C:\\Ah jS y7819")
	diagnoseCommandString("C:\\program files\\techtouc edd")
	diagnoseCommandString("C:\\program files\\techtouch editor (staging)\\")
	diagnoseCommandString("C:\\program files\\techtouch editor (staging)\\editor.exe")
	diagnoseCommandString("\"C:\\program files\\techtouch editor (staging)\\editor.exe\"")
	diagnoseCommandString("\"C:\\program files\\techtouch editor (staging)\\editor.exe\" \"path to bough\"")
	diagnoseCommandString("\"\"C:\\program files\\techtouch editor (staging)\\editor.exe\"")
}
