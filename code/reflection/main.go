package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/// Jimi Hendrix 構造体
type JimiHendrix struct {
	Stratocaster string
}

/// Jimy Page 構造体
type JimmyPage struct {
	LesPaul string
}

/// タイミングによって異なる実体を生成します。
func generateRandomStructure() interface{} {
	n := rand.Intn(256)
	if 210 <= n {
		return JimiHendrix{Stratocaster: "Purple Haze!"}
	}
	if 50 <= n {
		return JimmyPage{LesPaul: "Stairway To Heaven"}
	}
	return nil
}

/// 呼び出しフレームを取得します。
func getCallerName() string {

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return ""
	}

	return fn.Name()
}

func main() {
	// シード値を初期化
	rand.Seed(time.Now().UnixNano())

	frame := getCallerName()
	fmt.Printf("[TRACE] [%v] ### begin ###", frame)

	for {

		if 253 <= rand.Intn(256) {
			break
		}

		// 不明な構造体インスタンスを振りわけ
		switch message := generateRandomStructure().(type) { // nil safe
		case nil:
			fmt.Printf("[WARN] Type: [%v], Message: []\n", message)
		case JimiHendrix:
			fmt.Printf("[INFO] Type: [MessageA], Message: [%v]\n", message.Stratocaster)
		case JimmyPage:
			fmt.Printf("[INFO] Type: [MessageB], Description: [%v]\n", message.LesPaul)
		}

		time.Sleep(10 * 1000 * 1000)
	}

	fmt.Printf("[TRACE] [%v] --- end ---", frame)
}
