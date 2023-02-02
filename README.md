# goSimpleApi

goでDBを使わないシンプルなAPIを作成するための練習

## 起動方法

```bash
docker-compose up -d
```

最初にDB migrateしていますので起動するたびにデータも飛びます

## 使い方

```bash
curl -X POST http://localhost:8080/create \
    -H "Content-Type: application/json" -d '{"price":10000, "code":"hoge"}'
```

```bash
curl http://localhost:8080/product/1
```



## 検証方法

POSTMANで検証可能。

例: GET: http://localhost:8000/api/books など

## 参考文献

[Goで超簡単API \- Qiita](https://qiita.com/k-penguin-sato/items/8088b69304ee7e8f70be)

## やろうとしていること

- apiをつくる練習
- 一旦ORMをつくってデータをいい感じに出力する方法を色々ためしたい
- Webフレームワークを使ってもいいかもしれない

### ということで何をやるか

- まずはどんなORMがあるのかとかを調査して試してみる
- [Goの初心者が見ると幸せになれる場所　#golang - Qiita](https://qiita.com/tenntenn/items/0e33a4959250d1a55045)
