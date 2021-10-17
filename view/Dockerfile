FROM httpd:2.4

# node.jsなどをインストール
RUN apt-get update \
    && apt-get -y install --no-install-recommends \
    nodejs \
    npm \
    curl \
    # キャッシュ削除
    && apt-get clean

# node.jsを最新化
RUN npm install -g n \
    && n stable \
    && apt purge -y \
    nodejs \
    npm

# ワーキングディレクトリを作成
RUN mkdir -p /vue/

# プロジェクトのファイルやフォルダをワーキングディレクトリにコピー
COPY ./view/ /vue/

# カレントワーキングディレクトリとして '/vue' フォルダを指定
WORKDIR /vue/

# プロジェクトの依存ライブラリをインストール
RUN npm install

# 本番向けに圧縮しながらアプリケーションをビルド
RUN npm run build

# ビルドしたものをApacheのドキュメントルートにコピー
RUN cp -a /vue/dist/* /usr/local/apache2/htdocs/

# Apacheのポート番号を8081番に変更
RUN sed -i "s/Listen 80\$/Listen 8081/g" /usr/local/apache2/conf/httpd.conf

# Apacheで.htaccessを有効化
RUN sed -i "s/AllowOverride None/AllowOverride All/g" /usr/local/apache2/conf/httpd.conf

# Apacheのrewrite_moduleを有効化
RUN sed -i "s/#LoadModule rewrite_module modules\/mod_rewrite.so/LoadModule rewrite_module modules\/mod_rewrite.so/g" /usr/local/apache2/conf/httpd.conf

# RewiteRuleを設定
RUN cp /vue/.htaccess /usr/local/apache2/htdocs/.htaccess

# システムの8081番ポートを解放
EXPOSE 8081

# Apacheを再起動
RUN /usr/local/apache2/bin/apachectl -k restart