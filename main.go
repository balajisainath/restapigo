package main

import (
	"log"

	"github.com/balajisainath/restapigo/database"
	"github.com/balajisainath/restapigo/routers"
	
)

func main() {
	database.Setup() //established datbase connection
	engine := routers.Router()//get the customized engine 
	err:=engine.Run("127.0.0.1:8080") //start the engine
if err!=nil{
	log.Fatal(err)
}
}
