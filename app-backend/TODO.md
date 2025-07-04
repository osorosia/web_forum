# 掲示板バックエンド実装TODO

## ORM（bun）導入

- [ ] bun ORM のセットアップ
  - [ ] 必要な依存関係の追加
    - [ ] github.com/uptrace/bun
    - [ ] github.com/uptrace/bun/driver/pgdriver (PostgreSQL)
    - [ ] github.com/uptrace/bun/dialect/pgdialect
  - [ ] データベース接続設定の実装
  - [ ] モデル定義（Bunのタグを使用）
    - [ ] Threadモデル
    - [ ] Postモデル
  - [ ] マイグレーション管理の設定
  - [ ] リポジトリパターンの実装
    - [ ] ThreadRepository
    - [ ] PostRepository

## データベース設計

- [ ] データベーススキーマ設計
  - [ ] threadsテーブル（id, title, created_at）
  - [ ] postsテーブル（id, thread_id, post_number, name, content, created_at）
- [ ] データベース接続設定
- [ ] マイグレーションファイル作成

## API実装

### スレッド関連

- [ ] GET /threads - スレッド一覧取得API
  - [ ] 作成日順（新しい順）でソート
  - [ ] ページネーション対応（オプション）
- [ ] POST /threads - スレッド作成API
  - [ ] タイトル必須バリデーション
  - [ ] 最初の投稿も同時に作成
  - [ ] トランザクション処理

### 投稿関連

- [ ] GET /threads/:id/posts - スレッド内投稿一覧取得API
  - [ ] 投稿番号順でソート
  - [ ] ページネーション対応（オプション）
- [ ] POST /threads/:id/posts - 投稿作成API
  - [ ] 名前欄の処理（空の場合「名無しさん」）
  - [ ] 投稿番号の自動採番（スレッド内連番）
  - [ ] 投稿日時の記録

## 共通処理

- [ ] エラーハンドリング
  - [ ] 統一的なエラーレスポンス形式
  - [ ] 適切なHTTPステータスコード
- [ ] リクエストバリデーション
- [ ] レスポンス形式の統一（JSON）
- [ ] ログ出力設定

## テスト

- [ ] 単体テスト
  - [ ] 各APIエンドポイントのテスト
  - [ ] バリデーションのテスト
- [ ] 統合テスト
  - [ ] スレッド作成から投稿までのフロー

## その他

- [ ] 環境変数設定（DB接続情報など）
- [ ] Dockerファイル作成（オプション）
- [ ] API仕様書作成（OpenAPI/Swagger）（オプション）
- [ ] READMEの更新
