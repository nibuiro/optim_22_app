# optim_22_app

## ディレクトリ構成

```
.
├──configs               #各種設定ファイル（例：app.yaml, zap.yaml）。
├──controller               
│  ├──request
│  ├──top
│  ├──user     
│  │   └──api               #APIエンドポイントハンドラ。リクエストの正常性確認後serviceに渡す。
│  └── ...             
├──model
│  ├──user
│  │   ├──repository        #DBとの通信処理。ORMを⽤いたスクリプトなど。
│  │   └──service           #リクエストの内容がAPIの仕様に則しているか、repoを⽤いた⼀連の処理。
│  └── ...
├──mysql
├──pkg
│  ├──config
│  └──log
├──test
├──typefile
├──view
│   └──top
└── ...
```


## 環境構築手順

step1 optim-22-appのリポジトリをcloneする。  
```
git clone https://github.com/optim22team/optim_22_app.git
```
step2 .env,.gitignoreファイルを生成する。  
```
touch .env .gitignore
```
step3 docker-imageを作成する。  
```
docker-compose build
```
step4 docker-containerを作成する。  
```
docker-compose up -d
```
step5 go_container,mysql_containerが立ち上がっていることを確認する。  
```
docker ps
```
step6 データが出力されていることを確認する。(curl: (52) Empty reply from serverエラー,あるいは{"requests":null}が出力された場合、step7以降の作業が必要となる。)  
```
curl http://localhost:8080/
```
step7 mysql_containerに入る。  
```
docker exec -it mysql_container_ID bash
```
step8 mysqlのログインをする。  
```
mysql -u root -p
```
step9 mysqlのパスワードを入力する。(ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/var/run/mysqld
/mysqld.sock'),あるいは (ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: YES))というエラーが生じた場合、必要なファイルが生成されるまで待つ。  
```
rootpass
```
step10 mysqlへのログインが成功したら、mysqlをログアウトする。  
```
exit
```
step11 mysql-containerをログアウトする。  
```
exit
```
step12 go_containerに入る。  
```
docker exec -it go_container_ID bash
```
step13 カレントディレクトリで以下のコマンドを実行する。  
```
go run main.go
```
step14 go-containerをログアウトする。  
```
exit
```
step15 データが出力されていることを確認する。  
```
curl http://localhost:8080/
```