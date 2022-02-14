package main

import (
	"github.com/balajisainath/restapigo/database"
	"github.com/balajisainath/restapigo/routers"
)

func main() {
	database.Setup()
	engine := routers.Router
}
