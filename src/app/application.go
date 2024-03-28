package app

import (
	"bookstore_oauth-api/src/clients/cassandra"
	accesstoken "bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		fmt.Println(dbErr)
		panic(dbErr)
	}
	session.Close()
	atHandler := http.NewHandler(accesstoken.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8181")
}
