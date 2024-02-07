package sqlite

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/MangataL/BangumiBuddy/internal/auth"
	"github.com/MangataL/BangumiBuddy/pkg/log"
)

type userRepo struct {
	db *gorm.DB
}

var defaultUser = userScheme{
	Username: "admin",
	Password: "admin",
}

func (u *userRepo) GetUser(ctx context.Context) (auth.User, error) {
	var user userScheme
	err := u.db.WithContext(ctx).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return auth.User{
			Username: defaultUser.Username,
			Password: defaultUser.Password,
		}, u.db.Create(&defaultUser).Error
	}
	return auth.User{
		Username: user.Username,
		Password: user.Password,
	}, err
}

func (u *userRepo) ApplyUser(ctx context.Context, user auth.User) error {
	var curUser userScheme
	if err := u.db.WithContext(ctx).Take(&curUser).Error; err != nil {
		return err
	}
	curUser.Username = user.Username
	curUser.Password = user.Password
	return u.db.WithContext(ctx).Save(&curUser).Error
}

func (u *userRepo) GetUserCookie(ctx context.Context, cookie string) (auth.UserCookie, error) {
	uc := userCookieScheme{
		Cookie: cookie,
	}
	err := u.db.WithContext(ctx).Take(&uc).Error
	return auth.UserCookie{
		Username: uc.Username,
		Cookie:   uc.Cookie,
		ExpireAt: uc.ExpireAt,
	}, err
}

func (u *userRepo) SetCookie(ctx context.Context, cookie auth.UserCookie) error {
	return u.db.WithContext(ctx).Save(&userCookieScheme{
		Username: cookie.Username,
		Cookie:   cookie.Cookie,
		ExpireAt: cookie.ExpireAt,
	}).Error
}

func (u *userRepo) DeleteCookie(ctx context.Context, username string) error {
	return u.db.WithContext(ctx).Delete(&userCookieScheme{}, "username = ?", username).Error
}

func NewUserRepo(db *gorm.DB) auth.UserRepository {
	if err := db.AutoMigrate(&userScheme{}, &userCookieScheme{}); err != nil {
		log.Fatalf(context.Background(), "failed to migrate user scheme: %s", err)
	}
	return &userRepo{
		db: db,
	}
}
