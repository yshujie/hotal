package service

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hotel/app/entity"
	"github.com/hotel/app/libs"
	"strconv"
	"strings"
	"time"
)

type AuthService struct {
	loginUser *entity.User
}

func NewAuth() *AuthService {
	return new(AuthService)
}

func (s *AuthService) Init(token string) {
	beego.Trace("登录验证，token：", token)
	arr := strings.Split(token, "|")

	if len(arr) == 2 {
		idStr, password := arr[0], arr[2]
		userId, _ := strconv.ParseInt(idStr, 10, 64)

		if userId > 0 {
			user, err := UserService.GetUser(userId)

			if err == nil && password == libs.Md5([]byte(user.Password+user.Salt)) {
				s.loginUser = user
				beego.Trace("验证成功，用户信息：", user)
			}
		}
	}
}

// GetUser 获取当前的登录用户
func (s *AuthService) GetUser() *entity.User {
	return s.loginUser
}

// GetUserId 获取当前登录用户的用户ID
func (s *AuthService) GetUserId() int64 {
	if s.IsLogined() {
		return s.loginUser.Id
	}

	return 0
}

// GetUserName 获取当前登录用户的用户名
func (s *AuthService) GetUserName() string {
	if s.IsLogined() {
		return s.loginUser.UserName
	}

	return ""
}

// GetRealName 获取当前用户的真实姓名
func (s *AuthService) GetRealName() string {
	if s.IsLogined() {
		return s.loginUser.RealName
	}

	return ""
}

// IsLogined 是否登录
func (s *AuthService) IsLogined() bool {
	return s.loginUser != nil && s.loginUser.Id > 0
}

// Login 用户登录
func (s *AuthService) Login(userName, password string) (string, error) {
	user, err := UserService.GetUserByName(userName)
	if err != nil {
		if err == orm.ErrNoRows {
			return "", errors.New("账户或密码错误")
		} else {
			return "", errors.New("系统错误")
		}
	}

	if password != libs.Md5([]byte(user.Password+user.Salt)) {
		return "", errors.New("账户或密码错误")
	}
	if user.Status == -1 {
		return "", errors.New("账户已被禁用")
	}

	user.LastLogin = time.Now()
	UserService.UpdateUser(user, "LastLogin")
	s.loginUser = user

	token := fmt.Sprintf("%d|%s", user.Id, libs.Md5([]byte(user.Password+user.Salt)))

	return token, nil
}

// Logout 退出登录
func (s *AuthService) Logout() error {
	s.loginUser = nil
	return nil
}
