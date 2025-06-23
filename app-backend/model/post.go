package models

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`

	ID         string    `bun:"id,pk"`
	ThreadID   string    `bun:"thread_id,notnull"`
	PostNumber int       `bun:"post_number,notnull"`
	Name       string    `bun:"name,notnull,default:'名無しさん'"`
	Content    string    `bun:"content,notnull"`
	CreatedAt  time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`

	// Relations
	Thread *Thread `bun:"rel:belongs-to,join:thread_id=id"`
}

var _ bun.BeforeInsertHook = (*Post)(nil)

func (p *Post) BeforeInsert(ctx context.Context, query *bun.InsertQuery) error {
	if p.ID == "" {
		time.Sleep(time.Millisecond)
		id, err := ulid.New(ulid.Now(), rand.Reader)
		if err != nil {
			return err
		}
		p.ID = id.String()
	}
	return nil
}
