// Package config keeps config for used Globally.
package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type config struct {
	LDSRestPort      string
	LDSBucket        string
	LDSProject       string
	ResourceBasePath string
	BucketBasePath   string
	MockFlag         bool
}

// Config is the global configuration parsed from environment variables.
var Config config

func init() {
	mockFlag, _ := strconv.ParseBool(os.Getenv("MOCK"))
	if mockFlag {
		log.Println("enable mock mode!")
	}
	resourcePath := os.Getenv("LDS_RESOURCE_PATH")
	bucketBasePath := strings.TrimLeft(resourcePath, "/")
	resourceBasePath := resourcePath[0 : len(resourcePath)-len(bucketBasePath)]
	bucketBasePath = strings.TrimRight(bucketBasePath, "/") + "/" // Make sure the path end with "/".

	Config = config{
		LDSRestPort:      os.Getenv("LDS_REST_PORT"),
		LDSBucket:        os.Getenv("LDS_BUCKET"),
		LDSProject:       os.Getenv("LDS_PROJECT"),
		ResourceBasePath: resourceBasePath,
		BucketBasePath:   bucketBasePath,
		MockFlag:         mockFlag,
	}
	log.Println("using config:", Config)
}
