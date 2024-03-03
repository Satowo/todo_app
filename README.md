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
- go install github.com/cosmtrek/air@latest
```

## 3. .env ファイルの作成

ルートディレクトに.env ファイルを作成する.
（中身の値は吉田に聞く）

## 3. サーバーを起動する

```bash
make start
```

## 4. DB （MySQL）に接続する

user = root
database = todo_app_db
host = localhost
port = 3306
password = todo_pass

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
