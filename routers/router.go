// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/wwek/haoma/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/phone",
			beego.NSInclude(
				&controllers.PhoneController{},
			),
		),
		//beego.NSNamespace("/user",
		//	beego.NSInclude(
		//		&controllers.UserController{},
		//	),
		//),
	)
	beego.AddNamespace(ns)
	beego.Get("/", func(ctx *context.Context) {
		ctx.Redirect(302, "/gui/")
		//ctx.Output.Body([]byte("hello world"))
	})
}
