# app-large-data-sharing-golang

## Docker
### Start Web and Mocked REST server
```bash
export LDS_REST_HOST=api-mock
docker-compose up -d api-mock web
```

## Local Development Environment
### Prerequisite
- Golang
- gcloud CLi

### Authenticate gcloud
```base
gcloud auth login
gcloud auth application-default login
```


### Start REST server 
```
cd api
go run .
```

### ports

+ API server: localhost:8000
+ Web client: localhost:8080

## REST API
### get file list
http://localhost:8000/api/files?tags=<tag>&lastNo=<orderNo>

(tags and lastNo are optional)
use tag to filter files of specific tag  
file list is sorted in orderNo, and limited in 50 files per request  
use lastNo to get files after the specific file  

### Upload files example
```bash
curl -X POST http://{server_url}/api/files -H "Content-Type: multipart/form-data" -F "files=@{/path/to/file}" -F "files=@{/path/to/file}" -F "tags=tag1 tag2"
```

### Update file example
```bash
curl -X PUT http://{server_url}/api/files/{id} -H "Content-Type: multipart/form-data" -F "file=@{/path/to/file}" -F "tags=tag1 tag2"
```

### Upload files under a given folder
Upload all files under `path/to/myfolder` and tag them with `tag1` and `tag2`.

```bash
upload.sh tag1,tag2 path/to/myfolder
```
