package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"strconv"
	"net/http"
)

// HTTP 静的ファイルサーバーを起動します。
//
// @param location 位置
// @param port ポート番号
func serve(location string, port int) {
	
	fileInfo, err := os.Stat(location)
	if err != nil {
		log.Fatal(err)
	}
	if !fileInfo.IsDir() {
		log.Fatal(location, " is not a directory")
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(location))))
	portText := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(portText, nil))
}

func main() {

	location := ""
	port := 0

	currentOption := "";

	for _, e := range os.Args {
		if currentOption == "--location" {
			location = e
			currentOption = ""
		} else if currentOption == "--port" {
			port, _ = strconv.Atoi(e)
			currentOption = ""
		} else if strings.HasPrefix(e, "--") {
			currentOption = e
		}
	}

	if location == "" {
		log.Fatal("Location is required")
		return
	}

	if port == 0 {
		port = 80
	}

	serve(location, port)
}
