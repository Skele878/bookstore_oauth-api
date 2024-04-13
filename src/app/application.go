package app

import (
	"bookstore_oauth-api/src/clients/cassandra"
	"bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository/db"
	"bookstore_oauth-api/src/repository/rest"
	accesstokenservice "bookstore_oauth-api/src/services/access_token"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.GetSession()
	session.Close()
	atHandler := http.NewAccessTokenHandler(
		accesstokenservice.NewService(rest.NewRestUserRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8181")
}
