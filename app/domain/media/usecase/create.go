package domain_media_usecase

import (
	"context"
	"sync"

	domain "github.com/felixa1996/go_next_be/app/domain/media"
	dto "github.com/felixa1996/go_next_be/app/domain/media/dto"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *mediaUsecase) Create(c context.Context, dto dto.MediaDtoCreateInput) (domain.Media, error) {
	ctx := context.TODO()

	// process upload with go routines
	var wg sync.WaitGroup

	for _, file := range dto.Files {
		wg.Add(1)
		src, err := file.Open()
		if err != nil {
			return domain.Media{}, err
		}
		defer src.Close()

		contentType := file.Header["Content-Type"][0]
		fileName := "materials/" + file.Filename
		uri, err := u.minioWrapper.Upload(ctx, "gonextbe", fileName, src, file.Size, contentType)
		if err != nil {
			return domain.Media{}, err
		}
		u.logger.Info("Uploaded file", zap.String("Uri", uri))
		wg.Done()
	}

	wg.Wait()

	media := domain.Media{
		Id:  uuid.NewString(),
		Uri: dto.Uri,
	}

	res, err := u.repo.Create(ctx, media)
	if err != nil {
		u.logger.Error("Failed to create media usecase", zap.Error(err))
		return domain.Media{}, err
	}
	return res, nil
}
