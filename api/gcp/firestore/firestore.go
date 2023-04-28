// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package firestore is used to access GCP firestore.
package firestore

import (
	"context"
	"errors"
	"log"
	"time"

	"google/jss/ldsgo/config"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const timeout time.Duration = time.Second * 10
const collectionName string = "fileMeta"

const FieldPath string = "path"       // FieldPath field name of path
const FieldName string = "name"       // FieldName field name of name
const FieldSize string = "size"       // FieldSize field name of size
const FieldTags string = "tags"       // FieldTags field name of tags
const FieldOrderNo string = "orderNo" // FieldOrderNo field name of orderNo

// FileMetaRecord is used to create FileMeta in firestore.
type FileMetaRecord struct {
	Path     string   `firestore:"path"`
	Name     string   `firestore:"name"`
	FileSize int64    `firestore:"size"`
	Tags     []string `firestore:"tags"`
	OrderNo  string   `firestore:"orderNo"`
}

// FileMeta is queried from FileMeta collection.
type FileMeta struct {
	ID         string
	CreateTime time.Time
	UpdateTime time.Time
	FileMetaRecord
}

type service interface {
	NewClient(context.Context) (Client, error)
}

// Service used to creates client for firestore handling.
var Service service = new(firestoreService)

type firestoreService struct {
}

// NewClient creates the client for firestore handling.
func (*firestoreService) NewClient(ctx context.Context) (Client, error) {
	client, err := firestore.NewClient(ctx, config.Config.LDSProject)
	if err != nil {
		return nil, err
	}
	return &firestoreClient{client: client}, err
}

// Client is the interface of the firestore client for firestore handling.
type Client interface {
	Get(context.Context, string) (*FileMeta, error)
	ListByTags(context.Context, []string, string, int) ([]*FileMeta, error)
	Create(context.Context, string, *FileMetaRecord) (*FileMeta, error)
	Merge(context.Context, string, map[string]interface{}) (*FileMeta, error)
	Delete(context.Context, string) error
	DeleteAll(context.Context) error
	Close() error
}

type firestoreClient struct {
	client *firestore.Client
}

// Close close the underlying client.
func (c *firestoreClient) Close() error {
	return c.client.Close()
}

// Get gets the FileMeta from given ID.
func (c *firestoreClient) Get(ctx context.Context, id string) (*FileMeta, error) {
	doc := c.client.Collection(collectionName).Doc(id)
	return getByDoc(ctx, doc)
}

func getByDoc(ctx context.Context, doc *firestore.DocumentRef) (*FileMeta, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	snapshot, err := doc.Get(ctx)
	if err != nil {
		log.Printf("failed to get document from firestore: %v", err)
		return nil, err
	}
	result, err := toResult(snapshot)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func toResult(snapshot *firestore.DocumentSnapshot) (*FileMeta, error) {
	var result = new(FileMeta)
	if err := snapshot.DataTo(result); err != nil {
		log.Printf("failed to format result snapshot: %v, error: %v", snapshot, err)
		return result, err
	}
	result.ID = snapshot.Ref.ID
	result.CreateTime = snapshot.CreateTime
	result.UpdateTime = snapshot.UpdateTime
	return result, nil
}

// ListByTags lists the FileMeta from given tags.
func (c *firestoreClient) ListByTags(ctx context.Context, tags []string, orderNo string, size int) ([]*FileMeta, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	files := c.client.Collection(collectionName)
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
			log.Printf("firestore: failed to read from firestore, query: %v", query)
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

// Create creates the FileMeta associated with the given ID.
func (c *firestoreClient) Create(ctx context.Context, id string, record *FileMetaRecord) (*FileMeta, error) {
	return c.set(ctx, id, record)
}

// Merge updates the FileMeta identified by given ID, it only updates the given fields and leaves others untouched.
func (c *firestoreClient) Merge(ctx context.Context, id string, fields map[string]interface{}) (*FileMeta, error) {
	return c.set(ctx, id, fields, firestore.MergeAll)
}

func (c *firestoreClient) set(ctx context.Context, id string, record interface{}, opt ...firestore.SetOption) (*FileMeta, error) {
	ctxSet, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	doc := c.client.Collection(collectionName).Doc(id)
	if _, err := doc.Set(ctxSet, record, opt...); err != nil {
		log.Printf("firestore: failed to write to firestore, error:%v", err)
		return nil, err
	}
	return getByDoc(ctx, doc)
}

// Delete deletes the document identified by the given id.
func (c *firestoreClient) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if _, err := c.client.Collection(collectionName).Doc(id).Delete(ctx); err != nil {
		log.Printf("firestore: failed to delete document from firestore: %v", err)
		return err
	}
	return nil
}

// DeleteAll deletes all the documents in the colletion
func (c *firestoreClient) DeleteAll(ctx context.Context) error {
	col := c.client.Collection(collectionName)
	bulkwriter := c.client.BulkWriter(ctx)
	for {
		// Delete 50 documents each time.
		iter := col.Limit(50).Documents(ctx)
		numDeleted := 0

		for {
			doc, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				log.Printf("firestore: document iteration error: %v", err)
				return err
			}

			_, err = bulkwriter.Delete(doc.Ref)
			if err != nil {
				log.Printf("firestore: document deleted %v error: %v", doc.Ref.ID, err)
				return err
			}
			numDeleted++
		}

		if numDeleted == 0 {
			bulkwriter.End()
			break
		}
		bulkwriter.Flush()
	}
	return nil
}
