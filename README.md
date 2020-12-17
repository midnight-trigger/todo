# 使用技術
* バックエンド：Golang1.14（Echo）
* ORM：GORM
* DB：MySQL8.0.22
* 認証基盤：Cognito
* API基盤：API Gateway

# API概要
https://junya-it-nishi.atlassian.net/wiki/spaces/JN/pages/294913/Todo+REST+API

※Confluenceのログインが必要

# マイグレーション方法
 https://github.com/golang-migrate/migrate

 golang-migrate/migrateを利用する

```
# Macでbrewをお使いの場合
brew install golang-migrate
```

マイグレーションコマンド
```
migrate -database 'mysql://root:password@tcp(0.0.0.0:3306)/<DB名>' -path migrations up
```
※ golang-migrate/migrateでまだ管理されていないテーブルに流し込む場合は、いったんDBを空にしてから上記コマンドを実行する

 schema_migrationsテーブルが作成されてそこで実行管理される。

 マイグレーションファイルの作成
 ```
 migrate create -ext sql -dir migrations/ -seq <ファイル名>
 ```
