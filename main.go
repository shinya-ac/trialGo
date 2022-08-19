package main

import (
	"fmt"
	"strings"
	"time"
)

func goruotine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func main() {
	fmt.Println("hello world")
	fmt.Println("hello" + "world")
	fmt.Println(string("Hello world"[0]))

	var s string = "Hello World"

	fmt.Println(strings.Replace(s, "H", "X", 1))
}

//実行はmac右下の「Fnキー」を押しながら上にある「F5キー」を押すとできる。
