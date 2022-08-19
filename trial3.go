package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("text.txt") //ファイルをオープン
	if err != nil {
		fmt.Println("cannot open the file")
	}
	defer f.Close() //エラー処理〜f.Close()はテンプレ

	data := make([]byte, 1024) //1024バイトのバイト配列を作成
	count, err := f.Read(data) //1024バイト分、ファイルを読むという処理
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to read file")
	}

	fmt.Printf("read %d bytes:\n", count) //なんバイト読み込まれたかを出力
	fmt.Println(string(data[:count]))     //読み込んだデータ（バイトデータ）を文字列にして出力
}
