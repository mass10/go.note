package application

import (
	"log"
	"spawn-example/child"
	"spawn-example/configuration"
)

// アプリケーションインターフェイス
type Application interface {
	// アプリケーションを実行
	Run()
}

// アプリケーション構造体
type applicationCore struct {
}

func New() Application {
	return &applicationCore{}
}

// アプリケーションを実行
func (self *applicationCore) Run() {
	// コンフィギュレーション
	conf := configuration.Configure()
	if conf.IsChildMode() {
		log.Println("[DEBUG] (applicationCore.Run) アプリケーションをクライアントモードとして実行中...")
		// クライアントモードとして振る舞う
		child := child.New()
		child.Run()
		return
	}

	// 親としての通常の処理
	log.Println("[DEBUG] (applicationCore.Run) アプリケーションを実行中...")
	self.runForParentMain()
}
