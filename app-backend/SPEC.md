# 掲示板バックエンド仕様書

## 概要

2ちゃんねる形式の匿名掲示板のバックエンドAPI

## 基本仕様

### 機能一覧

1. スレッド一覧表示
2. スレッド作成
3. スレッド内投稿一覧表示
4. 投稿作成

### スレッド仕様

- スレッドタイトルは必須
- スレッド作成時に最初の投稿も同時に作成
- レス数上限なし
- スレッド一覧は作成日順（新しい順）で表示
- 削除機能なし

### 投稿仕様

- 名前欄あり（任意入力）
- 名前未入力時は「名無しさん」と表示
- 投稿番号は各スレッド内で1から連番
- 投稿日時を記録・表示
- 削除機能なし
- 画像アップロード機能なし
