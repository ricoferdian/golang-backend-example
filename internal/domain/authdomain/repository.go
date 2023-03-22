package authdomain

type UserAuthDatabaseRepo interface {
}

type UserAuthCacheRepo interface {
}

type UserAuthRepository interface {
	UserAuthDatabaseRepo
	UserAuthCacheRepo
}
