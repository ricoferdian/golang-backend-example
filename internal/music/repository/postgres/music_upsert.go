package postgres

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresMusicRepository) UpsertMusic(ctx context.Context, music model.MusicModel) (*model.MusicModel, error) {
	query, args := c.buildUpsertMusic(music)
	res, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		err = res.Scan(
			&music.MusicID,
		)
		break
	}
	return &music, nil
}

func (c PostgresMusicRepository) buildUpsertMusic(music model.MusicModel) (string, []interface{}) {
	if music.MusicID != 0 {
		return c.buildUpdateMusic(music)
	}
	return c.buildInsertMusic(music)
}

func (c PostgresMusicRepository) buildInsertMusic(music model.MusicModel) (string, []interface{}) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto(tableMasterMusic)
	ib.Cols(columnInsertMusic)
	ib.Values(
		music.Title,
		music.ArtistName,
	)
	addQ := sq.Buildf("%v RETURNING music_id", ib)

	return addQ.Build()
}

func (c PostgresMusicRepository) buildUpdateMusic(music model.MusicModel) (string, []interface{}) {
	ub := sq.NewUpdateBuilder()
	ub.Update(tableMasterMusic)
	if music.Title != "" {
		ub.SetMore(ub.Assign("title", music.Title))
	}
	if music.ArtistName != "" {
		ub.SetMore(ub.Assign("artist_name", music.ArtistName))
	}
	ub.Where(ub.Equal("music_id", music.MusicID))

	return ub.Build()
}
