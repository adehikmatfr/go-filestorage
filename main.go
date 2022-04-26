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
				AuthJsonPath: "assets/pubsub-credential.json",
				ProjectId:    "clodeo-universal-staging",
				Bucket:       []string{"clodeo-rebuild-content"},
			},
		},
	}
	c := filestorage.NewClient(fileStorage)

	// Open local file.
	f, _ := os.Open("notes.txt")
	defer f.Close()
	c.Upload("clodeo-rebuild-content", "txt/notes.txt", f)
	c.Download("clodeo-rebuild-content", "txt/notes.txt")
}
