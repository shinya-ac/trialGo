package main

import (
	"fmt"
)

type Number struct {
	Min int
}

type Sum_price struct {
	Sum int
}

//最小値を求める問題
func find_min() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4, 1}
	var num Number
	num.Min = l[0]
	for i := 0; i < len(l)-1; i++ {
		if l[i] < l[i+1] {
			if l[i] < num.Min {
				num.Min = l[i]
			}
		} else {
			if l[i+1] < num.Min {
				num.Min = l[i+1]
			}
		}
	}

	fmt.Println(num.Min)
}

//価格の合計を出す問題
func sum_price() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var sum Sum_price
	for _, v := range m {
		sum.Sum = sum.Sum + v
		fmt.Println(sum)
	}
	//fmt.Println("a")
}

func main() {
	find_min()
	sum_price()
}

// Q1 解答例

// package main

// import (
//     "fmt"
// )

// func main() {
//     l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

//     var min int

//     for i, num := range l {
//         if i == 0 {
//             min = num
//             continue
//         }

//         if min >= num {
//             min = num
//         }
//     }

//     fmt.Println(min)
// }
// Q2. 解答例

// package main

// import (
//     "fmt"
// )

// func main() {
//     m := map[string]int{
//         "apple":  200,
//         "banana": 300,
//         "grapes": 150,
//         "orange": 80,
//         "papaya": 500,
//         "kiwi":   90,
//     }

//     sum := 0
//     for _, v := range m {
//         sum += v
//     }

//     fmt.Println(sum)
// }
