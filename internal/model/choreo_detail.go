//go:generate easytags $GOFILE db
package model

import "database/sql"

type ChoreographyDetailModel struct {
	ChoreoDetailID       int64           `db:"choreo_detail_id"`
	Title                sql.NullString  `db:"title"`
	Duration             sql.NullFloat64 `db:"duration"`
	IsActive             sql.NullInt32   `db:"is_active"`
	VideoURL             sql.NullString  `db:"video_url"`
	VideoThumbnailURL    sql.NullString  `db:"video_thumbnail_url"`
	CDNVideoURL          sql.NullString  `db:"video_url_cdn"`
	CDNVideoThumbnailURL sql.NullString  `db:"video_thumbnail_url_cdn"`
	TestVideoURL         sql.NullString  `db:"test_video_url"`
	CDNTestVideoURL      sql.NullString  `db:"test_video_url_cdn"`
	ChoreoID             sql.NullInt64   `db:"choreo_id"`

	VisionTimeOffset     sql.NullFloat64 `db:"vision_time_offset"`
	VisionAngleThreshold sql.NullFloat64 `db:"vision_angle_threshold"`
	VisionBodyPose       sql.NullString  `db:"vision_body_pose"`
	// Will deprecate soon
	Position sql.NullInt32 `db:"position"`
}
