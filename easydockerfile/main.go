package main

import (
	_ "Easy-Dockerfile/easydockerfile/docs"
	_ "Easy-Dockerfile/easydockerfile/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
