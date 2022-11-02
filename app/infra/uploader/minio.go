package uploader

import (
	"context"
	"mime/multipart"

	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioWrapper struct {
	logger  *zap.Logger
	client  *minio.Client
	baseUrl string
}

func NewMinioWrapper(config config.Config, logger *zap.Logger) MinioWrapper {
	// todo need check
	// endpoint := "localhost:8998"
	// accessKeyID := "c23hil9KtahYIwKM"
	// secretAccessKey := "dxDXoKow5EhI60bRMJC8kqJcDjYsPJ6r"
	// useSSL := false

	minioClient, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretAccessKey, ""),
		Secure: config.MinioSSL,
	})
	if err != nil {
		logger.Fatal("Failed to initialize minio", zap.Error(err))
	}
	return MinioWrapper{
		baseUrl: config.MinioBaseUrl,
		logger:  logger,
		client:  minioClient,
	}
}

func (m MinioWrapper) Upload(ctx context.Context, bucketName string, fileName string, buffer multipart.File, fileSize int64, contentType string) (string, error) {
	info, err := m.client.PutObject(ctx, bucketName, fileName, buffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		m.logger.Error("Failed to upload to minio", zap.String("Content-Type", contentType), zap.Error(err))
		return "", err
	}

	m.logger.Info("Successfully uploaded", zap.String("fileName", fileName), zap.Int64("fileSize", info.Size))
	uri := m.baseUrl + bucketName + "/" + fileName
	return uri, nil
}
