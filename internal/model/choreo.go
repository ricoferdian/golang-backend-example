//go:generate easytags $GOFILE db
package model

import "database/sql"

type ChoreographyModel struct {
	ChoreoID             int64           `db:"choreo_id"`
	Title                sql.NullString  `db:"title"`
	Description          sql.NullString  `db:"description"`
	Difficulty           sql.NullInt32   `db:"difficulty"`
	Duration             sql.NullFloat64 `db:"duration"`
	IsActive             sql.NullInt32   `db:"is_active"`
	VideoPreviewURL      sql.NullString  `db:"video_preview_url"`
	VideoThumbnailURL    sql.NullString  `db:"video_thumbnail_url"`
	CDNVideoPreviewURL   sql.NullString  `db:"video_preview_url_cdn"`
	CDNVideoThumbnailURL sql.NullString  `db:"video_thumbnail_url_cdn"`
	ChoreographerID      sql.NullInt64   `db:"choreographer_id"`
	MusicID              sql.NullInt64   `db:"music_id"`
	AdditionalInfo       sql.NullString  `db:"additional_info"`
	// Will deprecate soon
	Position  sql.NullInt32 `db:"position"`
	TempPrice sql.NullInt64 `db:"temp_price"`
}
