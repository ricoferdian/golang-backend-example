//go:generate easytags $GOFILE
package entity

type UserFilterEntity struct {
	UserID       int64  `json:"user_id"`
	UserIdentity string `json:"user_identity"`
	AuthType     int
}

type AuthUserResponseEntity struct {
	TokenData AuthTokenEntity `json:"token_data"`
	UserData  UserEntity      `json:"user_data"`
}

type AuthTokenEntity struct {
	AccessToken string `json:"access_token"`
	ExpiryTime  string `json:"expiry_timestamp"`
}

type LoginUserEntity struct {
	UserID             int64  `json:"user_id"`
	UserIdentity       string `json:"user_identity"`
	PasswordIdentifier string `json:"password_identifier"`
}

type RoleFilterEntity struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

type RoleEntity struct {
	RoleID   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}

type UserEntity struct {
	UserID             int64         `json:"user_id"`
	UserIdentity       string        `json:"user_identity"`
	PasslessIdentity   string        `json:"passless_identity"`
	FirstName          string        `json:"first_name"`
	LastName           string        `json:"last_name"`
	UserType           int16         `json:"user_type"`
	RoleData           *[]RoleEntity `json:"role_data"`
	CurrentRoleData    *RoleEntity   `json:"current_role_data"`
	PasswordIdentifier string
}

type AuthenticatedUserEntity struct {
	UserID          int64       `json:"user_id"`
	UserIdentity    string      `json:"user_identity"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	UserType        int16       `json:"user_type"`
	CurrentRoleData *RoleEntity `json:"current_role_data"`
}
