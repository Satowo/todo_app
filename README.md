# 環境構築

## 0. Go の環境構築

Go をインストールしていない場合は
[こちら](https://zenn.dev/777kkk/books/bb6b650b7ba677/viewer/045ce9)を参考に Go のインストール、PATH を通す

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

まず最初に開発に必要な以下のコマンドをインストールする.

- [docker](https://docs.docker.com/get-docker/)
- [make](https://www.gnu.org/software/make/)
- [mysql](https://dev.mysql.com/doc/refman/8.0/ja/mysql.html)
- [go](https://golang.org/doc/install)

## 3. .env ファイルの作成

ルートディレクトに.env ファイルを作成する.
（中身の値は吉田に聞いてください）

<!-- ## 3. DB の初期値を投入する

**_（現状 seed データは仮のデータで git 上で共有されているためこの部分は不要）_**
db の初期値をもらい, リポジトリのルートに置いて解凍する.

```bash
tar -zxvf seed.tar.gz
```

初期値 DB に投入する.

```bash
make run-seed
```

このコマンドは何度も実行できるので, 予期せぬエラーが発生したり, データベースの値を変えたりした場合には再度実行する. -->

## 4. サーバーを起動する

```bash
make start
```

<!-- ## 5. 疎通確認

```bash
curl -v http://localhost:1324/api/health
``` -->

## 5. API 定義書を確認（作成途中）

http://localhost:8082 にアクセスして swagger UI が立ち上がるか確認する.
