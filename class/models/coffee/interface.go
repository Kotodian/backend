package coffee

import(
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego"
)

// func GetAllCoffee()  (*[]Coffee,error){
// 	o := orm.NewOrm()
// 	lists := new([]Coffee)
// 	_, err := o.QueryTable("coffee").All(&lists)
// 	if err == nil{
// 		return lists,nil
// 	}
// 	return lists,err
// }
func GetAllCoffee()  ([]orm.Params,error){
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM coffee").Values(&maps)
	if(err == nil && num != 0){
		return maps,nil
	}
	return maps,err
}

func FindCoffeeByType(typ string)  ([]*Coffee){

	o := orm.NewOrm()
	var coffees []*Coffee
	o.QueryTable("coffee").Filter("type",typ).All(&coffees)
	return coffees
}

func FindCoffeeById(coffeeId string) (Coffee,error) {
	o := orm.NewOrm()
	cof := Coffee{}
	cof.Coffee_id = coffeeId
	err := o.Read(&cof,"coffee_Id")
	if err == nil {
		return cof, nil
	}
	return cof,err
}

