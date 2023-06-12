package bootstrap

import (
	"awesomeProject/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func InitMinIO() *minio.Client {

	minioInfo := global.App.Config.MinIO
	log.Println(minioInfo)
	// Initialize minio client object.false是关闭https证书校验
	minioClient, err := minio.New(minioInfo.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioInfo.AccessKeyID, minioInfo.SecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}
	//客户端注册到全局变量中
	return minioClient

}
