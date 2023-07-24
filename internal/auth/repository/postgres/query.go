package postgres

const (
	tableRbacUser     = "rbac_user"
	tableRbacUserRole = "rbac_user_role"

	columnSelectUserByIdentity = "user_id,user_identity,password_identifier,first_name,last_name,user_type"
	columnSelectUserRoleByUser = "user_role_id,user_id,role_id"

	columnInsertUser = "user_identity,password_identifier,first_name,last_name,user_type,passless_identity"

	UserStatusActive   = 1
	UserStatusInactive = 0
)
