package postgres

import (
	"context"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresChoreoRepository) DeleteChoreoByID(ctx context.Context, choreoID int64) error {
	query, args := c.buildDeleteChoreoByID(choreoID)
	_, err := c.dbCli.Exec(c.dbCli.Rebind(query), args...)
	return err
}

func (c PostgresChoreoRepository) buildDeleteChoreoByID(choreoID int64) (string, []interface{}) {
	sb := sq.NewDeleteBuilder()
	sb.DeleteFrom(tableMasterChoreo)
	sb.Where(sb.Equal("choreo_id", choreoID))

	return sb.Build()
}

func (c PostgresChoreoRepository) DeleteChoreoDetailByID(ctx context.Context, choreoID int64) error {
	query, args := c.buildDeleteChoreoDetailByID(choreoID)
	_, err := c.dbCli.Exec(c.dbCli.Rebind(query), args...)
	return err
}

func (c PostgresChoreoRepository) buildDeleteChoreoDetailByID(choreoID int64) (string, []interface{}) {
	sb := sq.NewDeleteBuilder()
	sb.DeleteFrom(tableMasterChoreoDetail)
	sb.Where(sb.Equal("det_choreo_id", choreoID))

	return sb.Build()
}
