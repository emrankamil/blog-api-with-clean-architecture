package usecase

import (
	"blog-api_with-clean-architecture/domain"
	"blog-api_with-clean-architecture/utils"
	"context"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return utils.GenerateAccessToken(user, expiry, secret)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.GenerateRefreshToken(user, expiry, secret)
}

func (lu *loginUsecase) LogoutUser(c context.Context, email string) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	user, err := lu.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	updatedUser := &domain.User{
		ID: user.ID,
		Token: "",
		Refresh_token: "",
	}
	
	return lu.userRepository.UpdateUser(ctx, updatedUser)
}
