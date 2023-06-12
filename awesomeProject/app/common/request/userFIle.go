package request

type UploadFileRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type UserRepositoryLinkRequest struct {
	ParentId           int    `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
	Type               string `json:"type"`
}

type UserFileListRequest struct {
	Id   int    `form:"id,optional"`
	Page int    `form:"page"`
	Size int    `form:"size"`
	Type string `form:"type"`
}

type UserFileNameEditRequest struct {
	ParentID string `json:"parent_id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserDirCreateRequest struct {
	ParentId int    `json:"parentId"`
	Name     string `json:"name"`
}

type UserDeleteFileRequest struct {
	Id int `json:"id,optional" form:"id"`
}

type GetUserRepositoryByIdRequest struct {
	Id int `json:"identity,optional" form:"identity"`
}

type UserFileMovedRequest struct {
	Identity       string `json:"identity"`
	ParentIdentity string `json:"parentIdentity"`
}
