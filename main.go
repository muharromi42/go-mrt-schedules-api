package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muharromi42/go-mrt-schedules-api.git/modules/station"
)

func main() {
	InitiateRouter()
}

func InitiateRouter(){
	var ( 
		router = gin.Default()
		api = router.Group("/v1/api")
	)


	station.Initiate(api)

	router.Run(":8080");
}