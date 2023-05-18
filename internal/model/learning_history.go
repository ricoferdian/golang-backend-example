//go:generate easytags $GOFILE db
package model

import "database/sql"

type LearningHistoryFilter struct {
	LearningHistoryID int64 `db:"learning_history_id"`
	ChoreoDetailID    int64 `db:"choreo_detail_id"`
	UserID            int64 `db:"user_id"`
}

type LearningHistoryModel struct {
	LearningHistoryID int64           `db:"learning_history_id"`
	ChoreoDetailID    int64           `db:"choreo_detail_id"`
	UserID            int64           `db:"user_id"`
	ChoreoDetail      sql.NullString  `db:"choreo_detail"`
	Device            sql.NullString  `db:"device"`
	Downloaded        sql.NullInt32   `db:"downloaded"`
	Expired           sql.NullInt32   `db:"expired"`
	Progress          sql.NullFloat64 `db:"progress"`
	RecordUrl         sql.NullString  `db:"record_url"`
	ThumbnailUrl      sql.NullString  `db:"thumbnail_url"`
}

type SubmitLearningHistoryModel struct {
	LearningHistoryID int64           `db:"learning_history_id"`
	ChoreoDetailID    int64           `db:"choreo_detail_id"`
	UserID            int64           `db:"user_id"`
	ChoreoDetail      sql.NullString  `db:"choreo_detail"`
	Device            sql.NullString  `db:"device"`
	Progress          sql.NullFloat64 `db:"progress"`
	RecordUrl         sql.NullString  `db:"record_url"`
	ThumbnailUrl      sql.NullString  `db:"thumbnail_url"`
}
