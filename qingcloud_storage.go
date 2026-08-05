package qingcloud_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// QingCloudStorage 青云对象存储实现 (S3兼容)
type QingCloudStorage struct {
	*s3_storage.S3Storage
}

func NewQingCloudStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*QingCloudStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &QingCloudStorage{S3Storage: s3Storage}, nil
}

const StorageName = "qingcloud-storage"

func (x *QingCloudStorage) GetName() string {
	return StorageName
}

func (x *QingCloudStorage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}
