package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {


	url := "http://localhost:8888/tokenize"

	// 分かち書きする文字列をファイルから読み込ませる

	targettext := FileRead("sample.txt")

	reqBody := RequestBody{
		Sentence: targettext,
		Mode:     "normal",
	}

	// JSONに変換
	jsonValue, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(jsonValue)))

	// ヘッダーをセット
	req.Header.Set("Content-type", "application/json;charset=UTF-8")

	client := new(http.Client)
	resp, err := client.Do(req)

	// URLがnilだったり、Timeoutが発生した場合にエラーを返す模様。
	// サーバーからのレスポンスとなる 401 Unauthroized Error などはResponseをチェックする。
	// サーバーとの疎通が開始する前の動作のよう。
	if err != nil {
		fmt.Println("Error Request:", err)
		return
	}
	// resp.Bodyはクローズすること。クローズしないとTCPコネクションを開きっぱなしになる。
	defer resp.Body.Close()

	// 200 OK 以外の場合はエラーメッセージを表示して終了
	if resp.StatusCode != 200 {
		fmt.Println("Error Response:", resp.Status)
		return
	}

	// とりあえずResponsの構造体を全部出力
	//fmt.Printf("%-v", resp)

	// Response Body を読み取り
	body, _ := ioutil.ReadAll(resp.Body)

	// JSONを構造体にエンコード
	var messages TokenizedMessages

	if err := json.Unmarshal(body, &messages); err != nil {
		fmt.Println("Error Unmarshal:", err)
	}

	//fmt.Printf("%+v\n", messages)

	for _, v := range messages.Tokens {

		fmt.Printf("%s %+v\n", v.Surface, v.Pos)

	}

}

func FileRead(fileName string) (str string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	str = (string(bytes))
	return str
}

type TokenizedMessages struct {
	Status bool `json:"status"`
	Tokens []struct {
		ID            int      `json:"id"`
		Start         int      `json:"start"`
		End           int      `json:"end"`
		Surface       string   `json:"surface"`
		Class         string   `json:"class"`
		Pos           []string `json:"pos"`
		BaseForm      string   `json:"base_form"`
		Reading       string   `json:"reading"`
		Pronunciation string   `json:"pronunciation"`
		Features      []string `json:"features"`
	} `json:"tokens"`
}
type RequestBody struct {
	Sentence string `json:"sentence"`
	Mode     string `json:"mode"`
}
