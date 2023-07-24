package postgres

import (
	"context"
	"database/sql"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresUserAuthRepository) DeactivateUser(ctx context.Context, userID int64) error {
	query, args := c.buildDeactivateUser(userID)
	_, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (c PostgresUserAuthRepository) buildDeactivateUser(userID int64) (string, []interface{}) {
	ib := sq.NewUpdateBuilder()
	ib.Update(tableRbacUser)
	ib.Set(ib.Assign("is_active", UserStatusInactive))
	ib.Where(ib.Equal("user_id", userID))

	return ib.Build()
}

func (c PostgresUserAuthRepository) ReactivateUser(ctx context.Context, userID int64) error {
	query, args := c.buildDeactivateUser(userID)
	_, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (c PostgresUserAuthRepository) buildReactivateUser(userID int64) (string, []interface{}) {
	ib := sq.NewUpdateBuilder()
	ib.Update(tableRbacUser)
	ib.Set(ib.Assign("is_active", UserStatusActive))
	ib.Where(ib.Equal("user_id", userID))

	return ib.Build()
}
