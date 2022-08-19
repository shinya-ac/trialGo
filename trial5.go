package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//まずはURLの定義
	base, _ := url.Parse("https://example.com")           //ベースURLをパースして不正でないかチェックする
	reference, _ := url.Parse("/test?a=1&b=2")            //クエリ部分をパースして不正でないかチェックする
	endpoint := base.ResolveReference(reference).String() //ベースURLとクエリをがっちゃんこ
	fmt.Println(endpoint)

	//からのリクエストを作る
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password"))) //この時点ではまだリクエストを投げていない
	req.Header.Add("hogehoge", "fugafuga")                                           //header情報を追記

	query := req.URL.Query() //クエリの情報は「req.URL.Query()」に入っているので、これを見る
	fmt.Println(query)

	query.Add("c", "3&%")
	fmt.Println(query)          //エンコードする前
	fmt.Println(query.Encode()) //エンコードした後

	req.URL.RawQuery = query.Encode() //エンコードした上で元のURLに置き換えた

	var client *http.Client = &http.Client{} //「*http.Client型」のクライアントをhttp.Client{}という構造体
	//で初期化した
	resp, _ := client.Do(req) //この段階で実際にこのクライアントでURLにGETでアクセスできる
	body, _ := ioutil.ReadAll(resp.Body)
	println(string(body))

}
