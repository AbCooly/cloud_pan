package services

import (
	"awesomeProject/global"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"io"
	"log"
	"net/url"
	"time"
)

// CreateMinoBuket 创建minio 桶
func CreateMinoBuket(bucketName string) {
	err := global.App.MinioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: false})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := global.App.MinioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	//
	err = global.App.MinioClient.SetBucketPolicy(context.Background(), bucketName, "")
	if err != nil {
		fmt.Println("SetBucketPolicy error", err)
		return
	}
}

// UploadFile 上传文件给minio指定的桶中
func UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) (ok bool) {
	n, err := global.App.MinioClient.PutObject(context.Background(), bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return true
}

// GetFileUrl 获取文件url
func GetFileUrl(bucketName string, fileName string, expires time.Duration) string {
	//time.Second*24*60*60
	reqParams := make(url.Values)
	presignedURL, err := global.App.MinioClient.PresignedGetObject(context.Background(), bucketName, fileName, expires, reqParams)
	if err != nil {
		zap.L().Error(err.Error())
		return ""
	}
	return fmt.Sprintf("%s", presignedURL)
}
