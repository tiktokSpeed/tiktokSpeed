package minio_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/tiktokSpeed/tiktokSpeed/conf"
	"github.com/tiktokSpeed/tiktokSpeed/pkg/minio"
)

func TestPresigned(t *testing.T) {
	reqParams := make(url.Values)
	presignedURL, err := minio.Client().PresignedGetObject(context.Background(),
		conf.GetConf().Minio.Bucket,
		"avatar.png",
		time.Second*24*60*60,
		reqParams)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Successfully generated presigned URL", presignedURL)
}
