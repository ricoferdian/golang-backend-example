package postgres

import (
	"context"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresChoreographerRepository) DeleteChoreographerByID(ctx context.Context, choreographID int64) error {
	query, args := c.buildDeleteByID(choreographID)
	_, err := c.dbCli.ExecContext(ctx, c.dbCli.Rebind(query), args...)
	return err
}

func (c PostgresChoreographerRepository) buildDeleteByID(choreographID int64) (string, []interface{}) {
	sb := sq.NewDeleteBuilder()
	sb.DeleteFrom(tableMasterChoreographer)
	sb.Where(sb.Equal("choreographer_id", choreographID))

	return sb.Build()
}
