package utils

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"net/http"
)

//router.POST("/image_upload", func(c *gin.Context) {
//	// 单文件 前端的name
//	//file, err := c.FormFile("file")
//
//	// Multipart form
//	form, err := c.MultipartForm()
//	if err != nil {
//		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
//		return
//	}
//	files := form.File["img"]
//	for _, file := range files {
//		dst := "file/" + file.Filename
//		if err := c.SaveUploadedFile(file, dst); err != nil {
//			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
//			return
//		}
//	}
//})

func GetUUID() string {
	return uuid.NewV4().String()
}

func GetBucketName(name string) string {
	return name + "xxx"
}

func GetYourType(name string) string {
	if isImage(name) {
		return "image"
	} else if isVideo(name) {
		return "video"
	} else if isDoc(name) {
		return "doc"
	}
	return "other"
}

func isImage(name string) bool {
	ns := []string{"jpg", "jpeg", "png", "bmp"}
	for _, s := range ns {
		if s == name {
			return true
		}
	}
	return false
}
func isDoc(name string) bool {
	ns := []string{"doc", "docx", "txt", "ini", "ppt", "pptx"}
	for _, s := range ns {
		if s == name {
			return true
		}
	}
	return false
}
func isVideo(name string) bool {
	ns := []string{"mp4", "avi"}
	for _, s := range ns {
		if s == name {
			return true
		}
	}
	return false
}

func save(c *gin.Context, file *multipart.FileHeader, dst string) {
	dst = "file/" + dst
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}
}
