# app-large-data-sharing-golang

## Docker
### Start Web and Mocked REST server
```bash
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

# Frontend

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 15.2.0.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
