package main

import (
	"fmt"
)

type UserScoreDb struct {
	Name            string
	MathScore       int
	SocialScore     int
	ChemistoryScore int
}

func (u UserScoreDb) average() int {
	return (u.MathScore + u.SocialScore + u.ChemistoryScore) / 3
}

func hoge() *UserScoreDb {
	userScore := &UserScoreDb{Name: "mike", MathScore: 80, SocialScore: 70, ChemistoryScore: 60}
	return userScore
}

func main() {
	userScore := hoge() //UserScoreDbを返すhogeメソッド
	fmt.Println(userScore.average())
}
