package postgres

import (
	"context"
	"database/sql"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresChoreoRepository) InsertChoreoDetail(ctx context.Context, detail model.ChoreographyDetailModel) (result model.ChoreographyDetailModel, err error) {
	query, args := c.buildInsertChoreoDetail(detail)
	queryResult, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err != nil {
		return detail, err
	}
	defer queryResult.Close()
	for queryResult.Next() {
		err = queryResult.Scan(
			&detail.ChoreoDetailID,
		)
		break
	}
	return detail, nil
}

func (c PostgresChoreoRepository) buildInsertChoreoDetail(detail model.ChoreographyDetailModel) (string, []interface{}) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto(tableMasterChoreoDetail)
	ib.Cols(columnInsertChoreoDetail)
	ib.Values(
		detail.ChoreoID,
		detail.Title,
		detail.Duration,
		detail.IsActive,
		detail.Position,
		detail.CDNVideoURL,
		detail.CDNVideoThumbnailURL,
		detail.CDNTestVideoURL,
		detail.VisionBodyPose,
		detail.VisionAngleThreshold,
		detail.VisionTimeOffset,
	)
	addQ := sq.Buildf("%v RETURNING det_choreo_id", ib)
	return addQ.Build()
}

func (c PostgresChoreoRepository) UpdateChoreoDetail(ctx context.Context, detail model.ChoreographyDetailModel) (result model.ChoreographyDetailModel, err error) {
	query, args := c.buildUpdateChoreoDetail(detail)
	_, err = c.dbCli.ExecContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return detail, nil
	}
	if err != nil {
		return detail, err
	}
	return detail, nil
}

func (c PostgresChoreoRepository) buildUpdateChoreoDetail(detail model.ChoreographyDetailModel) (string, []interface{}) {
	ub := sq.NewUpdateBuilder()
	ub.Update(tableMasterChoreoDetail)

	ub.SetMore(ub.Assign("choreo_id", detail.ChoreoID))
	ub.SetMore(ub.Assign("title", detail.Title))
	if detail.Duration.Float64 != 0 {
		ub.SetMore(ub.Assign("duration", detail.Duration.Float64))
	}
	if detail.IsActive.Int32 != 0 {
		ub.SetMore(ub.Assign("is_active", detail.IsActive.Int32))
	}
	if detail.Position.Int32 != 0 {
		ub.SetMore(ub.Assign("position", detail.Position.Int32))
	}
	if detail.CDNVideoURL.String != "" {
		ub.SetMore(ub.Assign("vid_url_cdn", detail.CDNVideoURL.String))
	}
	if detail.CDNVideoThumbnailURL.String != "" {
		ub.SetMore(ub.Assign("vid_thumbnail_url_cdn", detail.CDNVideoThumbnailURL.String))
	}
	if detail.CDNTestVideoURL.String != "" {
		ub.SetMore(ub.Assign("vid_test_url_cdn", detail.CDNTestVideoURL.String))
	}
	if detail.VisionBodyPose.String != "" {
		ub.SetMore(ub.Assign("vision_body_pose", detail.VisionBodyPose.String))
	}
	if detail.VisionAngleThreshold.Float64 != 0 {
		ub.SetMore(ub.Assign("vision_angle_threshold", detail.VisionAngleThreshold.Float64))
	}
	if detail.VisionTimeOffset.Float64 != 0 {
		ub.SetMore(ub.Assign("vision_time_offset", detail.VisionTimeOffset.Float64))
	}
	ub.Where(ub.Equal("det_choreo_id", detail.ChoreoDetailID))
	return ub.Build()
}

func (c PostgresChoreoRepository) UpdateChoreoDetailLink(ctx context.Context, detail model.ChoreographyDetailModel) (result model.ChoreographyDetailModel, err error) {
	query, args := c.buildUpdateChoreoDetailLink(detail)
	_, err = c.dbCli.ExecContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return detail, nil
	}
	if err != nil {
		return detail, err
	}
	return detail, nil
}

func (c PostgresChoreoRepository) buildUpdateChoreoDetailLink(detail model.ChoreographyDetailModel) (string, []interface{}) {
	ub := sq.NewUpdateBuilder()
	ub.Update(tableMasterChoreoDetail)
	if detail.CDNVideoURL.String != "" {
		ub.SetMore(ub.Assign("vid_url_cdn", detail.CDNVideoURL.String))
	}
	if detail.CDNVideoThumbnailURL.String != "" {
		ub.SetMore(ub.Assign("vid_thumbnail_url_cdn", detail.CDNVideoThumbnailURL.String))
	}
	if detail.CDNTestVideoURL.String != "" {
		ub.SetMore(ub.Assign("vid_test_url_cdn", detail.CDNTestVideoURL.String))
	}
	ub.Where(ub.Equal("det_choreo_id", detail.ChoreoDetailID))
	return ub.Build()
}
