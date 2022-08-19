package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct { //情報を受け取る構造体を定義
	Title string
	Body  []byte
}

func (p *Page) save() error { //情報をファイルに保存するメソッドを定義
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

func viewHundler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s<div>", p.Title, p.Body)

}

func main() {
	http.HandleFunc("/view/", viewHundler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type UserDb struct {
	Age      int
	BirthDay DateTime
	Name     string
	NickName string
}
