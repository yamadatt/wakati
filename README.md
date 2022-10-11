形態素解析をするためのツール。

kagomeに文字列を渡して、形態素解析した結果を受け取る。

以下のように、dockerで80ポートを8888ポートにポートフォワードで外部公開する。dockerでkagomeを動かすことで、クライアントだけ作れば良いようになる。（はず）

```
docker run --rm -d -p 8888:80 --name kagomeAPI ikawaha/kagome server -http=":80"
```



以下で動くか確認する。

```
curl -s -X PUT localhost:8888/tokenize -d'{"sentence":"すもももももももものうち", "mode":"normal"}' | jq .
```



参考にしたURL

[Docker で日本語形態素解析を Kagome の Web API で手軽に利用する（分かち書き解析） \- Qiita](https://qiita.com/KEINOS/items/8b5e3a251430db89de3f)


上記の動作確認が終わったら、```main.go```を動かす。


