# フリーマーケットAPI

[Udemy](https://www.udemy.com/course/gin-golang/learn/lecture/42461892)のgin学習コースで作成したAPIを拡張する

## 技術スタック

golang, gin, gorm, postgresql

## 実装すること

- テスト作成
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
|POST|`/auth/signin`|ユーザーを登録する||
|POST|`/auth/login`|ログインする||
|GET|`/orders`|注文履歴を取得する||
|POST|`/orders/:id`|注文する||

## トランザクション管理が必要な処理

注文する処理