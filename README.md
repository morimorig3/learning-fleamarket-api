# フリーマーケットAPI

[Udemy](https://www.udemy.com/course/gin-golang/learn/lecture/42461892)のgin学習コースで作成したAPIを拡張する

## 技術スタック

golang, gin, gorm, postgresql

## 実装すること

- 同時アクセスによるデータ不整合が起こらないよう排他制御をする
- 特定のドメインからしかアクセスされないようCORS設定を行う

### バッチ処理

できれば

- 半年間ログインされていないユーザーを削除する
  - 紐づく商品も削除される

## エンドポイント

|メソッド|エンドポイント|説明|TODO|
|:---:|:---|:---|:---|
|GET|`/items`|全商品一覧を取得する|ページング|
|POST|`/items`|商品を登録する||
|GET|`/items/:id`|商品を取得する||
|PUT|`/items/:id`|商品を更新する||
|DELETE|`/items/:id`|商品を削除する||
|GET|`/auth/user`|ユーザー情報を取得する|未実装|
|POST|`/auth/user`|ユーザーを登録する|`/auth/signin`で実装済み|
|DELETE|`/auth/user`|ユーザーを削除する|未実装|
|POST|`/auth/login`|ログインする||
|POST|`/auth/logout`|ログアウトする|未実装|
