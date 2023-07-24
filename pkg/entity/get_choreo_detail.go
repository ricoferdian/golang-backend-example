//go:generate easytags $GOFILE
package entity

type ChoreoDetailFilterEntity struct {
	ChoreoDetailID int64 `json:"choreo_detail_id"`
	ChoreoID       int64 `json:"choreo_id"`
}

type ChoreoFilterEntity struct {
	ChoreoID        int64 `json:"choreo_id"`
	Difficulty      int   `json:"difficulty"`
	Price           int64 `json:"price"`
	ChoreographerID int64 `json:"choreographer_id"`
}

type ChoreographyDetailEntity struct {
	ChoreoDetailID       int64               `json:"choreo_detail_id"`
	Title                string              `json:"title"`
	Duration             float64             `json:"duration"`
	IsActive             int32               `json:"is_active"`
	VideoURL             string              `json:"video_url"`
	VideoThumbnailURL    string              `json:"video_thumbnail_url"`
	CDNVideoURL          string              `json:"video_url_cdn"`
	CDNVideoThumbnailURL string              `json:"video_thumbnail_url_cdn"`
	TestVideoURL         string              `json:"test_video_url"`
	CDNTestVideoURL      string              `json:"test_video_url_cdn"`
	ChoreoID             int64               `json:"choreo_id"`
	ChoreoData           *ChoreographyEntity `json:"choreo_data,omitempty"`

	VisionTimeOffset     float64 `json:"vision_time_offset"`
	VisionAngleThreshold float64 `json:"vision_angle_threshold"`
	VisionBodyPose       string  `json:"vision_body_pose"`
	// Will deprecate soon
	Order int32 `json:"order"`
}
