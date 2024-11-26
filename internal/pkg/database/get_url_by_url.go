package database

import (
	"context"
	"net/url"
)

func GetUrlByUrl(ctx context.Context, db *Database, u *url.URL) (*Url, error) {
	query := `
		SELECT
			id,
			slug,
			url,
			clicked,
			created_at,
			expires_at
		FROM urls
		WHERE url = $1
			AND expires_at > now()
		LIMIT 1`

	var queryUrl Url
	if err := db.Db.GetContext(ctx, &queryUrl, query, u.String()); err != nil {
		return nil, err
	}

	return &queryUrl, nil
}
