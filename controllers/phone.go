package controllers

import (
	"github.com/wwek/haoma/libs/api"
	"github.com/wwek/haoma/libs/phone"

	"github.com/astaxie/beego"
)

type pr *phone.Phone

// Operations about Phone
type PhoneController struct {
	beego.Controller
}

// @Title Login
// @Description 手机号码固号好吗标记查询
// @Param	phonenumber		query 	string	true		"电话号码, 053266114000 18870208731 051288178411"
// @Param	from      		query 	string	false		"标记来源渠道不选为所有渠道，可选360shoujiweishi,baidushoujiweishi,sogouhaomatong"
// @Success 200 {string} 返回json
// @Failure 403 失败
// @router /tag [get]
func (p *PhoneController) Get() {
	apiMsg := api.DefaultApi{}
	phonenumber := p.GetString("phonenumber")
	from := p.GetString("from")
	if phonenumber != "" {
		pr := phone.New(phonenumber)
		apiMsg.ErrCode = 0
		apiMsg.ErrMsg = "ok"
		if from != "" {
			apiMsg.Data, _ = pr.QueryOne(from)
			p.Data["json"] = apiMsg
		} else {
			apiMsg.Data, _ = pr.QueryAll()
			p.Data["json"] = apiMsg
		}

	}
	p.ServeJSON()
}
