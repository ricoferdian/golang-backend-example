//go:generate easytags $GOFILE db
package model

type RbacRoleModel struct {
	RoleID   int64  `db:"role_id"`
	RoleName string `db:"role_name"`
}
