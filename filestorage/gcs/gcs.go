package gcs

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GoogleCloudStorageAdapter struct {
	Storage *GoogleCloudStorage
}

type GoogleCloudStorage struct {
	Cfg Config
}

type Config struct {
	AuthJsonPath string
	ProjectId    string
	Host         string
	Bucket       []string
}

var c *storage.Client

func (ga *GoogleCloudStorageAdapter) Init() {
	err := ga.Storage.initProcess()
	if err != nil {
		fmt.Println(err)
	}
}

func (ga *GoogleCloudStorageAdapter) Upload(b string, path string, object io.Reader) error {
	return ga.Storage.put(b, path, object)
}

func (ga *GoogleCloudStorageAdapter) Download(b string, path string) (io.Reader, error) {
	return ga.Storage.get(b, path)
}

func (g *GoogleCloudStorage) initProcess() error {
	err := g.setEnvCredential()
	if err != nil {
		return err
	}

	ctx := context.Background()
	optHost := option.WithEndpoint(g.Cfg.Host)
	c, err := storage.NewClient(ctx, optHost)
	if err != nil {
		return err
	}
	setLocalClient(c)

	err = g.validateBucketList()
	if err != nil {
		return err
	}

	return nil
}

func (g *GoogleCloudStorage) put(b string, path string, object io.Reader) error {
	ctx := context.Background()

	attrs, err := c.Bucket(b).Object(path).Attrs(ctx)
	if err != nil && err != storage.ErrObjectNotExist {
		return fmt.Errorf("error while getting object attr")
	}

	if attrs != nil {
		return fmt.Errorf("object already exist")
	}

	wc := c.Bucket(b).Object(path).NewWriter(ctx)
	if _, err := io.Copy(wc, object); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}

func (l *GoogleCloudStorage) get(b string, path string) (io.Reader, error) {
	ctx := context.Background()
	reader, err := c.Bucket(b).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (g *GoogleCloudStorage) setEnvCredential() error {
	dir, e := os.Getwd()
	if e != nil {
		return e
	}

	dir = fmt.Sprintf("%s/%s", dir, g.Cfg.AuthJsonPath)
	url, e := url.ParseRequestURI(dir)
	if e != nil {
		return e
	}
	_, e = os.Stat(dir)
	if e != nil {
		return e
	}

	e = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", url.String())
	if e != nil {
		return e
	}

	return nil
}

func (g *GoogleCloudStorage) validateBucketList() error {
	ctx := context.Background()
	for _, v := range g.Cfg.Bucket {
		_, err := c.Bucket(v).Attrs(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func setLocalClient(client *storage.Client) {
	c = client
}
