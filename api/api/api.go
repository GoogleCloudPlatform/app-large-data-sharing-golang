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

// Package api the REST API of group "/api".
package api

import (
	"context"
	"log"
	"net/http"

	"google/jss/ldsgo/gcp/bucket"
	"google/jss/ldsgo/gcp/firestore"

	"github.com/gin-gonic/gin"
)

// Healthcheck is function for /api/healthchecker GET endpoint.
// This API is provided for Cloud Run to check the health of the server.
func Healthcheck(c *gin.Context) {
	c.String(http.StatusNoContent, "")
}

// Reset is function for /api/reset DELETE endpoint.
// This API resets the server, deleting all files in the system.
func Reset(c *gin.Context) {
	log.Println("Start to reset server")
	ctx := context.Background()

	dbClient, err := firestore.Service.NewClient(ctx)
	if err != nil {
		log.Panicln(err)
	}
	defer dbClient.Close() // nolint: errcheck

	if err := dbClient.DeleteAll(ctx); err != nil {
		log.Panicln(err)
	}

	client, err := bucket.Service.NewClient(ctx)
	if err != nil {
		log.Panicln(err)
	}
	defer client.Close() // nolint: errcheck

	if err := client.DeleteAll(ctx); err != nil {
		c.String(400, err.Error())
		return
	}
	c.String(204, "success")
}
