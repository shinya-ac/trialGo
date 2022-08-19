package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

func (v Vertex) plus_ten() {
	v.X = v.X + 10
	v.Y = v.Y + 10
	return
}

func (v *Vertex) plus_ten_pointer() {
	v.X = v.X + 10
	v.Y = v.Y + 10
	return
}

//コンストラクタメソッド↓
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

//以下は平面のVertexというストラクトを利用して、それを元に立体のVertex3Dを作成する工程
//(Embeddedという。オブジェクト指向でいうところの「継承」にあたる)

type Vertex3D struct {
	//X, Y intではなく、この二つの性質を持つVertexという構造体をそのままこの構造体に入れ込んだ
	Vertex
	Z int
}

func (v Vertex3D) Area3D() int {
	return v.X * v.Y * v.Z
}

func (v *Vertex3D) plus_ten_pointer3D() {
	v.X = v.X + 10
	v.Y = v.Y + 10
	v.Z = v.Z + 10
	return
}

func main() {
	v := Vertex{3, 4}
	//fmt.Println(Vertex.X)←構造体自体にはXとかYとかの内容（具体的な値）は存在しない
	fmt.Println(v.X, v.Y)

	fmt.Println(v.Area()) //面積

	v.plus_ten()
	fmt.Println(v.X, v.Y) //このplus_tenメソッドは値渡しなので元のv.X,v.Yに影響を与えない

	v.plus_ten_pointer()
	fmt.Println(v.X, v.Y) //このplus_ten_pointerメソッドはポインタ渡しなので元のv.X,v.Yに影響を与える

	k := New(8, 9) //コンストラクタでNew()を呼び出して初期化
	fmt.Println(k.X, k.Y)

	//3D
	v2 := Vertex3D{Vertex{3, 4}, 5} //ここがちょい独特
	//v2 := Vertex3D{3, 4, 5}ではないことに注意
	fmt.Println(v2.Area3D()) //3Dの体積
	v2.plus_ten_pointer3D()
	fmt.Println(v2.X, v2.Y, v2.Z)
}
