package models

import (
	"strconv"
	"time"
)

type User struct {
	Id       uint      `json:"id" xorm:"pk autoincr"`
	Name     string    `json:"name" xorm:"not null comment('用户名称')"`
	Password string    `json:"password" xorm:"not null default('') comment('用户密码')"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
	Deleted  time.Time `xorm:"deleted"`
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.Id))
}
