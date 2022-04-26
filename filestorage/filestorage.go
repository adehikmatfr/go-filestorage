package filestorage

import (
	"io"
)

type FileStorage interface {
	Init()
	Upload(bucket string, path string, object io.Reader) error
	Download(bucket string, path string) (io.Reader, error)
}

type Client struct{}

var fs FileStorage

func setLocalFileStorage(s FileStorage) {
	fs = s
}

func NewClient(fsg FileStorage) *Client {
	setLocalFileStorage(fsg)
	fsg.Init()
	return &Client{}
}

func (c *Client) Upload(bucket string, path string, object io.Reader) error {
	return fs.Upload(bucket, path, object)
}

func (c *Client) Download(bucket string, path string) (io.Reader, error) {
	return fs.Download(bucket, path)
}
