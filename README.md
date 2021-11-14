使い方

デバッグ時にサイトを巡回するというのを避けたいため、フィアルをあらかじめダウンロードしておく。

## サイトマップからダウンロードする

```
go run get-sitemap-url.go
```

既にダウンロードしたフィアルを再び取得するのを避けるため、差分をdiffする。

```
diff old_list new_list
```

差分のみをwgetでダウンロードする。負荷を考慮して、2秒ウエイト。

```
wget -i list.txt -w 2 --user-agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"
```

## movable typeのインポート形式に変換

```
go run make_movable.go
```

