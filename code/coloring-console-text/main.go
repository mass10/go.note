//
// コンソールの文字に色を付けるサンプルです。
//

package main

import (
	"fmt"
)

func main() {

	// TODO: Windows のコマンドプロンプトだと色が出ない。

	fmt.Println("\x1b[31mRED\x1b[0m")
	fmt.Println("\x1b[32mGREEN\x1b[0m")
}
