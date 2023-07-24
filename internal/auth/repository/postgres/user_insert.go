package postgres

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/auth/helper"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresUserAuthRepository) InsertSingleUser(ctx context.Context, user entity.UserEntity) (*model.RbacUserModel, error) {
	query, args := c.buildInsertSingleUser(user)
	userId, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err != nil {
		return nil, err
	}
	defer userId.Close()
	for userId.Next() {
		err = userId.Scan(
			&user.UserID,
		)
		break
	}
	result := helper.UserEntityToModel(user)

	return &result, nil
}

func (c PostgresUserAuthRepository) buildInsertSingleUser(user entity.UserEntity) (string, []interface{}) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto(tableRbacUser)
	ib.Cols(columnInsertUser)
	ib.Values(
		user.UserIdentity,
		user.PasswordIdentifier,
		user.FirstName,
		user.LastName,
		user.UserType,
		user.PasslessIdentity,
	)
	addQ := sq.Buildf("%v RETURNING user_id", ib)

	return addQ.Build()
}
