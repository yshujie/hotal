package idgenerator

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

const DEFAULT_QUEUE_LEN = 10

type mysql struct {
	queue []int64
	o     orm.Ormer
}

type nextId struct {
	Nextid int64
}

func NewMysql() *mysql {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")

	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	orm.RegisterDataBase("fcqxdb", "mysql", dsn)
	o := orm.NewOrm()
	o.Using("fcqxdb")

	return &mysql{
		queue: make([]int64, 0, DEFAULT_QUEUE_LEN),
		o:     o,
	}
}

func (this *mysql) GetNextID(sizes ...int) int64 {
	var size int
	if len(sizes) > 0 {
		size = sizes[0]
	}
	if size == 0 || size > 100 {
		size = DEFAULT_QUEUE_LEN
	}
	if len(this.queue) < 1 {
		this.preLoad(size)
	}
	if len(this.queue) > 0 {
		nextid := this.queue[0]
		this.queue = this.queue[1:]
		return nextid
	}

	return 0
}

func (this *mysql) preLoad(size int) error {
	sql := fmt.Sprintf("UPDATE idgenerator SET nextid=LAST_INSERT_ID(nextid+%d)", size)

	o := this.o
	if _, err := o.Raw(sql).Exec(); err != nil {
		return err
	}

	sql = "SELECT LAST_INSERT_ID() AS nextid"
	var nextid nextId
	if err := o.Raw(sql).QueryRow(&nextid); err != nil {
		return err
	}

	beego.Trace(nextid)
	for i := size; i > 0; i-- {
		this.queue = append(this.queue, nextid.Nextid-int64(i))
	}
	beego.Trace(this.queue)
	return nil
}
