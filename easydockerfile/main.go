package main

import (
	_ "github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/docs"
	_ "github.com/wangzhezhe/Easy-Dockerfile/easydockerfile/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
