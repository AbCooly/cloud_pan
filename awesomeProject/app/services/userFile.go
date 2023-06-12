package services

import (
	"awesomeProject/app/common/request"
	"awesomeProject/app/common/response"
	"awesomeProject/app/models"
	"awesomeProject/global"
	"fmt"
	"log"
	"time"
)

type userFileService struct {
	fileTable string
}

var UserFileService = userFileService{"user_file"}

//func (userService *userFileService) UserFileList(params request.UserFileListRequest) (err error, fileList models.UserRepository) {
//	has, err := global.App.DB.Table("userRepository").Where("id = ?", params.Id).Get(&fileList)
//	if err != nil {
//		fmt.Println(err)
//	} else if !has {
//		err = errors.New("id none")
//	}
//	return
//}

func (ur userFileService) PutUserFileInfo(uf models.UserFile) string {
	has, err := global.App.DB.Table(UserFileService.fileTable).Where("name = ? and user_id = ? and parent_id = ?", uf.Name, uf.UserId, uf.ParentId).Exist()
	if err != nil {
		return err.Error()
	}
	if has {
		return "文件已存在"
	}
	_, err = global.App.DB.Table(UserFileService.fileTable).Insert(uf)
	if err != nil {
		log.Printf(err.Error())
		return err.Error()
	}
	return ""
}

func (ur userFileService) UserDirCreate(req request.UserDirCreateRequest, userId int) string {
	//判断文件夹是否存在
	has, err := global.App.DB.Table(UserFileService.fileTable).Where("name = ? and user_id = ?", req.Name, userId).Exist()
	if err != nil {
		return err.Error()
	}
	if has {
		return "文件夹名称已存在"
	}
	uf := new(models.UserFile)
	uf.UserId = userId
	uf.ParentId = req.ParentId
	uf.Name = req.Name
	uf.Type = "dir"
	_, err = global.App.DB.Table(UserFileService.fileTable).Insert(uf)
	if err != nil {
		log.Printf(err.Error())
		return err.Error()
	}
	return ""
}

// UserFileList 根据用户查询用户文件列表
func (ur userFileService) UserFileList(req request.UserFileListRequest, userId int) (response.UserFileResponse, error) {
	table := ur.fileTable
	//分页参数
	size := req.Size
	if size == 0 {
		size = 20
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	session := global.App.DB.NewSession()
	//判断类型
	if req.Type != "all" {
		session = global.App.DB.Table(table).Where("parent_id=? and user_id=? and type = ?", req.Id, userId, req.Type)
	} else {
		session = global.App.DB.Table(table).Where("parent_id=? and user_id=?", req.Id, userId)
	}
	// ??
	var ufs []models.UserFile
	err := session.OrderBy("ext").Limit(size, (page-1)*size).Find(&ufs)
	resp := response.UserFileResponse{List: ufs, Count: len(ufs)}
	return resp, err
}

// UserFileList 根据用户查询用户文件列表
func (ur userFileService) UserRepositoryLink(req request.UserRepositoryLinkRequest, userId int) (response.UserFileResponse, error) {
	table := ur.fileTable
	//分页参数
	var ufs []models.UserFile

	//判断类型
	var err error
	if req.Type != "all" {
		err = global.App.DB.Table(table).Where("parent_id=? and user_id=? and type = ?", req.ParentId, userId, req.Type).Find(&ufs)
		if err != nil {
			return response.UserFileResponse{}, err
		}
	} else {
		err = global.App.DB.Table(table).Where("parent_id=? and user_id=?", req.ParentId, userId).Find(&ufs)
		if err != nil {
			return response.UserFileResponse{}, err
		}
	}
	resp := response.UserFileResponse{List: ufs, Count: len(ufs)}
	return resp, err
}

// GetFileNameByUser 根据用户查询用户文件是否存在
func (ur userFileService) GetFileNameByUser(req request.UserFileNameEditRequest) int64 {
	count, err := global.App.DB.Where("name = ? and parent_id = (select parent_id from user_repository ur Where ur.identity = ?)", req.Name, req.Identity).Count(&ur)
	if err != nil {
		return 0
	}
	return count
}

// Edit 修改数据
func (ur userFileService) Edit(req request.UserFileNameEditRequest) (int64, error) {
	return global.App.DB.Where("id=? AND name=?", req.Identity, req.Name).Update(ur)
}

// GetByName 根据名称查询数据
func (ur userFileService) GetByName(userId int, name string) (uf models.UserFile, err error) {
	_, err = global.App.DB.Where("name = ? and user_id = ?", name, userId).Get(&uf)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// Delete 删除用户文件数据
func (ur userFileService) Delete(id int) error {
	_, err := global.App.DB.Where("id = ?", id).Delete(new(models.UserFile))
	return err
}

// GetByIdentityAndUserIdentity 根据Identity和UserIdentity查询资源
//func (ur userFileService) GetByIdentityAndUserIdentity() (*UserRepository, error) {
//	_, err := global.App.DB.Where("id=? AND userId = ?", ur.Id, ur.UserId).Get(&ur)
//	if err != nil {
//		return nil, err
//	}
//	return &ur, nil
//}

// GetUserById 根据Identity和UserIdentity查询资源
func (ur userFileService) GetUserById(req request.GetUserRepositoryByIdRequest) (res models.UserFile, err error) {
	_, err = global.App.DB.Where("id=?", req.Id).Get(&res)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// 根据parentId查询下面是否有文件

func (ur userFileService) GetParentIdCount(parent int) (int64, error) {
	return global.App.DB.Table(ur.fileTable).Where("parent_id = ? ", parent).Where("delete_time = ? OR delete_time IS NULL", time.Time{}.Format("2006-01-02 15:04:05")).Count()
}

// GetByIdentityAndUserIdentity 根据Identity和UserIdentity查询资源

//func (ur userFileService) GetByRepositoryIdentityAndUserIdentity() (int64, error) {
//	return global.App.DB.Table(ur.TableName()).Where("repository_identity=? AND user_identity = ? And parent_id = ?", ur.RepositoryId, ur.UserId, ur.ParentId).Where("delete_time = ? OR delete_time IS NULL", time.Time{}.Format("2006-01-02 15:04:05")).Count()
//}
