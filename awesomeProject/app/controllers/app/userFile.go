package app

import (
	"awesomeProject/app/common/request"
	"awesomeProject/app/common/response"
	"awesomeProject/app/models"
	"awesomeProject/app/services"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"time"
)

func LinkRepository(c *gin.Context) {
	var form request.UserRepositoryLinkRequest
	if err := c.Bind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	userId, _ := strconv.Atoi(c.Keys["id"].(string))
	if fileList, err := services.UserFileService.UserRepositoryLink(form, userId); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, fileList)
	}
}

// 获取用户文件列表
func GetUserFile(c *gin.Context) {
	var form request.UserFileListRequest
	if err := c.Bind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	userId, _ := strconv.Atoi(c.Keys["id"].(string))
	if fileList, err := services.UserFileService.UserFileList(form, userId); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, fileList)
	}
}
func CreatDir(c *gin.Context) {
	var form request.UserDirCreateRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	userIntId, _ := strconv.Atoi(c.Keys["id"].(string))
	if errString := services.UserFileService.UserDirCreate(form, userIntId); errString == "" {
		response.Success(c, "finish")
	} else {
		response.BusinessFail(c, errString)
	}
}

// PutHeaderImage 上传用户头像

func PutFileUpload(c *gin.Context) {
	bucketName := utils.GetBucketName(c.Keys["id"].(string))
	services.CreateMinoBuket(bucketName)
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 把文件上传到minio对应的桶中
	ok := services.UploadFile(bucketName, file.Filename, fileObj, file.Size)
	if !ok {
		response.Fail(c, 401, "file上传失败")
		return
	}
	fileUrl := services.GetFileUrl(bucketName, file.Filename, time.Second*24*60*60)
	if fileUrl == "" {
		response.Fail(c, 400, "获取file失败")
		return
	}
	// insert info in db
	var uf models.UserFile
	uf.UserId, _ = strconv.Atoi(c.Keys["id"].(string))
	uf.Path = fileUrl
	if ss := strings.Split(file.Filename, "."); len(ss) > 1 {
		uf.Type = utils.GetYourType(ss[len(ss)-1])
		uf.Ext = ss[1]
	}
	uf.Name = file.Filename
	uf.ParentId, _ = strconv.Atoi(c.PostForm("parentId"))
	log.Printf(uf.Name)
	errString := services.UserFileService.PutUserFileInfo(uf)
	if errString != "" {
		log.Printf(errString)
	}
	//TODO 把用户的头像地址存入到对应user表中head_url 中
	response.Success(c, map[string]interface{}{
		"msg": "success",
	})
}

func DeleteFile(c *gin.Context) {
	var form request.UserDeleteFileRequest
	if err := c.Bind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	fmt.Println(form.Id)
	err := services.UserFileService.Delete(form.Id)
	if err != nil {
		log.Printf(err.Error())
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, map[string]interface{}{
			"msg": "success",
		})
	}

}
