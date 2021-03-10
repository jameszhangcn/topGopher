package main

import(
	"github.com/astaxie/beego"
    "github.com/jameszhangcn/redisApi")

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("Hello my test")
}

func (this *MainController) Post() {
	user := this.GetString("username")
	if user == "" {
		this.Ctx.WriteString("username not get, please check it")
	}else {
		this.Ctx.WriteString("111")
	}
}

func main(){
	beego.Router("/", &MainController{})
	beego.Run("127.0.0.1:9999")
	redisApi.String()
}