package address

import (
	"github.com/astaxie/beego/orm"
)

func GetAddress(username string)  ([]*UserAddress){
		o := orm.NewOrm()
		var userAddress []*UserAddress
		o.QueryTable("user_address").Filter("username",username).All(&userAddress)
		return userAddress
}

func SaveAddress(adre string, username string, nickname string, phone string, sex string,isdefault string)  (error){
	o := orm.NewOrm()
	userAddress := UserAddress{}
	userAddress.Address = adre
	userAddress.Username = username
	userAddress.Nickname = nickname
	userAddress.Phone = phone
	userAddress.Sex = sex
	userAddress.Isdefault = isdefault
	_,err := o.Insert(&userAddress)
	if err == nil{
		return nil
	}
	return err
}
func GetAddressById(id int)  (UserAddress){
	o := orm.NewOrm()
	adre := UserAddress{Id:id}
	err := o.Read(&adre)
	if(err == nil){
		return adre
	}
	return adre
}