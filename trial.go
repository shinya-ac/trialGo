package main

//コメント
import (
	"fmt"
)

// 値渡しをするone()関数
func one(x int) {
	x = 1
}

func pointer_one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	//↑「これは0x0004というアドレスが差している数値型のメモリ（箱）に対して100という数値を格納して
	//その箱名前をnとする」というコード

	fmt.Println(n) //nという名前の箱の中にあるデータの本体（100）を出力
	//=>100
	fmt.Println(&n) //nという箱のあるアドレス（0x0004）を出力
	//=>0x0004

	var p *int = &n
	//↑これは&nという値（つまりnの位置を差しているアドレス）を0x0010というアドレスが
	//差しているポインタ型（正確には数値のポインタ型）のメモリ（箱）に格納している。
	//そしてそのメモリ（箱）の名前はpという風に名付けるという意味のコード

	fmt.Println(p) //pという名前の箱にあるデータの本体（0x0004）を出力
	//=>0x0004
	fmt.Println(*p) //pという名前の箱（ポインタ型の箱）の中身のアドレス（つまり0x0004）
	//が指し示しているデータの本体を出力
	//つまり今回の例で言うとnと同じ値になるはずの存在
	//=>100

	*p = 3000
	fmt.Println(*p)
	fmt.Println(n)

	fmt.Println("以下は値渡し")
	n = 100
	one(n)
	fmt.Println(n)

	fmt.Println("以下はポインタ渡し")
	n = 100
	pointer_one(&n)
	fmt.Println(n)
	fmt.Println(*&n)

	//①まずはnew()ありでポインタ型を宣言するバージョン
	var q *int = new(int)
	fmt.Println(*q)
	//=>0　　　←初期値(デフォルト値)が入っている
	*q++
	fmt.Println(*q)
	//=>1　　　←初期値に+1された数が入っている

	//②以下はnew()なしでポインタ型を宣言するバージョン
	var q2 *int
	fmt.Println(q2)
	//=>0　　　←nilが入っている
	//*q2++
	//fmt.Println(q2)
	//=>error　　　←nilに対して+1しているのでエラーになる

	fmt.Println("型の違い")
	s := make([]int, 0)
	fmt.Printf("%T\n", s)
	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var pointer *int = new(int)
	fmt.Printf("%T\n", pointer)
	var structure = new(struct{})
	fmt.Printf("%T\n", structure)
}
