package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Seat struct {
	height_location int
	width_location  int
	m_distance      int
}

func minInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[0]
}

func main() {
	var N, H, W, P, Q int
	//すでに予約されている座席の数 N、
	//映画館の座席の縦の数 H、映画館の座席の横の数 W、
	//最も見やすい席の p 座標、q 座標を表す P, Q
	fmt.Scan(&N, &H, &W, &P, &Q)
	fmt.Println(N, H, W, P, Q)

	var seats []Seat = []Seat{}
	var reserved_seats []Seat = []Seat{}
	//var m_distance []int

	//全席の構造体のスライスを作成
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {

			seat1 := Seat{height_location: i, width_location: j}
			seats = append(seats, seat1)
			//fmt.Println(seats)
		}
	}

	fmt.Println(seats)
	fmt.Println("三番目の席を出力")
	fmt.Println(seats[2])

	//予約済みの席の構造体のスライスを生成
	var reserved_p, reserved_q string
	for i := 0; i < N; i++ {
		fmt.Scan(&reserved_p, &reserved_q)
		int_p, _ := strconv.Atoi(reserved_p)
		int_q, _ := strconv.Atoi(reserved_q)
		seat := Seat{height_location: int_p, width_location: int_q}
		reserved_seats = append(reserved_seats, seat)
	}

	var available_seats []Seat = []Seat{}
	//全席から予約済みの席を引いた席一覧のスライス
	for i := 0; i < H*W; i++ {
		for k := 0; k < len(reserved_seats); k++ {
			if seats[k].height_location == reserved_seats[k].height_location && seats[k].width_location == reserved_seats[k].width_location {
				return
			} else {
				for q := 0; q < len(available_seats); q++ {
					if available_seats[q].height_location == seats[i].height_location && available_seats[q].width_location == seats[i].width_location {
						return
					} else {
						available_seats = append(available_seats, seats[i])
					}
				}
			}
		}
	}

	for i := 0; i < len(available_seats); i++ {
		fmt.Println("予約可能な席の一覧")
		fmt.Println(available_seats[i])
	}

	//全席のマンハッタン距離を算出して、その席（構造体）のm_distanceに代入
	for i := 0; i < len(seats); i++ {
		var a, b int
		if P > seats[i].height_location {
			a = P - seats[i].height_location
		} else if P < seats[i].height_location {
			a = seats[i].height_location - P
		}
		if Q > seats[i].width_location {
			b = Q - seats[i].width_location
		} else if Q < seats[i].width_location {
			b = seats[i].width_location - Q
		}

		y := a + b
		seats[i].m_distance = y
	}

	var m_distances []int
	for i := 0; i < len(seats); i++ {
		m_distances = append(m_distances, seats[i].m_distance)
	}

	fmt.Println("マンハッタン距離の配列")
	fmt.Println(m_distances)

	var l int

	//全席のうちの各席のマンハッタン距離(m_distance)を比較して最初値を求めようとした（挫折）
	l = minInt(m_distances)

	fmt.Println("マンハッタン距離で算出した一番近い距離")
	fmt.Println(l)

	//最も見やすい席の構造体の配列
	var best_seats []Seat = []Seat{}

	//マンハッタン距離で算出した一番近い距離（l）と一致するseatsのm_distanceを持つ構造体を出す
	for i := 0; i < len(seats); i++ {
		if l == seats[i].m_distance {
			best_seats = append(best_seats, seats[i])
		}
	}
	fmt.Println("一番近い席の構造体一覧")
	fmt.Println(best_seats)

	for i := 0; i < len(best_seats); i++ {
		//ベストな席の座標
		for i := 0; i < len(reserved_seats); i++ {
			if best_seats[i].height_location == reserved_seats[i].height_location && best_seats[i].width_location == reserved_seats[i].width_location {
				return
			} else {
				fmt.Println(best_seats[i].height_location)
				fmt.Println(best_seats[i].width_location)
			}
		}
	}

	// func minInt(a []int) int {
	// 	sort.Sort(sort.IntSlice(a))
	// 	return a[0]
	// }

	//全席の回数（つまりH*W回）for文を回して「各席のマンハッタン距離」を出す。
	//マンハッタン距離が最も小さい席の一覧を出す。
	//その一覧のスライスから予約済みの席のスライスをひく
	//残った席の一覧を出力する。
	//全席の構造体のスライス-予約済みの席の構造体のスライス=自分が予約可能な席一覧

	fmt.Println(reserved_seats)

	//
}

//最終的には最も映画が見やすい席（の座標）を　全て出力するという問題
//→どの席も見やすくなかったとしても、何かしらの席はある想定？
//映画が見やすい席＝最も見やすい席とのマンハッタン距離が近い席
//マンハッタン距離＝＞ある 2 点 (p, q)、(s, t) に対して、|p - s| + |q - t| で与えられる距離
//つまりマンハッタン距離の値が一番小さい席を出力するという問題

//予約されている座標は省く
//残った座標でマンハッタン距離を出す
//全ての席は「0,0〜H-1,W-1」まである。数字としてはH*Wの値

// 入力例1
// 9 4 5 2 3
// 1 0
// 1 2
// 1 3
// 1 4
// 2 2
// 2 3
// 2 4
// 3 3
// 3 4

// 出力例1
// 0 3
// 2 1
// 3 2
// 入力例2
// 4 3 2 2 0
// 0 0
// 1 0
// 1 1
// 2 1
// 出力例2
// 2 0
