package Services

import (
	"2024_akutansi_project/Consts"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Utils"
	"context"
	"log"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

type (
	IStorageService interface {
		UploadFile(req Dto.StorageRequest, hashed bool) (path string, err error)
		DeleteFile(req Dto.StorageRequest) (err error)
		FileExists(req Dto.StorageRequest) (exist bool, err error)
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

func (s *StorageService) UploadFile(req Dto.StorageRequest, hashed bool) (path string, err error) {
	file, err := req.File.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	if hashed {
		req.ObjectKey = filepath.Join(req.ObjectKey, Utils.GenerateUniqueFileName(req.File.Filename))
	} else {
		req.ObjectKey = filepath.Join(req.ObjectKey, req.File.Filename)
	}

	log.Println(req.ObjectKey)

	_, err = s.minio.PutObject(context.Background(), Consts.BUCKET_NAME, req.ObjectKey, file, req.File.Size, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	log.Printf("Successfully uploaded file to %s/%s\n", Consts.BUCKET_NAME, req.ObjectKey)
	return s.setFilePath(req.ObjectKey), nil
}

func (s *StorageService) DeleteFile(req Dto.StorageRequest) (err error) {
	err = s.minio.RemoveObject(context.Background(), Consts.BUCKET_NAME, req.ObjectKey, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	log.Printf("Successfully deleted file to %s/%s\n", Consts.BUCKET_NAME, req.ObjectKey)

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

func (s *StorageService) setFilePath(objectKey string) string {
	return filepath.Join(Consts.URL_MINIO + objectKey)
}
