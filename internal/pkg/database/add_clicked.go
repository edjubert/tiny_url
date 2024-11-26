package database

import (
	"context"
)

func AddClicked(ctx context.Context, db *Database, slug string) error {
	query := `
		UPDATE urls
		SET clicked = clicked + 1
		WHERE slug = $1
			AND expires_at > now();`

	tx := db.Db.MustBeginTx(ctx, nil)

	tx.MustExec(query, slug)

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
