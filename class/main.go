package main

import (
	"class/models/address"
	"go/types"
	"class/models"
	"encoding/gob"
	"class/models/user"
	_ "class/routers"
	"github.com/astaxie/beego"
	"class/models/coffee"
)

func main() {
	models.Initdb()
	InitSession()
	beego.Run()
}

func InitSession(){
	gob.Register(new(user.User))
	gob.Register(new(coffee.Coffee))
	gob.Register(new(types.Type))
	gob.Register(new(address.UserAddress))
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "demo"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./data"
}
