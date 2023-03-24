package postgres

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
	"kora-backend/internal/model"
)

func (c PostgresMusicRepository) GetMusicByIdsMap(ctx context.Context, musicIDs []int64) (map[int64]model.MusicModel, error) {
	if len(musicIDs) == 0 {
		return nil, errors.New("musicIDs must be supplied")
	}
	query, args := c.buildGetMusicByIDs(musicIDs)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int64]model.MusicModel)
	for rows.Next() {
		musicData, err := c.scanMusicData(rows)
		if err != nil {
			return nil, err
		}
		result[musicData.MusicID] = musicData
	}
	return result, nil
}

func (c PostgresMusicRepository) buildGetMusicByIDs(musicIDs []int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllMusic)
	sb.From(tableMasterMusic)

	anySb := sq.Buildf("%v WHERE music_id = any(%v)", sb, pq.Array(musicIDs))

	return anySb.Build()
}

func (c PostgresMusicRepository) GetMusicById(ctx context.Context, musicID int64) (*model.MusicModel, error) {
	if musicID == 0 {
		return nil, errors.New("musicID must be supplied")
	}
	query, args := c.buildGetMusicByID(musicID)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var musicData model.MusicModel
	for rows.Next() {
		musicData, err = c.scanMusicData(rows)
		if err != nil {
			return nil, err
		}
		break
	}
	return &musicData, nil
}

func (c PostgresMusicRepository) buildGetMusicByID(musicID int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllMusic)
	sb.From(tableMasterMusic)

	sb.Where(sb.Equal("music_id", musicID))

	return sb.Build()
}

func (c PostgresMusicRepository) scanMusicData(row *sql.Rows) (result model.MusicModel, err error) {
	err = row.Scan(
		&result.MusicID,
		&result.ArtistName,
		&result.Title,
	)

	return result, err
}
