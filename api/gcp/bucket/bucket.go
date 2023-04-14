// Package bucket is used to access GCP storage bucket
package bucket

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/cienet/ldsgo/config"
)

const timeout time.Duration = time.Second * 10

// Transcoder the function to transcode data from reader to writer
type Transcoder func(writer io.Writer, reader io.Reader) (int64, error)

// NewClient returns the storage client for handle bucket
func NewClient(ctx context.Context) *storage.Client {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Fail to new storage client: %v", err)
		panic(err)
	}
	return client
}

// Write reads from <reader> and write it to <path> of cloud storage bucket.
func Write(ctx context.Context, client *storage.Client, path string, reader io.Reader) (int64, error) {
	return TransWrite(ctx, client, path, reader, io.Copy)
}

// TransWrite reads from <reader>, tanscode and write it to <path> of cloud storage bucket.
func TransWrite(ctx context.Context, client *storage.Client, path string, reader io.Reader, transcoder Transcoder) (size int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	writer := client.Bucket(config.Config.LDSBucket).Object(path).NewWriter(ctx)
	defer func() {
		err = writer.Close() // Propagate the error if fail to close
	}()

	if transcoder == nil {
		transcoder = io.Copy
	}
	size, err = transcoder(writer, reader)
	return
}

// Delete deletes the given paths in bucket, return the first failed path if any.
func Delete(ctx context.Context, client *storage.Client, paths ...string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	bucketHandler := client.Bucket(config.Config.LDSBucket)
	for _, path := range paths {
		o := bucketHandler.Object(path)
		if err := o.Delete(ctx); err != nil {
			if errors.Is(err, storage.ErrObjectNotExist) {
				log.Printf("Ignore error, file %s does not exist while deleting", path)
			} else {
				log.Printf("Fail to delete file %s in bucket", path)
				return path, err
			}
		}
	}
	return "", nil
}
