package main

import (
	"fmt"
	"math"
)

func create_card_set(all_cards []int, number_of_sets_int, M int) [][]int {
	var all_cards_set [][]int
	var card_set []int

	for j := 0; j < number_of_sets_int; j++ {
		if len(all_cards) < M {
			all_cards_set = append(all_cards_set, all_cards)
		} else {
			card_set = all_cards[:M]

			all_cards = all_cards[M:]

			all_cards_set = append(all_cards_set, card_set)
		}
	}
	return all_cards_set
}

func shuffle_card_set(shuffled_all_cards_set [][]int) [][]int {
	for i := 0; i < len(shuffled_all_cards_set)/2; i++ {
		shuffled_all_cards_set[i], shuffled_all_cards_set[len(shuffled_all_cards_set)-i-1] = shuffled_all_cards_set[len(shuffled_all_cards_set)-i-1], shuffled_all_cards_set[i]
	}
	return shuffled_all_cards_set
}

func main() {
	//Nは全カードの枚数
	//Mは1セットのカード数
	//Kはシャッフルをする回数
	var N, M, K int
	fmt.Scan(&N, &M, &K)

	var floatN, floatM float64
	floatN = float64(N)
	floatM = float64(M)

	var number_of_sets float64 = math.Ceil(floatN / floatM)
	var number_of_sets_int int = int(number_of_sets)

	var all_cards []int
	for i := 1; i <= N; i++ {
		all_cards = append(all_cards, i)
	}

	// シャッフルの回数分回す
	for s := 0; s < K; s++ {
		all_cards_set := create_card_set(all_cards, number_of_sets_int, M)

		all_cards_set = shuffle_card_set(all_cards_set)
		var arr []int

		for i := 0; i < len(all_cards_set); i++ {
			for j := 0; j < len(all_cards_set[i]); j++ {
				arr = append(arr, all_cards_set[i][j])
			}
		}
		all_cards = arr
	}

	for b := 0; b < len(all_cards); b++ {
		fmt.Println(all_cards[b])
	}
}
