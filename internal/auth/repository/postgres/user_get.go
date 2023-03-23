package postgres

import (
	"context"
	"database/sql"
	sq "github.com/huandu/go-sqlbuilder"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

func (c PostgresUserAuthRepository) GetSingleUserByUniqueFilter(ctx context.Context, filter entity.UserFilterEntity) (result *model.RbacUserModel, err error) {
	query, args := c.buildGetSingleUserByFilter(filter)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return result, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		data, err := c.scanUserData(rows)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}
	return result, nil
}

func (c PostgresUserAuthRepository) buildGetSingleUserByFilter(filter entity.UserFilterEntity) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectUserByIdentity)
	sb.From(tableRbacUser)

	if filter.UserID != 0 {
		sb.Where(sb.Equal("user_id", filter.UserID))
	}
	if filter.UserIdentity != "" {
		sb.Where(sb.Equal("user_identity", filter.UserIdentity))
	}

	return sb.Build()
}

func (c PostgresUserAuthRepository) scanUserData(row *sql.Rows) (result model.RbacUserModel, err error) {
	err = row.Scan(
		&result.UserID,
		&result.UserIdentity,
		&result.HashPasswordIdentifier,
		&result.FirstName,
		&result.LastName,
		&result.UserType,
	)

	return result, err
}
