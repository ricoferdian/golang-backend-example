package postgres

import (
	"context"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresMusicRepository) DeleteMusicByID(ctx context.Context, musicID int64) error {
	query, args := c.buildDeleteByID(musicID)
	_, err := c.dbCli.ExecContext(ctx, c.dbCli.Rebind(query), args...)
	return err
}

func (c PostgresMusicRepository) buildDeleteByID(musicID int64) (string, []interface{}) {
	sb := sq.NewDeleteBuilder()
	sb.DeleteFrom(tableMasterMusic)
	sb.Where(sb.Equal("music_id", musicID))

	return sb.Build()
}
