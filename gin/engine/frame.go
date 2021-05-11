package engine

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
const(
	DEBUG =gin.DebugMode
)

var engine *gin.Engine
var addr string=":8888"

func GetEngine()*gin.Engine{
	return engine
}

func GetAddr()string{
	return addr
}



func Run(){

	engine := GetEngine()
	engine=gin.Default()
	gin.SetMode(DEBUG)

	//根据文件路径访问,./static下面的资源
	engine.Static("/file","static")
	//设置一个文件的特殊路由
	engine.StaticFS("/special",http.Dir("./static/huaji/2.jpg"))

	err := engine.Run(GetAddr())
	if err!=nil{
		log.Fatalln()
	}




}
