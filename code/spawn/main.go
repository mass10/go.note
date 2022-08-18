package main

import (
	"io"
	"log"
	"os"

	"spawn-example/application"
	"spawn-example/configuration"
)

// エントリーポイント
func main() {

	// ========== ログファイルを準備 ==========
	filepath := "application.log"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	// ========== ロガーをセットアップ ==========
	{
		log.SetOutput(io.MultiWriter(os.Stderr, file))
		log.Default().SetPrefix("[my-logging] ")
		log.Default().SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	}

	log.Println("### start ###")

	for i, e := range os.Args {
		log.Printf("コマンドライン引数 %d: [%v]\n", i, e)
	}

	// ========== アプリケーションを初期化 ==========
	{
		configuration.Configure()
		app := application.New()
		app.Run()
	}

	// ========== SHUTDOWN ==========
	log.Println("--- end ---")
}
