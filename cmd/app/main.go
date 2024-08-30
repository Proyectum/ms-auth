package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	routes "github.com/proyectum/ms-auth/internal/adapters/in/http"
	"github.com/proyectum/ms-auth/internal/boot"
	"log"
	"net/http"
	"time"
)

func main() {
	boot.LoadConfig()
	boot.ExecuteMigrations()
	r := gin.Default()
	r.GET("/ping", ping())
	routes.RegisterRoutes(r)
	s := createServer(r)
	log.Fatal(s.ListenAndServe())
}

func createServer(r *gin.Engine) *http.Server {
	srvConf := boot.CONFIG.Server
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", srvConf.Port),
		Handler:        r,
		ReadTimeout:    srvConf.ReadTimeout * time.Second,
		WriteTimeout:   srvConf.WriteTimeout * time.Second,
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
