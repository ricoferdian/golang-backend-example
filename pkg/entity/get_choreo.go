//go:generate easytags $GOFILE
package entity

type ChoreographyEntity struct {
	ChoreoID             int64   `json:"choreo_id"`
	Title                string  `json:"title"`
	Description          string  `json:"description"`
	Difficulty           int32   `json:"difficulty"`
	Duration             float64 `json:"duration"`
	IsActive             int32   `json:"is_active"`
	VideoPreviewURL      string  `json:"video_preview_url"`
	VideoThumbnailURL    string  `json:"video_thumbnail_url"`
	CDNVideoPreviewURL   string  `json:"video_preview_url_cdn"`
	CDNVideoThumbnailURL string  `json:"video_thumbnail_url_cdn"`
	ChoreographerID      int64   `json:"choreographer_id"`
	MusicID              int64   `json:"music_id"`
	LikeCount            int64   `json:"like_count"`
	Like                 bool    `json:"liked"`
	Save                 bool    `json:"saved"`
	AdditionalInfo       string  `json:"additional_info"`
	// Related to payment
	Unlocked     bool   `json:"unlocked"`
	CurrentPrice string `json:"current_price"`
	// Will deprecate soon
	Order int32 `json:"order"`
	// Relation
	ChoreographerData *ChoreographerEntity `json:"choreographer_data,omitempty"`
	MusicData         *MusicEntity         `json:"music_data,omitempty"`
}

type ChoreographerEntity struct {
	ChoreographerID   int64  `json:"choreographer_id"`
	ChoreographerName string `json:"choreographer_name"`
	Description       string `json:"description"`
	ProfileImageURL   string `json:"profile_image_url"`
}

type MusicEntity struct {
	MusicID    int64  `json:"music_id"`
	Title      string `json:"title"`
	ArtistName string `json:"artist_name"`
}
