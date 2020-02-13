package main

import (
	"fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func main() {
    //ディレクトリを指定する
    fs := http.FileServer(http.Dir("templates"))
    //ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
    http.Handle("/", fs)

    log.Println("Listening...")
    // 3000ポートでサーバーを立ち上げる
    http.ListenAndServe(":3000", nil)
}
