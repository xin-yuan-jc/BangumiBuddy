package sqlite

import (
	"time"
)

type userScheme struct {
	ID       int64  `gorm:"column:id;AUTO_INCREMENT;NOT NULL"`
	Username string `gorm:"column:username;NOT NULL"`
	Password string `gorm:"column:password;NOT NULL"`
}

func (userScheme) TableName() string {
	return "user"
}

type userCookieScheme struct {
	ID       int64     `gorm:"column:id;primaryKey;type:int;AUTO_INCREMENT;NOT NULL"`
	Username string    `gorm:"column:username;type:varchar(32);NOT NULL"`
	Cookie   string    `gorm:"column:cookie;type:varchar(32);NOT NULL"`
	ExpireAt time.Time `gorm:"column:expire_at;type:datetime;NOT NULL"`
}

func (userCookieScheme) TableName() string {
	return "user_cookie"
}
