package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("redis_secret"))

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(sessions.Sessions("session", store))

	r.GET("/login/twitter/auth", LoginByTwitter)
	r.GET("/login/twitter/auth/callback", TwitterCallback)

	r.Run()
}
