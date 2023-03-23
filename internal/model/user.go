//go:generate easytags $GOFILE db
package model

import "database/sql"

type RbacUserModel struct {
	UserID                 int64          `db:"user_id"`
	UserIdentity           string         `db:"user_identity"`
	HashPasswordIdentifier string         `db:"password_identifier"`
	FirstName              sql.NullString `db:"first_name"`
	LastName               sql.NullString `db:"last_name"`
	UserType               sql.NullInt16  `db:"user_type"`
}
