//go:generate easytags $GOFILE db
package model

type MusicModel struct {
	MusicID    int64  `db:"music_id"`
	Title      string `db:"title"`
	ArtistName string `db:"artist_name"`
}
