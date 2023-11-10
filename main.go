package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello, world\n")
	fmt.Print("Edit by Yu\n")

	r := gin.New()
	r.Use(gin.Recovery())

	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger())
	}

	co := cors.DefaultConfig()
	co.AllowAllOrigins = true
	co.MaxAge = 24 * 365 * time.Hour

	r.Use(cors.New(co))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})
	ginpprof.WrapGroup(r.Group("/testgit"))
	srv := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("[Error] listen Error, ", err)
		}
	}()

	quit := make(chan os.Signal, 5)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	fmt.Println("[Info] Server Shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("[Error] Server Shutdown", err)
	}
	<-time.After(5 * time.Second)
}
