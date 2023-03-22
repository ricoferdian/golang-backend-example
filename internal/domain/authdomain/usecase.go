package authdomain

import "context"

type UserAuthUseCase interface {
	Login(ctx context.Context)
	Register(ctx context.Context)
}
