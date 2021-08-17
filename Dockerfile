# ベースイメージ作成。(現在の最新バージョン)
FROM golang:1.16.7

#vimを利用するため。
RUN apt-get update && apt-get -y install vim

# カレントディレクトリの内容をコピーする
COPY . /go/src/app

# コンテナ内で各種命令を実行するためのカレントディレクトリを指定
WORKDIR /go/src/app

# ライブラリをインストールする
RUN go mod download

# イメージを実行する時、コンテナに対して何もオプションを指定しなければ、
# 自動的に実行するコマンドを CMD 命令で指定
CMD ["go","run","/go/src/app/main.go"]