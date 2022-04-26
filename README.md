# filestorage package
## create new client example
1. google cloud storage
```go
package main

import (
	"os"

	"gitlab.com/adehikmatfr/file-storage/filestorage"
	"gitlab.com/adehikmatfr/file-storage/filestorage/gcs"
)

func main() {
	fileStorage := &gcs.GoogleCloudStorageAdapter{
		Storage: &gcs.GoogleCloudStorage{
			Cfg: gcs.Config{
		        AuthJsonPath: "YOUR_CREDENTIAL_PATH",
		        ProjectId:    "YOUR_PROJECT_ID",
				 // not required
				 // if it is filled then your bucket will be validated
				Bucket:       []string{"clodeo-rebuild-content"},
			},
		},
	}
	c := filestorage.NewClient(fileStorage)
}
```

### upload file example
```go
	// Open local file.
	f, _ := os.Open("notes.txt")
	defer f.Close()
	c.Upload("clodeo-rebuild-content", "txt/notes.txt", f)
```

### download message example
```go
i, _ :=c.Download("clodeo-rebuild-content", "txt/notes.txt")
```