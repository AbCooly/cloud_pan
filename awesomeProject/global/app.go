package global

import (
	"awesomeProject/config"
	"awesomeProject/go-storage/storage"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	_ "github.com/minio/minio-go/v7/pkg/encrypt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *xorm.Engine
	Redis       *redis.Client
	MinioClient *minio.Client
}

var App = new(Application)

func (app *Application) Disk(disk ...string) storage.Storage {
	// 若未传参，默认使用配置文件驱动
	diskName := app.Config.Storage.Default
	if len(disk) > 0 {
		diskName = storage.DiskName(disk[0])
	}
	s, err := storage.Disk(diskName)
	if err != nil {
		panic(err)
	}
	return s
}
