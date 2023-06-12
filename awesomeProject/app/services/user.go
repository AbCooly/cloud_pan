package services

import (
	"awesomeProject/app/common/request"
	"awesomeProject/app/models"
	"awesomeProject/global"
	"awesomeProject/utils"
	"errors"
	"fmt"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	has, err := global.App.DB.Table("user").Where("name = ?", params.Name).Exist()
	if err != nil {
		fmt.Println(err)
		return
	} else if has {
		err = errors.New("name已存在")
		return
	}
	user = models.User{Name: params.Name, Password: utils.BcryptMake([]byte(params.Password))}
	_, err = global.App.DB.Insert(&user)
	return
}

func (userService *userService) Login(params request.Login) (err error, user models.User) {
	//bug here
	has, err := global.App.DB.Table("user").Where("name = ?", params.Name).Get(&user)
	//fmt.Println(user.Name)
	if err != nil {
		err = errors.New("查询失败")
	} else if !has {
		err = errors.New("用户名不存在")
	} else if !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		fmt.Println(user.Password)
		err = errors.New("密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	has, err := global.App.DB.ID(intId).Get(&user)
	if err != nil {
		return err, models.User{}
	}
	if !has {
		err = errors.New("数据不存在")
	}
	return
}

func (userService *userService) GetUserInfoByName(name string) (err error, user models.User) {
	has, err := global.App.DB.Where("name = ?", name).Get(&user)
	if err != nil {
		return err, models.User{}
	}
	if !has {
		err = errors.New("数据不存在")
	}
	return
}
