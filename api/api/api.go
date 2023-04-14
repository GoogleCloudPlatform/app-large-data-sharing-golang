// Package api the REST API of group "/api".
package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/cienet/ldsgo/config"
	"github.com/cienet/ldsgo/gcp/bucket"
	"github.com/cienet/ldsgo/gcp/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// Healthcheck is function for /api/healthchecker GET endpoint.
// This API is provided for Cloud Run to check the health of the server.
func Healthcheck(c *gin.Context) {
	c.String(http.StatusNoContent, "")
}

// Reset is function for /api/reset DELETE endpoint.
// This API resets the server, deleting all files in the system.
func Reset(c *gin.Context) {
	ctx := context.Background()
	dbClient := firestore.NewClient(ctx)
	col := dbClient.Collection("fileMeta")
	bulkwriter := dbClient.BulkWriter(ctx)
	for {
		// Delete 50 documents per time.
		iter := col.Limit(50).Documents(ctx)
		numDeleted := 0

		for {
			doc, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				log.Printf("Firestore document iteration error: %v", err)
			}

			_, err = bulkwriter.Delete(doc.Ref)
			if err != nil {
				log.Printf("Firestore document deleted %v error: %v", doc.Ref.ID, err)
			}

			numDeleted++
		}

		if numDeleted == 0 {
			bulkwriter.End()
			break
		}

		bulkwriter.Flush()
	}

	client := bucket.NewClient(ctx)
	bucketHandler := client.Bucket(config.Config.LDSBucket)
	it := bucketHandler.Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Printf("object iteration error: %v", err)
		}

		if failPath, err := bucket.Delete(ctx, client, attrs.Name); err != nil {
			if !strings.Contains(attrs.Name, "small") || errors.Is(err, storage.ErrObjectNotExist) {
				// Ignore the error of the thumbnail does not exist.
				log.Printf("Storage object (%v) deleting failed", failPath)
				c.String(400, err.Error())
				return
			}
		}
	}

	c.String(204, "success")

}
