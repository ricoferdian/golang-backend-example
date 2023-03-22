//go:generate easytags $GOFILE
package entity

type ChoreographyEntity struct {
	ChoreoID          int64   `json:"choreo_id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Difficulty        int32   `json:"difficulty"`
	Duration          float64 `json:"duration"`
	IsActive          int32   `json:"is_active"`
	VideoPreviewURL   string  `json:"video_preview_url"`
	VideoThumbnailURL string  `json:"video_thumbnail_url"`
	ChoreographerID   int64   `json:"choreographer_id"`
	MusicID           int64   `json:"music_id"`
	// Will deprecate soon
	Order int32 `json:"order"`
	// Relation
	ChoreographerData *ChoreographerEntity `json:"choreographer_data,omitempty"`
	MusicData         *MusicEntity         `json:"music_data,omitempty"`
}

type ChoreographerEntity struct {
	ChoreographerID   int64  `json:"choreographer_id"`
	ChoreographerName string `json:"choreographer_name"`
}

type MusicEntity struct {
	MusicID    int64  `json:"music_id"`
	Title      string `json:"title"`
	ArtistName string `json:"artist_name"`
}
