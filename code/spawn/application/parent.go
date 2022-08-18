package application

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

// 親としての通常の操作
func (self *applicationCore) runForParentMain() {

	command := exec.Command("go.exe", "run", "main.go", "--fork")

	// プロセスの標準入力
	requestIf, err := command.StdinPipe()
	if err != nil {
		log.Fatal(err)
		return
	}
	requestToChild := bufio.NewWriter(requestIf)

	// 子プロセスからの出力を読み込む
	responseIf, err := command.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return
	}

	// 子プロセスを起動
	log.Println("[DEBUG] (applicationCore.runForParentMain) 子プロセスを実行しています...")
	err = command.Start()
	if err != nil {
		log.Println("[ERROR] プログラムを起動できません。", err)
		log.Fatal(err)
		return
	}
	// command.ProcessState

	log.Println("[DEBUG] BEGIN LOOP")

	i := 0

	responseFromChild := bufio.NewScanner(responseIf)

	for {

		if 10 <= i {
			// 適度に終了
			break
		}

		requestString := "[MSG]Hello"

		log.Printf("親 > %s > 子", requestString)

		// >> 子プロセスへの要求 >>
		requestToChild.WriteString("[MSG]Hello\n")
		requestToChild.Flush()

		// 応答を読み取り
		for responseFromChild.Scan() {
			line := responseFromChild.Text()
			if !strings.HasPrefix(line, "[RPLY]") {
				continue
			}
			log.Printf("子 > %s > 親\n", line)
			break
		}

		i = i + 1
	}

	log.Println("[DEBUG] END LOOP")
}
