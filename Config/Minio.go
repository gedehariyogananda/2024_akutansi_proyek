package Config

import (
	"log"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() *minio.Client {
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	endpoint := os.Getenv("MINIO_ENDPOINT")
	secure, _ := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretAccessKey, ""),
		Secure: secure,
	})

	if err != nil {
		log.Fatalln("Error connected to Minio: ", err)
	}

	log.Println("Connected to Minio server")

	return minioClient
}
