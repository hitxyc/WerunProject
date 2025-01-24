package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	Register(r)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
