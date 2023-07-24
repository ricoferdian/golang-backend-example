package postgres

import (
	"context"
	"database/sql"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresChoreoRepository) InsertChoreo(ctx context.Context, choreo model.ChoreographyModel) (result model.ChoreographyModel, err error) {
	query, args := c.buildInsertChoreo(choreo)
	queryResult, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err != nil {
		return choreo, err
	}
	defer queryResult.Close()
	for queryResult.Next() {
		err = queryResult.Scan(
			&choreo.ChoreoID,
		)
		break
	}
	return choreo, nil
}

func (c PostgresChoreoRepository) buildInsertChoreo(choreo model.ChoreographyModel) (string, []interface{}) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto(tableMasterChoreo)
	ib.Cols(columnInsertChoreo)
	ib.Values(
		choreo.Title,
		choreo.Description,
		choreo.Difficulty,
		choreo.Duration,
		choreo.IsActive,
		choreo.Position,
		choreo.CDNVideoPreviewURL,
		choreo.CDNVideoThumbnailURL,
		choreo.ChoreographerID,
		choreo.MusicID,
		choreo.AdditionalInfo,
		choreo.TempPrice,
	)
	addQ := sq.Buildf("%v RETURNING choreo_id", ib)
	return addQ.Build()
}

func (c PostgresChoreoRepository) UpdateChoreo(ctx context.Context, choreo model.ChoreographyModel) (result model.ChoreographyModel, err error) {
	query, args := c.buildUpdateChoreo(choreo)
	_, err = c.dbCli.ExecContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return choreo, nil
	}
	if err != nil {
		return choreo, err
	}
	return choreo, nil
}

func (c PostgresChoreoRepository) buildUpdateChoreo(choreo model.ChoreographyModel) (string, []interface{}) {
	ub := sq.NewUpdateBuilder()
	ub.Update(tableMasterChoreo)

	ub.SetMore(ub.Assign("title", choreo.Title.String))
	if choreo.Description.String != "" {
		ub.SetMore(ub.Assign("description", choreo.Description.String))
	}
	if choreo.Difficulty.Int32 != 0 {
		ub.SetMore(ub.Assign("difficulty", choreo.Difficulty.Int32))
	}
	if choreo.Duration.Float64 != 0 {
		ub.SetMore(ub.Assign("duration", choreo.Duration.Float64))
	}
	ub.SetMore(ub.Assign("is_active", choreo.IsActive.Int32))
	if choreo.Position.Int32 != 0 {
		ub.SetMore(ub.Assign("position", choreo.Position.Int32))
	}
	if choreo.CDNVideoPreviewURL.String != "" {
		ub.SetMore(ub.Assign("vid_preview_url_cdn", choreo.CDNVideoPreviewURL.String))
	}
	if choreo.CDNVideoThumbnailURL.String != "" {
		ub.SetMore(ub.Assign("vid_thumbnail_url_cdn", choreo.CDNVideoThumbnailURL.String))
	}
	if choreo.ChoreographerID.Int64 != 0 {
		ub.SetMore(ub.Assign("choreographer_id", choreo.ChoreographerID.Int64))
	}
	if choreo.MusicID.Int64 != 0 {
		ub.SetMore(ub.Assign("music_id", choreo.MusicID.Int64))
	}
	if choreo.AdditionalInfo.String != "" {
		ub.SetMore(ub.Assign("additional_info", choreo.AdditionalInfo.String))
	}
	if choreo.TempPrice.Int64 != 0 {
		ub.SetMore(ub.Assign("temp_price", choreo.TempPrice.Int64))
	}
	ub.Where(ub.Equal("choreo_id", choreo.ChoreoID))
	return ub.Build()
}

func (c PostgresChoreoRepository) UpdateChoreoLink(ctx context.Context, choreo model.ChoreographyModel) (result model.ChoreographyModel, err error) {
	query, args := c.buildUpdateChoreoLink(choreo)
	_, err = c.dbCli.ExecContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return choreo, nil
	}
	if err != nil {
		return choreo, err
	}
	return choreo, nil
}

func (c PostgresChoreoRepository) buildUpdateChoreoLink(choreo model.ChoreographyModel) (string, []interface{}) {
	ub := sq.NewUpdateBuilder()
	ub.Update(tableMasterChoreo)
	if choreo.CDNVideoPreviewURL.String != "" {
		ub.SetMore(ub.Assign("vid_preview_url_cdn", choreo.CDNVideoPreviewURL.String))
	}
	if choreo.CDNVideoThumbnailURL.String != "" {
		ub.SetMore(ub.Assign("vid_thumbnail_url_cdn", choreo.CDNVideoThumbnailURL.String))
	}
	ub.Where(ub.Equal("choreo_id", choreo.ChoreoID))
	return ub.Build()
}
