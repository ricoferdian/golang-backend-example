package postgres

import (
	"context"
	"database/sql"
	sq "github.com/huandu/go-sqlbuilder"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

func (c PostgresChoreoRepository) GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) (result []model.ChoreographyDetailModel, err error) {
	query, args := c.buildGetChoreoDetailByChoreoID(filter)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return result, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		choreoDetData, err := c.scanChoreoDetailData(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, choreoDetData)
	}
	return result, nil
}

func (c PostgresChoreoRepository) buildGetChoreoDetailByChoreoID(filter entity.ChoreoDetailFilterEntity) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllChoreoDetail)
	sb.From(tableMasterChoreoDetail)

	if filter.ChoreoID != 0 {
		sb.Where(sb.Equal("choreo_id", filter.ChoreoID))
	}
	if filter.ChoreoDetailID != 0 {
		sb.Where(sb.Equal("choreo_detail_id", filter.ChoreoDetailID))
	}

	return sb.Build()
}

func (c PostgresChoreoRepository) scanChoreoDetailData(row *sql.Rows) (result model.ChoreographyDetailModel, err error) {
	err = row.Scan(
		&result.ChoreoDetailID,
		&result.ChoreoID,
		&result.Title,
		&result.Duration,
		&result.IsActive,
		&result.Position,
		&result.VideoURL,
		&result.VideoThumbnailURL,
		&result.VisionBodyPose,
		&result.VisionAngleThreshold,
		&result.VisionTimeOffset,
	)

	return result, err
}
