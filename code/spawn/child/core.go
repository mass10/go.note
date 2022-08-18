package child

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ChildProcess interface {
	Run()
}

type childProcessCore struct {
}

func New() ChildProcess {
	child := childProcessCore{}
	return &child
}

func (self *childProcessCore) Run() {
	log.Println("[DEBUG] (childProcessCore.Run) start")

	// response := os.Stdout

	// 標準入力にエラーの無い限り
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if !strings.HasPrefix(line, "[MSG]") {
			log.Println("[WARN] 不明な入力は無視されました。", line)
			continue
		}
		line = line[5:]
		if strings.HasPrefix(line, "SHUTDOWN") {
			break
		}
		if strings.HasPrefix(line, "QUIT") {
			break
		}

		// レスポンス
		fmt.Printf("[RPLY]%s\n", line)
	}

	log.Println("[DEBUG] (childProcessCore.Run) end")
}
