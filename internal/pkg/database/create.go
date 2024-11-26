package database

import (
	"context"
	"database/sql"
	"math"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/jxskiss/base62"
	"github.com/pkg/errors"
)

type Url struct {
	Id      int64  `db:"id"`
	Slug    string `db:"slug"`
	Url     string `db:"url"`
	Clicked int    `db:"clicked"`

	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}

func Create(ctx context.Context, db *Database, u *url.URL) (string, error) {
	queryUrl, err := GetUrlByUrl(ctx, db, u)
	if err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return "", err
		}
	}

	if queryUrl != nil {
		return queryUrl.Slug, nil
	}

	r := rand.Int64N(math.MaxInt64)
	encode := base62.FormatInt(r)
	expiresAt := time.Now().Add(time.Hour * 720 * 3)

	query := `INSERT INTO urls(id, slug, url, expires_at) VALUES ($1, $2, $3, $4)`

	tx := db.Db.MustBegin()

	tx.MustExecContext(ctx, query, r, string(encode), u.String(), expiresAt)

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return string(encode), nil
}
