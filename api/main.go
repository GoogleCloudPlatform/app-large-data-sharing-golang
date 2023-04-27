// Package main is the entrypoint of the server.
package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"google/jss/ldsgo/api"
	"google/jss/ldsgo/api/files"
	"google/jss/ldsgo/config"

	"github.com/gin-gonic/gin"
)

func main() {
	f, err := os.Create("gin.log")
	if err != nil {
		log.Panicln(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	apiRouter := router.Group("/api")
	apiRouter.GET("/healthchecker", api.Healthcheck)
	apiRouter.POST("/files", files.PostFiles)
	apiRouter.GET("/files", files.GetFileList)
	apiRouter.DELETE("/files/:id", files.DeleteFile)
	apiRouter.PUT("/files/:id", files.UpdateFile)
	apiRouter.DELETE("/reset", api.Reset)

	if config.Config.MockFlag {
		mockGcp()
	}

	server := &http.Server{
		Addr:    ":" + config.Config.LDSRestPort,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
