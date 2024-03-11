# 環境構築

## 0. Go の環境構築

Go をインストールしていない場合は
[こちら](https://zenn.dev/777kkk/books/bb6b650b7ba677/viewer/045ce9)を参考に Go をインストールし、PATH を通す.

## 1. レポジトリのクローン

- ssh

```bash
git clone git@github.com:Satowo/todo_app.git
```

- http

```bash
git clone https://github.com/Satowo/todo_app.git
```

## 2. パッケージのインストール

```bash
go install github.com/cosmtrek/air@latest && \
go install github.com/gorilla/mux
```

## 3. .env ファイルの作成

ルートディレクトリに.env ファイルを作成し[notion](https://www.notion.so/Peach-Tech-bf5d1684fddf427d98e139b3f0a5ecea?pvs=4)の内容をコピペする.

## 3. サーバーを起動する

```bash
make start
```

## 4. DB （MySQL）に接続する

user = root <br>
database = todo_app_db <br>
host = localhost <br>
port = 3305 <br>
password = todo_pass <br>

でローカルの DB に接続する.

## 4. DB の初期値を投入する

DB の GUI ツールの SQL エディタに
/database/schema.sql をコピペして、初期値を投入する.

## 5. API 定義書を確認

http://localhost:8082 にアクセスして swagger UI が立ち上がるか確認する.

## 6. API サーバの疎通確認

curl コマンドで情報を取得できるか確認する.

```bash
curl -v -X GET localhost:8083/boards
```
