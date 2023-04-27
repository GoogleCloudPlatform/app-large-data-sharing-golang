// Package main is the entrypoint of the server.
package main

import (
	"google/jss/ldsgo/api/files"
	"google/jss/ldsgo/gcp/bucket"
	"google/jss/ldsgo/gcp/firestore"

	"github.com/stretchr/testify/mock"
)

func mockGcp() {
	mockBucket()
	mockFirestore()
}

func mockFirestore() {
	_, firestoreClient := firestore.MockService()

	tags := []string{"tag1", "tag2"}

	fileMeta := files.NewDummyFileMeta(1, tags)
	firestoreClient.On("Get", mock.Anything, mock.Anything).Return(fileMeta, nil)

	var fileMetas []*firestore.FileMeta
	for i := 0; i < 50; i++ {
		fileMetas = append(fileMetas, files.NewDummyFileMeta(i, tags))
	}
	firestoreClient.On("ListByTags", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(fileMetas, nil)

	firestoreClient.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(fileMeta, nil)
	firestoreClient.On("Merge", mock.Anything, mock.Anything, mock.Anything).Return(fileMeta, nil)
	firestoreClient.On("Delete", mock.Anything, mock.Anything).Return(nil)
	firestoreClient.On("DeleteAll", mock.Anything).Return(nil)

}

func mockBucket() {
	_, bucketClient := bucket.MockService()
	bucketClient.On("TransWrite", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(1000, nil)
	bucketClient.On("Delete", mock.Anything, mock.Anything).Return("", nil)
	bucketClient.On("DeleteAll", mock.Anything).Return(nil)
}
