package models

import "time"

type Media struct {
	ID         int64     `json:"id" xorm:"pk"`
	DiskType   string    `json:"disk_type" xorm:"index;not null;comment('存储类型')"`
	SrcType    int8      `json:"src_type" xorm:"notnull;comment('链接类型 1相对路径 2外链')"`
	Src        string    `json:"src" xorm:"notnull;comment('资源链接"')`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
