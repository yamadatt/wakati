package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)


func main() {
    fmt.Println("Start!")

    url := "http://localhost:8888/tokenize"
    authHeaderName := "x-cdata-authtoken"
    authHeaderValue := "7y3E6q4b6V1v9f0D2m9j"

    req, _ := http.NewRequest(http.MethodGet, url, nil)
    req.Header.Set(authHeaderName, authHeaderValue)

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
    fmt.Printf("%-v", resp)

    // Response Body を読み取り
    body, _ := io.ReadAll(resp.Body)

    // JSONを構造体にエンコード
    var Books TokenizedMessages
    json.Unmarshal(body, &Books)

    fmt.Printf("%-v", Books)

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