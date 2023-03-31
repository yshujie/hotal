package controlers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/hotel/app/service"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"time"
)

type MainController struct {
	BaseController
}

// Index 首页
func (c *MainController) Index() {
	c.Data["pageTitle"] = "系统概况"
	c.Data["hostname"], _ = os.Hostname()
	c.Data["os"] = runtime.GOOS
	c.Data["goroutineNum"] = runtime.NumGoroutine()
	c.Data["cpuNum"] = runtime.NumCPU()
	c.Data["arch"] = runtime.GOARCH
	c.Data["dbVersion"] = service.DBVersion()
	c.Data["dataDir"] = beego.AppConfig.String("data_dir")
	up, day, hour, min, sec := c.getUptime()
	c.Data["uptime"] = fmt.Sprintf("%s，已运行 %d天 %d小时 %d分钟 %d秒", beego.Date(up, "Y-m-d H:i:s"), day, hour, min, sec)

	c.display()
}

// 获取开启时间
func (c *MainController) getUptime() (up time.Time, day, hour, min, sec int) {
	ts, _ := beego.AppConfig.Int64("up_time")
	up = time.Unix(ts, 0)
	uptime := int(time.Now().Sub(up) / time.Second)
	if uptime >= 86400 {
		day = uptime / 86400
		uptime %= 86400
	}
	if uptime >= 3600 {
		hour = uptime / 3600
		uptime %= 60
	}

	sec = uptime
	return
}

// Profile 个人信息
func (c *MainController) Profile() {
	beego.ReadFromRequest(&c.Controller)
	user := c.auth.GetUser()

	if c.isPost() {
		flash := beego.NewFlash()
		email := c.GetString("email")
		sex, _ := c.GetInt("sex")
		password1 := c.GetString("password1")
		password2 := c.GetString("password2")

		user.Email = email
		user.Sex = sex
		service.UserService.UpdateUser(user, "Email", "Sex")

		if password1 != "" {
			if len(password1) < 6 {
				flash.Error("密码长度必须大于6位")
				flash.Store(&c.Controller)
				c.redirect(beego.URLFor(".Profile"))
			} else if password2 != password1 {
				flash.Error("两次密码输入不一致")
				flash.Store(&c.Controller)
				c.redirect(beego.URLFor(".Profile"))
			} else {
				service.UserService.ModifyPassword(c.userId, password1)
			}
		}

		flash.Success("修改成功")
		flash.Store(&c.Controller)
		c.redirect(beego.URLFor(".Profile"))
	}

	c.Data["pageTitle"] = "个人信息"
	c.Data["user"] = user
	c.display()
}

// Login 登录
func (c *MainController) Login() {
	if c.userId > 0 {
		c.redirect("/")
	}

	beego.ReadFromRequest(&c.Controller)
	if c.isPost() {
		flash := beego.NewFlash()
		username := c.GetString("username")
		password := c.GetString("password")
		remember := c.GetString("remember")
		if username != "" && password != "" {
			token, err := c.auth.Login(username, password)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
				c.redirect("/login")
			} else {
				if remember == "yes" {
					c.Ctx.SetCookie("auth", token, 7*86400)
				} else {
					c.Ctx.SetCookie("auth", token)
				}

				c.redirect(beego.URLFor(".Index"))
			}
		}
	}

	c.TplName = "main/login.html"
}

// Register 注册
func (c *MainController) Register() {
	if c.isPost() {
		valid := validation.Validation{}

		username := c.GetString("username")
		realname := c.GetString("realname")
		email := c.GetString("email")
		mobile := c.GetString("mobile")
		sex, _ := c.GetInt("sex")
		password1 := c.GetString("password1")
		password2 := c.GetString("password2")

		valid.Required(username, "username").Message("请输入登录名")
		valid.Required(realname, "realname").Message("请输入真实姓名")
		valid.Required(email, "email").Message("请输入Email")
		valid.Email(email, "email").Message("Email 无效")
		valid.Required(mobile, "mobile").Message("请输入手机号")
		valid.Required(password1, "password1").Message("请输入密码")
		valid.Required(password2, "password2").Message("请输入确认密码")
		valid.MinSize(password1, 6, "password1").Message("密码长度不能小于 6 位")
		valid.Match(password1, regexp.MustCompile("^"+regexp.QuoteMeta(password2)+"$"), "password2").Message("两次输入的密码不同")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				c.showMsg(err.Message, MSG_ERR)
			}
		}

		_, err := service.UserService.AddUser(username, realname, email, mobile, password1, sex)
		c.checkError(err)

		// 更新角色
		roleIds := make([]int, 0)
		for _, v := range c.GetStrings("role_ids") {
			if roleId, _ := strconv.Atoi(v); roleId > 0 {
				roleIds = append(roleIds, roleId)
			}
		}

		c.redirect(beego.URLFor("UserController.List"))
	}

	c.Data["pageTitle"] = "用户注册"
	c.TplName = "main/register.html"
}

// Logout 退出登录
func (c *MainController) Logout() {
	c.userId = 0
	c.auth.Logout()

	c.Ctx.SetCookie("auth", "")
	c.redirect(beego.URLFor(".Login"))
}
