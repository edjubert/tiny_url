package database

import (
	"context"
)

func GetUrlBySlug(ctx context.Context, db *Database, slug string) (*Url, error) {
	query := `
		SELECT
			id,
			slug,
			url,
			clicked,
			created_at,
			expires_at
		FROM urls
		WHERE slug = $1
			AND expires_at > now()
		LIMIT 1`

	var queryUrl Url
	if err := db.Db.GetContext(ctx, &queryUrl, query, slug); err != nil {
		return nil, err
	}

	return &queryUrl, nil
}
