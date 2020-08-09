# 開発環境構築

```
docker-compose up -d
mysql -h 127.0.0.1 -uroot -p -P 3306 gin < docker/mysql_helper/init.sql
mysql -h 127.0.0.1 -uroot -p -P 3306 gin < docker/mysql_helper/data.sql
```

## serverside

```
export ENV=dev; go run main.go
```
localhost:8080にアクセスすることでサーバー側の確認ができます。

## frontend

```
yarn dev
```

ブラウザでlocalhost:3000にアクセスするとフロント側を確認できます。

## E2E

```
docker-compose up -d
cd ./docker/mysql_helper/
./init.sh
```

# xorm

## reverse
$ xorm reverse mysql root:root@/gin?charset=utf8mb4 $GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm


# ref

[Golang のパッケージ完全に理解した ← わかってない - くろのて](https://note.crohaco.net/2019/golang-package/)
[ginを最速でマスターしよう - Qiita](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)