// Package files defines REST API /api/files.
package files

import (
	"strconv"
	"time"

	"google/jss/ldsgo/gcp/firestore"

	"github.com/google/uuid"
)

// NewDummyFileMeta creates some fake data for testing.
func NewDummyFileMeta(i int, tags []string) *firestore.FileMeta {
	num := i + 1
	name := "file" + strconv.Itoa(num)
	id := uuid.New().String()
	path := toBucketPath(id)

	return &firestore.FileMeta{
		ID:         uuid.New().String(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		FileMetaRecord: firestore.FileMetaRecord{
			Path:     path,
			Name:     name,
			FileSize: int64(num) * 1000,
			Tags:     tags,
			OrderNo:  getOrderNo(id),
		},
	}
}
