// Package main is the entrypoint of the server.
package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/cienet/ldsgo/api"
	"github.com/cienet/ldsgo/api/files"
	"github.com/cienet/ldsgo/config"
	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	apiRouter := router.Group("/api")
	if config.Config.MockFlag {
		apiRouter.GET("/healthchecker", api.Healthcheck)
		apiRouter.POST("/files", files.MockPostFiles)
		apiRouter.GET("/files", files.MockGetFileList)
		apiRouter.DELETE("/files/:id", files.MockDeleteFile)
		apiRouter.PUT("/files/:id", files.MockUpdateFile)
	} else {
		apiRouter.GET("/healthchecker", api.Healthcheck)
		apiRouter.POST("/files", files.PostFiles)
		apiRouter.GET("/files", files.GetFileList)
		apiRouter.DELETE("/files/:id", files.DeleteFile)
		apiRouter.PUT("/files/:id", files.UpdateFile)
		apiRouter.DELETE("/reset", api.Reset)
	}

	server := &http.Server{
		Addr:    ":" + config.Config.LDSRestPort,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
