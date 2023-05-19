//go:generate easytags $GOFILE json
package entity

type SubmitLearningHistoryEntity struct {
	LearningHistoryID int64   `json:"learning_history_id" form:"learning_history_id"`
	ChoreoDetailID    int64   `json:"choreo_detail_id" validate:"required" form:"choreo_detail_id"`
	UserID            int64   `json:"user_id"`
	ChoreoDetail      string  `json:"choreo_detail" form:"choreo_detail"`
	Device            string  `json:"device" form:"device"`
	Progress          float64 `json:"progress" form:"progress"`
	RecordUrl         string  `json:"record_url" form:"record_url"`
	ThumbnailUrl      string  `json:"thumbnail_url" form:"thumbnail_url"`
}
