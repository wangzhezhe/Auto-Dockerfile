package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:PageController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:PageController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:BuildController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:BuildController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:BuildController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:BuildController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/controllers:TestController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

}
