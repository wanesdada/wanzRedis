package main

import (
	_ "redis/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

