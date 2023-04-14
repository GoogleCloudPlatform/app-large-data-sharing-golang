// Package firestore is used to access GCP firestore
package firestore

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/cienet/ldsgo/config"
	"google.golang.org/api/iterator"
)

const timeout time.Duration = time.Second * 10
const collectionName string = "fileMeta"

const FieldPath string = "path"       // FieldPath field name of path
const FieldName string = "name"       // FieldName field name of name
const FieldSize string = "size"       // FieldSize field name of size
const FieldTags string = "tags"       // FieldTags field name of tags
const FieldOrderNo string = "orderNo" // FieldOrderNo field name of orderNo

// FileMetaRec is used to create FileMeta in firestore
type FileMetaRec struct {
	Path     string   `firestore:"path"`
	Name     string   `firestore:"name"`
	FileSize int64    `firestore:"size"`
	Tags     []string `firestore:"tags"`
	OrderNo  string   `firestore:"orderNo"`
}

// FileMeta is queried from FileMeta collection
type FileMeta struct {
	ID         string
	CreateTime time.Time
	UpdateTime time.Time
	FileMetaRec
}

// NewClient returns the firestore client
func NewClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, config.Config.LDSProject)
	if err != nil {
		log.Printf("Fail to new firestore client: %v", err)
		panic(err)
	}
	return client
}

// Get gets the FileMeta from given ID
func Get(ctx context.Context, dbClient *firestore.Client, id string) (*FileMeta, error) {
	doc := dbClient.Collection(collectionName).Doc(id)
	return getByDoc(ctx, doc)
}

func getByDoc(ctx context.Context, doc *firestore.DocumentRef) (*FileMeta, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	snapshot, err := doc.Get(ctx)
	if err != nil {
		log.Println("Fail to get document from firestore:", err)
		return nil, err
	}
	result, err := toResult(snapshot)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func toResult(snapshot *firestore.DocumentSnapshot) (*FileMeta, error) {
	result := &FileMeta{}
	if err := snapshot.DataTo(result); err != nil {
		log.Println("Fail to format result snapshot:", snapshot, "err:", err)
		return result, err
	}
	result.ID = snapshot.Ref.ID
	result.CreateTime = snapshot.CreateTime
	result.UpdateTime = snapshot.UpdateTime
	return result, nil
}

// ListByTags lists the FileMeta from given tags
func ListByTags(ctx context.Context, client *firestore.Client, tags []string, orderNo string, size int) ([]*FileMeta, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	files := client.Collection(collectionName)
	var query firestore.Query
	if len(tags) == 0 {
		query = files.OrderBy("orderNo", firestore.Desc)
	} else {
		query = files.Where("tags", "array-contains-any", tags).OrderBy("orderNo", firestore.Desc)
	}
	if orderNo != "" {
		query = query.StartAfter(orderNo)
	}
	query = query.Limit(size)
	iter := query.Documents(ctx)
	defer iter.Stop()
	var results []*FileMeta
	for {
		snapshot, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Println("Fail to read from firestore, query:", query)
			return nil, err
		}
		result, err := toResult(snapshot)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

// Create creates the FileMeta associated with the given ID
func Create(ctx context.Context, dbClient *firestore.Client, id string, record *FileMetaRec) (*FileMeta, error) {
	return set(ctx, dbClient, id, record)
}

// Merge updates the FileMeta identified by given ID, it only updates the given fields and leaves others untouched
func Merge(ctx context.Context, dbClient *firestore.Client, id string, fields *map[string]interface{}) (*FileMeta, error) {
	return set(ctx, dbClient, id, *fields, firestore.MergeAll)
}

func set(ctx context.Context, dbClient *firestore.Client, id string, record interface{}, opt ...firestore.SetOption) (*FileMeta, error) {
	ctxSet, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	doc := dbClient.Collection(collectionName).Doc(id)
	if _, err := doc.Set(ctxSet, record, opt...); err != nil {
		log.Printf("Failed to write to firestore, error:%v", err)
		return nil, err
	}
	return getByDoc(ctx, doc)
}

// Delete removes the FileMeta identified by id.
func Delete(ctx context.Context, dbClient *firestore.Client, id string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if _, err := dbClient.Collection(collectionName).Doc(id).Delete(ctx); err != nil {
		log.Printf("Fail to delete document from firestore: %v", err)
		return err
	}
	return nil
}
