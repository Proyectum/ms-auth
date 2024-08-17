package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	routes "github.com/proyectum/ms-auth/internal/adapters/in/http"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func main() {
	loadConfig()
	executeMigrations()
	r := gin.Default()
	r.GET("/ping", ping())
	routes.RegisterRoutes(r)
	s := createServer(r)
	log.Fatal(s.ListenAndServe())
}

func createServer(r *gin.Engine) *http.Server {
	port := viper.GetInt32("server.port")
	readTimeout := viper.GetDuration("server.read-timeout")
	writeTimeout := viper.GetDuration("server.write-timeout")
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
