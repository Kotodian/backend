package types

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func GetType()  ([]orm.Params, error){
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM type").Values(&maps)
	if(err == nil && num != 0){
		return maps,nil
	}
	beego.Info(maps)
	return maps,err
}