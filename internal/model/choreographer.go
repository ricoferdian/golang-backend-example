//go:generate easytags $GOFILE db
package model

import "database/sql"

type ChoreographerModel struct {
	ChoreographerID   int64          `db:"choreographer_id"`
	ChoreographerName string         `db:"choreographer_name"`
	Description       sql.NullString `db:"description"`
	ProfileImageURL   sql.NullString `db:"profile_image_url"`
}
