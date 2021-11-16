形態素解析をするためのツール。

kagomeに文字列を渡して、形態素解析した結果を受け取る。

以下のように、dockerで80ポートを8888ポートにポートフォワードで外部公開する。

```
$ docker run --rm -d -p 8888:80 --name kagomeAPI ikawaha/kagome server -http=":80"
```


参考にしたURL

[Docker で日本語形態素解析を Kagome の Web API で手軽に利用する（分かち書き解析） \- Qiita](https://qiita.com/KEINOS/items/8b5e3a251430db89de3f)


```
&{200 OK 200 HTTP/1.1 1 1 map[Content-Length:[1922] Content-Type:[application/json] Date:[Tue, 16 Nov 2021 13:25:28 GMT]] 0xc00018a080 1922 [] false false map[] 0xc00014c000 <nil>}{Status:true Tokens:[{ID:36163 Start:0 End:3 Surface:すもも Class:KNOWN Pos:[ 
名詞 一般 * *] BaseForm:すもも Reading:スモモ Pronunciation:スモモ Features:[名詞 一般 * * * * すもも スモモ スモモ]} {ID:73244 Start:3 End:4 Surface:も Class:KNOWN Pos:[助詞 係助詞 * *] BaseForm:も Reading:モ Pronunciation:モ Features:[助詞 係助詞 * * * *  
も モ モ]} {ID:74988 Start:4 End:6 Surface:もも Class:KNOWN Pos:[名詞 一般 * *] BaseForm:もも Reading:モモ Pronunciation:モモ Features:[名詞 一般 * * * * もも モモ モモ]} {ID:73244 Start:6 End:7 Surface:も Class:KNOWN Pos:[助詞 係助詞 * *] BaseForm:も Reading:モ Pronunciation:モ Features:[助詞 係助詞 * * * * も モ モ]} {ID:74988 Start:7 End:9 Surface:もも Class:KNOWN Pos:[名詞 一般 * 
*] BaseForm:もも Reading:モモ Pronunciation:モモ Features:[名詞 一般 * * * * もも モモ モモ]} {ID:55829 Start:9 End:10 Surface:の
 Class:KNOWN Pos:[助詞 連体化 * *] BaseForm:の Reading:ノ Pronunciation:ノ Features:[助詞 連体化 * * * * の ノ ノ]} {ID:8027 Start:10 End:12 Surface:うち Class:KNOWN Pos:[名詞 非自立 副詞可能 *] BaseForm:うち Reading:ウチ Pronunciation:ウチ Features:[名詞 非
自立 副詞可能 * * * うち ウチ ウチ]} {ID:98 Start:12 End:13 Surface:。 Class:KNOWN Pos:[記号 句点 * *] BaseForm:。 Reading:。 Pronunciation:。 Features:[記号 句点 * * * * 。 。 。]}]}
```

以下で動くか確認する。

```
curl -s -X PUT localhost:8888/tokenize -d'{"sentence":"すもももももももものうち", "mode":"normal"}' | jq .
```

返却されるJSON

```
{
  "status": true,
  "tokens": [
    {
      "id": 36163,
      "start": 0,
      "end": 3,
      "surface": "すもも",
      "class": "KNOWN",
      "pos": [
        "名詞",
        "一般",
        "*",
        "*"
      ],
      "base_form": "すもも",
      "reading": "スモモ",
      "pronunciation": "スモモ",
      "features": [
        "名詞",
        "一般",
        "*",
        "*",
        "*",
        "*",
        "すもも",
        "スモモ",
        "スモモ"
      ]
    },
    {
      "id": 73244,
      "start": 3,
      "end": 4,
      "surface": "も",
      "class": "KNOWN",
      "pos": [
        "助詞",
        "係助詞",
        "*",
        "*"
      ],
      "base_form": "も",
      "reading": "モ",
      "pronunciation": "モ",
      "features": [
        "助詞",
        "係助詞",
        "*",
        "*",
        "*",
        "*",
        "も",
        "モ",
        "モ"
      ]
    },
    {
      "id": 74988,
      "start": 4,
      "end": 6,
      "surface": "もも",
      "class": "KNOWN",
      "pos": [
        "名詞",
        "一般",
        "*",
        "*"
      ],
      "base_form": "もも",
      "reading": "モモ",
      "pronunciation": "モモ",
      "features": [
        "名詞",
        "一般",
        "*",
        "*",
        "*",
        "*",
        "もも",
        "モモ",
        "モモ"
      ]
    },
    {
      "id": 73244,
      "start": 6,
      "end": 7,
      "surface": "も",
      "class": "KNOWN",
      "pos": [
        "助詞",
        "係助詞",
        "*",
        "*"
      ],
      "base_form": "も",
      "reading": "モ",
      "pronunciation": "モ",
      "features": [
        "助詞",
        "係助詞",
        "*",
        "*",
        "*",
        "*",
        "も",
        "モ",
        "モ"
      ]
    },
    {
      "id": 74988,
      "start": 7,
      "end": 9,
      "surface": "もも",
      "class": "KNOWN",
      "pos": [
        "名詞",
        "一般",
        "*",
        "*"
      ],
      "base_form": "もも",
      "reading": "モモ",
      "pronunciation": "モモ",
      "features": [
        "名詞",
        "一般",
        "*",
        "*",
        "*",
        "*",
        "もも",
        "モモ",
        "モモ"
      ]
    },
    {
      "id": 55829,
      "start": 9,
      "end": 10,
      "surface": "の",
      "class": "KNOWN",
      "pos": [
        "助詞",
        "連体化",
        "*",
        "*"
      ],
      "base_form": "の",
      "reading": "ノ",
      "pronunciation": "ノ",
      "features": [
        "助詞",
        "連体化",
        "*",
        "*",
        "*",
        "*",
        "の",
        "ノ",
        "ノ"
      ]
    },
    {
      "id": 8027,
      "start": 10,
      "end": 12,
      "surface": "うち",
      "class": "KNOWN",
      "pos": [
        "名詞",
        "非自立",
        "副詞可能",
        "*"
      ],
      "base_form": "うち",
      "reading": "ウチ",
      "pronunciation": "ウチ",
      "features": [
        "名詞",
        "非自立",
        "副詞可能",
        "*",
        "*",
        "*",
        "うち",
        "ウチ",
        "ウチ"
      ]
    }
  ]
}
```
