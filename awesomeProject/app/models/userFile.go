package models

import "time"

type UserFile struct {
	Id           int       `json:"id" xorm:"pk autoincr"`
	UserId       int       `json:"user_id"`
	ParentId     int       `json:"parent_id"`
	RepositoryId string    `json:"repository_identity"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Ext          string    `json:"ext"`
	Path         string    `json:"path" xorm:"TEXT"`
	Created      time.Time `xorm:"created"`
	Updated      time.Time `xorm:"updated"`
	Deleted      time.Time `xorm:"deleted"`
}
