# サーバーエンジニア向け 2025新卒採用事前課題

あなたは歌手とアルバムを管理するAPIの機能開発にたずさわることになりました。

次の課題に順に取り組んでください。

できない課題があっても構いません。

面接中に課題に関して質問をしますので、分かる範囲で説明してください。

## 課題1
プログラムのコードを読み、中身を把握しましょう。

### 解答
- api：HTTPリクエストのエンドポイント（API）を定義し、それに対するハンドラーやルーターを提供
- controller：HTTPリクエストを受け取り、サービス層を呼び出して処理を行い、最終的にHTTPレスポンスを生成
- infra：データベースや外部サービスへのアクセス、永続化などの実装
- model：アプリケーション内で使用されるデータ構造やモデルを定義
- repository：データベースや外部データストアへのクエリやデータの永続化のためのインターフェースや実装
- service：コントローラーから呼び出され、データの操作や処理を行う




## 課題2
go をインストールし(各自で調べてください)、歌手を管理するAPIの動作を確認しましょう。

```
# (ターミナルを開いて)
# サーバーを起動する
go run main.go
```

```
# (別のターミナルを開いて)
# 歌手の一覧を取得する
curl http://localhost:8888/singers

# 指定したIDの歌手を取得する
curl http://localhost:8888/singers/1

# 歌手を追加する
curl -X POST -d '{"id":10,"name":"John"}' http://localhost:8888/singers

# 歌手を削除する
curl -X DELETE http://localhost:8888/singers/1
```

### 解答
api_logフォルダに格納

## 課題3
アルバムを管理するAPIを新規作成しましょう。

### 3-1
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer_id":1},{"id":2,"title":"Alice's 2nd Album","singer_id":1},{"id":3,"title":"Bella's 1st Album","singer_id":2}]
```

### 3-2
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer_id":1}
```

### 3-3
アルバムを追加するAPI
```
curl -X POST -d '{"id":10,"title":"Chris 1st","singer_id":3}' http://localhost:8888/albums

# このようなレスポンスを期待しています
{"id":10,"title":"Chris 1st","singer_id":3}

# そして、アルバムを取得するAPIでは、追加したものが存在するように
curl http://localhost:8888/albums/10
```

### 3-4
アルバムを削除するAPI
```
curl -X DELETE http://localhost:8888/albums/1
```

## 課題4
アルバムを取得するAPIでは、歌手の情報も付加するように改修しましょう。

### 4-1
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}
```

### 4-2
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]
```

## 課題5
歌手とそのアルバムを管理するという点で、現状のAPIの改善点を検討し思いつく限り書き出してください。

実装をする必要はありません。
