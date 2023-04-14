// Package files defines REST API /api/files.
package files

import (
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MockPostFiles returns fake data for files POST endpoint.
func MockPostFiles(c *gin.Context) {
	obj := FileUploadRequest{}
	var filesarray []map[string]interface{}

	if err := c.Bind(&obj); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	form, _ := c.MultipartForm()
	files := form.File["files"]
	tags := form.Value["tags"][0]
	tagList := strings.Fields(tags)

	for _, file := range files {
		log.Println("process uploaded file:", file.Filename)
		id := uuid.New().String()

		filesarray = append(
			filesarray,
			map[string]interface{}{
				"id":         id,
				"name":       file.Filename,
				"tags":       tagList,
				"url":        toResourceURL(id),
				"thumbURL":   toResourceURL(toThumbnailPath(id)),
				"size":       10635,
				"orderNo":    "1679474966505" + "-" + id,
				"createTime": "2023-02-24T09:12:59.355Z",
				"updateTime": "2023-02-24T09:12:59.355Z",
			})
		log.Printf("Uploaded file: %v\n", file.Filename)
	}
	c.JSON(http.StatusCreated, gin.H{"files": filesarray})
}

// MockGetFileList return fake data for files GET endpoint.
func MockGetFileList(c *gin.Context) {
	FileMetaSlice := []FileMeta{}
	for i := 0; i < 100; i++ {
		id := uuid.New().String()
		FileMetaSlice = append(FileMetaSlice, FileMeta{
			ID:         id,
			Name:       "file_" + id + ".jpg",
			Tags:       []string{"image", "photo"},
			URL:        toResourceURL(id),
			ThumbURL:   toResourceURL(toThumbnailPath(id)),
			FileSize:   10635,
			OrderNo:    strconv.FormatInt(time.Now().UnixMilli(), 10) + "-" + id,
			CreateTime: time.Now().Format("2006-01-02T15:04:05.000Z"),
			UpdateTime: time.Now().Format("2006-01-02T15:04:05.000Z"),
		})
	}
	data := map[string][]FileMeta{"files": FileMetaSlice}
	c.JSON(200, data)
}

// MockUpdateFile returns Fake data for files/{id} UPDATE endpoint.
func MockUpdateFile(c *gin.Context) {
	id := c.Param("id")
	form, _ := c.MultipartForm()
	var file *multipart.FileHeader
	if len(form.File["file"]) != 0 {
		file = form.File["file"][0]
	} else {
		file = nil
	}
	tags := parseTags(form.Value["tags"][0])

	newPath := uuid.New().String()

	data := FileMeta{
		ID:         id,
		Name:       file.Filename,
		Tags:       tags,
		ThumbURL:   toResourceURL(toThumbnailPath(newPath)),
		URL:        toResourceURL(newPath),
		FileSize:   10635,
		CreateTime: time.Now().Format("2006-01-02T15:04:05.000Z"),
		UpdateTime: time.Now().Format("2006-01-02T15:04:05.000Z"),
		OrderNo:    strconv.FormatInt(time.Now().UnixMilli(), 10) + "-" + id,
	}

	response := map[string]FileMeta{"files": data}

	c.JSON(200, response)
}

// MockDeleteFile returns success response for files/{id} DELETE encpoint.
func MockDeleteFile(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Delete object: %v", id)
	c.JSON(200, gin.H{"status": "success"})
}
