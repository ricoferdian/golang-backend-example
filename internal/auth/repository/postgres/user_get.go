package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresUserAuthRepository) GetSingleUserByUniqueFilter(ctx context.Context, filter entity.UserFilterEntity) (result *model.RbacUserModel, err error) {
	query, args, err := c.buildGetSingleUserByFilter(filter)
	if err != nil {
		return nil, err
	}
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

func (c PostgresUserAuthRepository) buildGetSingleUserByFilter(filter entity.UserFilterEntity) (string, []interface{}, error) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectUserByIdentity)
	sb.From(tableRbacUser)

	if filter.UserID != 0 {
		sb.Where(sb.Equal("user_id", filter.UserID))
	}
	if filter.AuthType == 0 {
		return "", nil, errors.New("auth type cannot be empty")
	}
	if filter.UserIdentity == "" {
		return "", nil, errors.New("identity cannot be empty")
	}
	if filter.AuthType == constants.AuthTypeUserPassword {
		sb.Where(sb.Equal("user_identity", filter.UserIdentity))
		sb.Where(sb.IsNotNull("user_identity"))
		sb.Where(sb.IsNotNull("password_identifier"))
	}
	if filter.AuthType == constants.AuthTypePasswordlessOtp {
		sb.Where(sb.Equal("passless_identity", filter.UserIdentity))
		sb.Where(sb.IsNotNull("passless_identity"))
	}
	sb.Where(sb.Equal("is_active", UserStatusActive))

	q, qb := sb.Build()

	return q, qb, nil
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
