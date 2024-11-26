package database

import (
	"context"
	"database/sql"
	"time"
)

func ExtendSlug(ctx context.Context, db *Database, slug string) (time.Time, error) {
	date := time.Now().Add(time.Hour * 720 * 3)
	query := `
		UPDATE urls
		SET expires_at = $2
		WHERE slug = $1
  `

	tx := db.Db.MustBeginTx(ctx, nil)

	res := tx.MustExec(query, slug, date)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return time.Time{}, err
	}

	if rowsAffected == 0 {
		return time.Time{}, sql.ErrNoRows
	}

	if err := tx.Commit(); err != nil {
		return time.Time{}, err
	}

	return date, nil
}
