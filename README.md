# 使用技術
* バックエンド：Golang1.14（Echo）
* ORM：GORM
* DB：MySQL8.0.22
* 認証基盤：Cognito
* API基盤：API Gateway

# API概要
https://junya-it-nishi.atlassian.net/wiki/spaces/JN/pages/294913/Todo+REST+API

※Confluenceのログインが必要

# ディレクトリ構成
```
├── api
│   ├── controller
│   ├── definition
│   ├── domain
│   └── error_handing
├── configs
├── dev_tools
├── infra
├── logger
└── third_party
```
* `api` : API関係
   * controller： API Controller
   * definition： API リクエスト/レスポンス構造体定義
   * domain： API ビジネスロジック
   * error_handing： API エラー処理
   * route.go, route_gen/go: ルーティング関連
* `configs` : 環境毎（local, dev, qa, stg prod etc...）の定義ファイルなど
* `dev_tools` : 開発効率化ためのツールなど
* `infra` : DB、Cognito関連の定義と処理
* `logger` : log処理
* `third_party` : 使用するライブラリ関係の処理など

# API動作確認手順
※MacOSを使う前提で書いています。

（１）ローカルにGolangの開発環境を整える
https://code-graffiti.com/set-golang-environment-on-mac-with-vscode/

（２）設定した`$GOPATH/src`配下に本APIをcloneする
（３）パッケージをインストールする
`source ~/.bash_profile` ※bashの場合
`go build`

（4）ローカルにNgrokを導入・起動する
* 導入
https://qiita.com/hirokisoccer/items/7033c1bb9c85bf6789bd

* 起動
`ngrok http 8080`

※API Gatewayでエンドポイントを指定する際、ローカル環境を外部公開する必要があり、そのためのツールとしてNgrokの使用を想定しています。今回はローカルホストのポート番号8080を使用します。

（5）CloudFormationを使ってAWS環境を構築する
* テンプレートファイルの環境変数を設定する（cloudformation.yml 412行目）
  * Ngrok…（４）で起動したNgrokのURLを設定
  * Domain…Cognito用ドメイン。任意の文字列を入力
* CloudFormationにてcloudformation.ymlをインポートしてスタックを作成
* 作成されたAPI Gatewayを確認し、デプロイ
  * デプロイ用のステージを適宜用意して下さい
* 作成されたCognitoユーザプールIDとアプリクライアントIDをメモ

（6）ローカルでEchoの開発サーバを起動
* 環境変数を定義
  * APIルートディレクトリにて`cp .env.example .env`
  * .envにてDB情報、AWS Credential、（５）でメモしたCognito情報を設定
* DBマイグレーション
  * 下記『マイグレーション方法』を参照
* 開発サーバ起動
  * APIルートディレクトリにて`go run main.go`
* 動作確認
  * curlコマンドやPostmanを使ってAPIを叩く
  * デプロイされたAPI GatewayのURLを適宜確認して下さい
  * 例）会員登録API：<API GatewayのURL>/users/signup
    * APIの仕様については『API概要』を参照

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
