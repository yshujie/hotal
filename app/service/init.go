package service

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hotel/app/entity"
	"github.com/hotel/app/entity/idgenerator"
	"net/url"
)

var (
	o                orm.Ormer
	tablePrefix      string
	tableSuffix      string
	UserService      *userService
	HotelService     *hotelService
	RoomService      *roomService
	CustomerService  *customerService
	BookOrderService *bookOrderService
)

func init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	tablePrefix = beego.AppConfig.String("db.prefix")
	tableSuffix = beego.AppConfig.String("db.suffix")

	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	orm.RegisterDataBase("default", "mysql", dsn)
	// todo：这是在做什么？
	orm.RegisterModelWithSuffix(tableSuffix,
		new(entity.Hotel),
		new(entity.User),
		new(entity.Room),
		new(entity.Customer),
		new(entity.BookOrder),
	)

	// todo：这是在做什么？
	if beego.AppConfig.String("runmode") != "prod" {
		orm.Debug = true
	}

	// todo：这是在做什么？
	o = orm.NewOrm()
	orm.RunCommand()

	// 初始化ID生成器
	initIDGenerator()

	// 初始化服务
	initService()
}

func initIDGenerator() {
	idgenerator.Init()
}

func initService() {
	UserService = &userService{}
	HotelService = &hotelService{}
	RoomService = &roomService{}
	CustomerService = &customerService{}
	BookOrderService = &bookOrderService{}
}

// 获取真实表名
func tableName(name string) string {
	return tablePrefix + name + tableSuffix
}

func GetOrm() orm.Ormer {
	return o
}

func concatenateError(err error, stderr string) error {
	if len(stderr) == 0 {
		return err
	}
	return fmt.Errorf("%v: %s", err, stderr)
}

func DBVersion() string {
	var lists []orm.ParamsList
	o.Raw("SELECT VERSION()").ValuesList(&lists)
	return lists[0][0].(string)
}

// 获取分页参数
func getPagination(page, pageSize int) (offset, limit int) {
	if pageSize == -1 {
		offset = 0
		limit = 1000
	} else {
		offset = calcOffset(page, pageSize)
		limit = pageSize
	}

	return offset, limit
}

func calcOffset(page, pageSize int) (offset int) {
	if page <= 0 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}

	return offset
}
