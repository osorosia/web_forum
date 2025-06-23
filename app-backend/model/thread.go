package models

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"
)

type Thread struct {
	bun.BaseModel `bun:"table:threads,alias:t"`

	ID        string    `bun:"id,pk"`
	Title     string    `bun:"title,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`

	// Relations
	Posts []Post `bun:"rel:has-many,join:id=thread_id"`
}

var _ bun.BeforeInsertHook = (*Thread)(nil)

func (t *Thread) BeforeInsert(ctx context.Context, query *bun.InsertQuery) error {
	if t.ID == "" {
		time.Sleep(time.Millisecond)
		id, err := ulid.New(ulid.Now(), rand.Reader)
		if err != nil {
			return err
		}
		t.ID = id.String()
	}
	return nil
}
