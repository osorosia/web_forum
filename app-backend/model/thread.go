package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Thread struct {
	bun.BaseModel `bun:"table:threads,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Title     string    `bun:"title,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`

	// Relations
	Posts []Post `bun:"rel:has-many,join:id=thread_id"`
}
