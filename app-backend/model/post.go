package models

import (
	"time"

	"github.com/uptrace/bun"
)

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
