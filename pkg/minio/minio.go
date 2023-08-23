package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tiktokSpeed/tiktokSpeed/conf"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var (
	client *minio.Client
	once   sync.Once
)

func Client() *minio.Client {
	once.Do(func() {
		minioConf := conf.GetConf().Minio
		minioClient, err := minio.New(minioConf.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(minioConf.AccessKey, minioConf.SecretKey, ""),
			Secure: true,
		})
		if err != nil {
			panic(err)
		}
		client = minioClient
	})
	return client
}

type File struct {
	Size int64
	Body io.Reader
	Name string
}

func Upload(ctx context.Context, userId int64, f *File) (string, error) {
	keyPrefix := fmt.Sprintf("%d/%s", userId, uuid.NewString())
	fileKey := path.Join(keyPrefix, "video"+path.Ext(f.Name))
	_, err := Client().PutObject(ctx, conf.GetConf().Minio.Bucket, fileKey, f.Body, f.Size, minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("upload video error: %w", err)
	}
	fURL, err := PresignedURL(ctx, fileKey)
	if err != nil {
		return "", err
	}
	srcBuf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(fURL).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 10)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(srcBuf, os.Stdout).
		Run()
	if err != nil {
		_ = Delete(ctx, fileKey)
		return "", fmt.Errorf("ffmpeg error: %w", err)
	}
	coverKey := path.Join(keyPrefix, "cover.jpg")
	_, err = Client().PutObject(ctx, conf.GetConf().Minio.Bucket, coverKey, srcBuf, int64(srcBuf.Len()), minio.PutObjectOptions{})
	if err != nil {
		_ = Delete(ctx, fileKey)
		return "", fmt.Errorf("upload thumb error: %w", err)
	}
	return fileKey, nil
}

func PresignedURL(ctx context.Context, key string) (string, error) {
	reqParams := make(url.Values)
	presignedURL, err := Client().PresignedGetObject(ctx,
		conf.GetConf().Minio.Bucket,
		key,
		time.Second*24*60*60,
		reqParams)
	if err != nil {
		return "", fmt.Errorf("presigned url error: %w", err)
	}
	return presignedURL.String(), nil
}

func Delete(ctx context.Context, key string) error {
	return Client().RemoveObject(ctx, conf.GetConf().Minio.Bucket, key, minio.RemoveObjectOptions{})
}
