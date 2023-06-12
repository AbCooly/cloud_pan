package bootstrap

import (
	"awesomeProject/global"
	"awesomeProject/go-storage/kodo"
	"awesomeProject/go-storage/local"
)

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
	_, _ = kodo.Init(global.App.Config.Storage.Disks.QiNiu)
	// _, _ = oss.Init(global.App.Config.Storage.Disks.AliOss)
}
