package repository

import (
	"basic-go/webook/internel/domain"
	"basic-go/webook/internel/repository/dao"
	"context"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (repo *UserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *UserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (repo *UserRepository) UpdateById(ctx context.Context, id int64, nickname string, birthday string, aboutMe string) error {
	u, err := repo.dao.FindById(ctx, id)
	if err != nil {
		return err
	}
	if len(nickname) == 0 {
		nickname = u.Nickname
	}
	if len(birthday) == 0 {
		birthday = u.Birthday
	}
	if len(aboutMe) == 0 {
		aboutMe = u.AboutMe
		if len(aboutMe) == 0 {
			aboutMe = "什么也没写"
		}
	}
	return repo.dao.UpdateById(ctx, id, nickname, birthday, aboutMe)
}

func (repo *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	u, err := repo.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		AboutMe:  u.AboutMe,
	}, nil
}
