package main

import (
	//"github.com/gin-gonic/gin"
	"apidemo/controllers"
	_ "apidemo/conf"
	_ "apidemo/db"
	//"apidemo/conf"
	"fmt"
	"apidemo/db"
)


func main() {
	//r := gin.Default()
	//
	//// template
	//api := r.Group("/api/v1")
	//
	//content := api.Group("/content")
	//{
	//	content.GET("/template", (&controllers.Template{}).Get)
	//}
	//
	//r.Run()


	template := controllers.Template{}
	conn := db.GetMysqlConn()
	conn.First(&template)
	fmt.Println(template.ID)
	//a()
	//a(1)
	//r := gin.Default()
	//r.GET("/ping", controllers.Get)
	//r.Run() // listen and serve on 0.0.0.0:8080
}