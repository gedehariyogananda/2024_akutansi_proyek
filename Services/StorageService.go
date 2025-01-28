package Services

import (
	"2024_akutansi_project/Consts"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Utils"
	"context"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
)

type (
	IStorageService interface {
		UploadFile(req Dto.StorageRequest, hashed bool) (path string, err error)
		DeleteFile(req Dto.StorageRequest) (err error)
		FileExists(req Dto.StorageRequest) (exist bool, err error)
		SignedUrl(req Dto.StorageRequest) (url string, err error)
	}

	StorageService struct {
		minio *minio.Client
	}
)

func StorageServiceProvider(minio *minio.Client) *StorageService {
	return &StorageService{
		minio: minio,
	}
}

func (s *StorageService) setFilePath(objectKey string) string {
	return Consts.URL_MINIO + "/" + Consts.BUCKET_NAME + "/" + objectKey
}

func (s *StorageService) UploadFile(req Dto.StorageRequest, hashed bool) (path string, err error) {
	file, err := req.File.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	if hashed {
		req.ObjectKey = filepath.ToSlash(filepath.Join(req.ObjectKey, Utils.GenerateUniqueFileName(req.File.Filename)))
	} else {
		req.ObjectKey = filepath.ToSlash(filepath.Join(req.ObjectKey, req.File.Filename))
	}

	_, err = s.minio.PutObject(context.Background(), Consts.BUCKET_NAME, req.ObjectKey, file, req.File.Size, minio.PutObjectOptions{
		ContentType: req.File.Header.Get("Content-Type"),
	})

	if err != nil {
		return "", err
	}

	return req.ObjectKey, nil
}

func (s *StorageService) DeleteFile(req Dto.StorageRequest) (err error) {
	err = s.minio.RemoveObject(context.Background(), Consts.BUCKET_NAME, req.ObjectKey, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *StorageService) FileExists(req Dto.StorageRequest) (exist bool, err error) {
	_, err = s.minio.StatObject(context.Background(), Consts.BUCKET_NAME, req.ObjectKey, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s *StorageService) SignedUrl(req Dto.StorageRequest) (url string, err error) {
	signedUrl, err := s.minio.PresignedGetObject(
		context.Background(),
		Consts.BUCKET_NAME,
		req.ObjectKey,
		5*time.Minute,
		nil,
	)

	if err != nil {
		return "", err
	}

	return signedUrl.String(), nil
}
