package models

import (
	"class/models/address"
	"class/models/types"
	"class/models/coffee"
	"class/models/user"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Initdb()  {
	dbHost := beego.AppConfig.String("dbhost")
	dbPort := beego.AppConfig.String("dbport")
	dbUser := beego.AppConfig.String("dbuser")
	dbPass := beego.AppConfig.String("dbpass")
	dbName := beego.AppConfig.String("dbname")
	dbStr := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/"+dbName + "?charset=utf8"
	e := orm.RegisterDataBase("default","mysql",dbStr)
	orm.RegisterModel(new(user.User),new(coffee.Coffee),new(types.Type),new(address.UserAddress))
	if e != nil {
		panic(e.Error())
	}	
	e = orm.RunSyncdb("default",false,true)
	if e != nil {
		panic(e.Error())
	}
}