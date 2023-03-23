//go:generate easytags $GOFILE db
package model

type RbacUserRoleModel struct {
	UserRoleID int64 `db:"user_role_id"`
	RoleID     int64 `db:"role_id"`
	UserID     int64 `db:"user_id"`
}
