FROM node:lts-alpine

# 静的コンテンツを配信するシンプルな http サーバをインストールする
RUN npm install -g http-server

# ワーキングディレクトリを作成する
RUN mkdir -p /vue/

# プロジェクトのファイルやフォルダをワーキングディレクトリにコピーする
COPY ./view/ /vue/

# カレントワーキングディレクトリとして '/vue/view' フォルダを指定する
WORKDIR /vue/

# プロジェクトの依存ライブラリをインストールする
RUN npm install

# 本番向けに圧縮しながらアプリケーションをビルドする
RUN npm run build

# 8081番ポートを解放する
EXPOSE 8081

CMD [ "http-server","--cors","-p8081","dist" ]