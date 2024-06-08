package service

import (
	"basic-go/webook/internel/domain"
	"basic-go/webook/internel/repository"
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Signup(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email string, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 检查密码对不对
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserService) UpdateUserInfo(ctx context.Context, userId int64, nickname string, birthday string, aboutMe string) error {
	return svc.repo.UpdateById(ctx, userId, nickname, birthday, aboutMe)

}

func (svc *UserService) UserProfile(ctx context.Context, userId int64) (string, error) {
	user, err := svc.repo.FindById(ctx, userId)
	if err != nil {
		return "", err
	}
	userinfo, _ := json.Marshal(user)
	return string(userinfo), nil
}
