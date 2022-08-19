package main

import "fmt"

type BeetlesCoordinate struct {
	redBeetleCoordinate   int
	greenBeetleCoordinate int
	blueBeetleCoordinate  int
}

// 三匹のカブトムシの初期座標設定
func (b *BeetlesCoordinate) Set(redBeetleCoordinate, greenBeetleCoordinate, blueBeetleCoordinate int) {
	b.redBeetleCoordinate = redBeetleCoordinate
	b.greenBeetleCoordinate = greenBeetleCoordinate
	b.blueBeetleCoordinate = blueBeetleCoordinate
}

func (b *BeetlesCoordinate) Moov(direction, color string) {
	if direction == "L" {
		switch color {
		case "R":
			b.redBeetleCoordinate -= 1
		case "G":
			b.greenBeetleCoordinate -= 1
		case "B":
			b.blueBeetleCoordinate -= 1
		case "Y": //赤と緑のカブトムシ
			b.redBeetleCoordinate -= 1
			b.greenBeetleCoordinate -= 1

		case "M": //赤と青のカブトムシ
			b.redBeetleCoordinate -= 1
			b.blueBeetleCoordinate -= 1

		case "C": //青と緑のカブトムシ
			b.blueBeetleCoordinate -= 1
			b.greenBeetleCoordinate -= 1

		case "W": //全部のカブトムシ
			b.redBeetleCoordinate -= 1
			b.greenBeetleCoordinate -= 1
			b.blueBeetleCoordinate -= 1
		}
	} else {
		//case1:白→全部が+1
		//case2:...
		switch color {
		case "R":
			b.redBeetleCoordinate += 1
		case "G":
			b.greenBeetleCoordinate += 1
		case "B":
			b.blueBeetleCoordinate += 1
		case "Y": //赤と緑のカブトムシ
			b.redBeetleCoordinate += 1
			b.greenBeetleCoordinate += 1

		case "M": //赤と青のカブトムシ
			b.redBeetleCoordinate += 1
			b.blueBeetleCoordinate += 1

		case "C": //青と緑のカブトムシ
			b.blueBeetleCoordinate += 1
			b.greenBeetleCoordinate += 1

		case "W": //全部のカブトムシ
			b.redBeetleCoordinate += 1
			b.greenBeetleCoordinate += 1
			b.blueBeetleCoordinate += 1
		}

		// b.redBeetleCoordinate = redBeetleCoordinate
		// b.greenBeetleCoordinate = greenBeetleCoordinate
		// b.blueBeetleCoordinate = blueBeetleCoordinate
	}
}

// trueなら3匹のカブトムシの座標が同じ
func (b *BeetlesCoordinate) Check() bool {
	if b.redBeetleCoordinate == b.greenBeetleCoordinate && b.greenBeetleCoordinate == b.blueBeetleCoordinate {
		return true
	}
	return false
}

func main() {
	// N:命令の個数を表す整数 N
	var N int
	fmt.Scan(&N)

	// それぞれのカブトムシの座標を取得する
	var x_R, x_G, x_B int
	fmt.Scan(&x_R, &x_G, &x_B)

	var b BeetlesCoordinate
	//座標初期化
	b.Set(x_R, x_G, x_B)

	// direction:ライトの当たる方向（左、右）
	// color:ライトの色
	var direction, color string
	for i := 0; i < N; i++ {
		fmt.Scan(&direction, &color)
		b.Moov(direction, color)

		if b.Check() == true {
			fmt.Println(b.redBeetleCoordinate)
			return
			//returnでifもforも強制終了＝＞「fmt.Println("no")」に辿り着かない
		}
	}

	// 重ならなかったとき
	fmt.Println("no")
}

//最終：同じ位置に集まる→最初に同じ位置になった座標（~3~3のどっか）を出力
//同じ位置に集まらない→noを出力

//同じ位置かどうかの判定→構造体の三匹の座標（数値）が一致したかどうかで判定

//道の座標↓
//-100,....-3,-2,-1,0,1,2,3,...100
// L: -1にする
// R: +1にする

//カブトムシの座標
// {
// 	R: 3
// 	G: 3
// 	B: 3
// }

//色一覧
//"R", "G", "B", "Y", "M",    "C",    "W"
//赤、  緑、  青、 黄、  マゼンタ、シアン、 白

//type Animal struct {
//    Name string
//   Age  int
//}

// a := Animal{}
// var a Animal = Animal{} // 上記と同じ意味
// var a Animal           //  上記と同じ意味

// a.Name = "hoge"
// a.Age = 5

// func (a *Animal) Set(Name string, Age int) {
//     a.Name = name
//     a.Age = age
// }

//p.Set("Tanaka", 31)
