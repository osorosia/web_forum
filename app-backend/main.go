package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

// Todoモデル定義
type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Title     string    `bun:"title,notnull"`
	Done      bool      `bun:"done,notnull,default:false"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
}

func main() {
	// SQLiteデータベース接続
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:test.db?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	// bunインスタンスの作成
	db := bun.NewDB(sqldb, sqlitedialect.New())
	defer db.Close()

	ctx := context.Background()

	// テーブルの作成
	fmt.Println("=== テーブル作成 ===")
	_, err = db.NewCreateTable().Model((*Todo)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("todosテーブルを作成しました")

	// レコードの挿入
	fmt.Println("\n=== レコード挿入 ===")
	todos := []Todo{
		{Title: "Goの勉強", Done: false, CreatedAt: time.Now()},
		{Title: "bunの使い方を学ぶ", Done: true, CreatedAt: time.Now()},
		{Title: "Webアプリを作る", Done: false, CreatedAt: time.Now()},
	}

	_, err = db.NewInsert().Model(&todos).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d件のTodoを挿入しました\n", len(todos))

	// 全レコードの取得
	fmt.Println("\n=== 全レコード取得 ===")
	var allTodos []Todo
	err = db.NewSelect().Model(&allTodos).Order("id ASC").Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, todo := range allTodos {
		status := "未完了"
		if todo.Done {
			status = "完了"
		}
		fmt.Printf("ID: %d, タイトル: %s, ステータス: %s, 作成日時: %s\n",
			todo.ID, todo.Title, status, todo.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// 条件付き取得（未完了のタスクのみ）
	fmt.Println("\n=== 未完了のタスク取得 ===")
	var undoneTodos []Todo
	err = db.NewSelect().
		Model(&undoneTodos).
		Where("done = ?", false).
		Order("id ASC").
		Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, todo := range undoneTodos {
		fmt.Printf("ID: %d, タイトル: %s\n", todo.ID, todo.Title)
	}

	// 単一レコードの取得
	fmt.Println("\n=== 単一レコード取得（ID=1） ===")
	todo := new(Todo)
	err = db.NewSelect().Model(todo).Where("id = ?", 1).Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %d, タイトル: %s, ステータス: %v\n", todo.ID, todo.Title, todo.Done)
}
