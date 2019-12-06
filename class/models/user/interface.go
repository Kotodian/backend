package user

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func Login(username string,password string) (User,error){
	o := orm.NewOrm()
	u := User{}
	u.Username = username
	e := o.Read(&u,"Username")
	if e == nil{
		if u.PasswordHash ==  password{
			
			return u, nil
		}
		return User{Username:username} ,PasswordError
	}
	beego.Info(username)
	return u,UsernameNotFind
}

func Register(username string,password string,image string)  (User,error){
	o := orm.NewOrm()
	u := User{}
	u.Username = username
	e := o.Read(&u,"Username")
	if e == orm.ErrNoRows {
			u.PasswordHash = password
			u.Image = image
			_,e = o.Insert(&u)
			return u,e
		}
		return User{Username:username},UserExist
}


