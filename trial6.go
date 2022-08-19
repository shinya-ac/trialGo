package main

import (
	"encoding/json"
	"fmt"
)

type PersonStruct struct {
	Name      string   `json:"XXXXXXXXXX"`
	Age       int      `json:"age,string"`
	Nicknames []string `json:"nicknames"`
}

func main() {
	b := []byte(`{"name": "Mike", "age": "20", "nicknames":["a", "b", "c"]}`)
	var p PersonStruct
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))

	m := make(map[string]int, 2)

	m["a"] = 100
	m["b"] = 200
	m["add"] = 300

	fmt.Println(m)

}
