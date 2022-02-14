package routers

import (
	"github.com/balajisainath/restapigo/database"
	"github.com/balajisainath/restapigo/handler"
	"github.com/gin-gonic/gin"
)


func Router()*gin.Engine{
	router:=gin.Default()//getting default engine for customization
	api:=handler.Handler{
		 DB: database.GetDB(), //get the handler db
	}

	router.GET("/books",api.GetBooks) //set the function for  http://localhost:8080/books:Get request
	//while calling handler function ,gin will pass gin.context in the handler function  

	router.POST("/books,",api.SaveBook )
	return router
}