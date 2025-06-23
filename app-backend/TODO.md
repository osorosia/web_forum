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

## Bun ORM 設計方針

### 採用理由
- 高性能なGo向けORM（SQLクエリビルダー）
- 型安全性が高く、構造体タグによる直感的なマッピング
- PostgreSQL、MySQL、SQLiteなど主要DBをサポート
- マイグレーション機能内蔵
- トランザクション管理が容易

### プロジェクト構造
```
app-backend/
├── cmd/
│   └── server/
│       └── main.go          # アプリケーションエントリーポイント
├── internal/
│   ├── config/
│   │   └── config.go        # 環境変数管理
│   ├── database/
│   │   ├── connection.go    # DB接続管理
│   │   └── migrations/      # SQLマイグレーションファイル
│   ├── models/
│   │   ├── thread.go        # Threadモデル定義
│   │   └── post.go          # Postモデル定義
│   ├── repository/
│   │   ├── interface.go     # リポジトリインターフェース
│   │   ├── thread.go        # Thread用DB操作
│   │   └── post.go          # Post用DB操作
│   └── handlers/
│       ├── thread.go        # Thread APIハンドラ
│       └── post.go          # Post APIハンドラ
└── pkg/
    └── response/
        └── response.go      # 統一レスポンス形式
```

### モデル設計例

#### Threadモデル
```go
type Thread struct {
    bun.BaseModel `bun:"table:threads,alias:t"`
    
    ID        int64     `bun:"id,pk,autoincrement"`
    Title     string    `bun:"title,notnull"`
    CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
    
    // Relations
    Posts []Post `bun:"rel:has-many,join:id=thread_id"`
}
```

#### Postモデル
```go
type Post struct {
    bun.BaseModel `bun:"table:posts,alias:p"`
    
    ID         int64     `bun:"id,pk,autoincrement"`
    ThreadID   int64     `bun:"thread_id,notnull"`
    PostNumber int       `bun:"post_number,notnull"`
    Name       string    `bun:"name,notnull,default:'名無しさん'"`
    Content    string    `bun:"content,notnull"`
    CreatedAt  time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
    
    // Relations
    Thread *Thread `bun:"rel:belongs-to,join:thread_id=id"`
}
```

### 実装方針
1. **リポジトリパターン採用**
   - DB操作をリポジトリ層に集約
   - テスタビリティの向上
   - ビジネスロジックとDB操作の分離

2. **トランザクション管理**
   - スレッド作成時は必ずトランザクション内で処理
   - Bunの`db.RunInTx()`を活用

3. **エラーハンドリング**
   - DB操作エラーは適切にラップして返す
   - NotFoundエラーとその他のエラーを区別

4. **パフォーマンス考慮**
   - 必要に応じてインデックスを設定
   - N+1問題を避けるため、適切にPreloadを使用
   - ページネーションはLIMIT/OFFSETで実装
