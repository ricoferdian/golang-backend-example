//go:generate easytags $GOFILE
package entity

type ChoreoDetailFilterEntity struct {
	ChoreoDetailID int64 `json:"choreo_detail_id"`
	ChoreoID       int64 `json:"choreo_id"`
}

type ChoreographyDetailEntity struct {
	ChoreoDetailID    int64               `json:"choreo_detail_id"`
	Title             string              `json:"title"`
	Duration          float64             `json:"duration"`
	IsActive          int32               `json:"is_active"`
	VideoURL          string              `json:"video_url"`
	VideoThumbnailURL string              `json:"video_thumbnail_url"`
	ChoreoID          int64               `json:"choreo_id"`
	ChoreoData        *ChoreographyEntity `json:"choreo_data,omitempty"`

	VisionTimeOffset     float64 `json:"vision_time_offset"`
	VisionAngleThreshold float64 `json:"vision_angle_threshold"`
	VisionBodyPose       string  `json:"vision_body_pose"`
	// Will deprecate soon
	Order int32 `json:"order"`
}
