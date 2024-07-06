package postgres

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresChoreographerRepository) UpsertChoreographerByIds(ctx context.Context, choreographerData model.ChoreographerModel) (*model.ChoreographerModel, error) {
	query, args := c.buildUpsertChoreographer(choreographerData)
	res, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		err = res.Scan(
			&choreographerData.ChoreographerID,
		)
		break
	}
	return &choreographerData, nil
}

func (c PostgresChoreographerRepository) buildUpsertChoreographer(choreographerData model.ChoreographerModel) (string, []interface{}) {
	if choreographerData.ChoreographerID != 0 {
		return c.buildUpdateChoreographer(choreographerData)
	}
	return c.buildInsertChoreographer(choreographerData)
}

func (c PostgresChoreographerRepository) buildInsertChoreographer(choreographerData model.ChoreographerModel) (string, []interface{}) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto(tableMasterChoreographer)
	ib.Cols(columnInsertChoreographer)
	ib.Values(
		choreographerData.ChoreographerName,
		choreographerData.Description,
		choreographerData.ProfileImageURL,
	)
	addQ := sq.Buildf("%v RETURNING choreographer_id", ib)

	return addQ.Build()
}

func (c PostgresChoreographerRepository) buildUpdateChoreographer(choreographerData model.ChoreographerModel) (string, []interface{}) {
	ub := sq.NewUpdateBuilder()
	ub.Update(tableMasterChoreographer)
	if choreographerData.ChoreographerName != "" {
		ub.SetMore(ub.Assign("name", choreographerData.ChoreographerName))
	}
	if choreographerData.Description.String != "" {
		ub.SetMore(ub.Assign("description", choreographerData.Description))
	}
	if choreographerData.ProfileImageURL.String != "" {
		ub.SetMore(ub.Assign("profile_image_url", choreographerData.ProfileImageURL))
	}
	ub.Where(ub.Equal("choreographer_id", choreographerData.ChoreographerID))

	return ub.Build()
}
