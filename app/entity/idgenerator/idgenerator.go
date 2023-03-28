package idgenerator

import "github.com/astaxie/beego"

type idgenerator interface {
	GetNextID(size ...int) int64
}

var idg idgenerator

func Init() {
	idgeneratorType := beego.AppConfig.String("idgenerator")

	switch idgeneratorType {
	case "mysql":
		idg = NewMysql()
	default:
		idg = NewMysql()
	}
}

func GetNextID() int64 {
	return idg.GetNextID()
}
