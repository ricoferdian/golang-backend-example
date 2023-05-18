//go:generate easytags $GOFILE json
package entity

type LearningHistoryEntity struct {
	LearningHistoryID int64   `json:"learning_history_id"`
	ChoreoDetailID    int64   `json:"choreo_detail_id"`
	UserID            int64   `json:"user_id"`
	ChoreoDetail      string  `json:"choreo_detail"`
	Device            string  `json:"device"`
	Downloaded        int32   `json:"downloaded"`
	Expired           int32   `json:"expired"`
	Progress          float64 `json:"progress"`
	RecordUrl         string  `json:"record_url"`
	ThumbnailUrl      string  `json:"thumbnail_url"`
}
