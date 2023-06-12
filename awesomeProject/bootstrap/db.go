package bootstrap

import (
	"awesomeProject/app/models"
	"awesomeProject/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yxlimo/xormzap"
	"go.uber.org/zap"
	"strconv"
	"xorm.io/xorm"
)

func InitializeDB() *xorm.Engine {
	// 根据驱动配置进行初始化
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMySqlXorm()
	default:
		return initMySqlXorm()
	}
}

// 初始化 mysql gorm.DB
func initMySqlXorm() *xorm.Engine {
	dbConfig := global.App.Config.Database

	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"

	engine, err := xorm.NewEngine(dbConfig.Driver, dsn)
	if err != nil {
		global.App.Log.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	}
	engine.SetLogger(xormzap.Logger(zapLogger))
	engine.SetMaxIdleConns(dbConfig.MaxIdleConns)
	engine.SetMaxOpenConns(dbConfig.MaxOpenConns)
	err = checkTable(engine)
	if err != nil {
		return nil
	}
	return engine
}

func checkTable(engine *xorm.Engine) error {
	tableCheck := map[string]interface{}{
		"user":      new(models.User),
		"user_file": new(models.UserFile),
	}
	var err error
	for key, value := range tableCheck {
		if exist, _ := engine.IsTableExist(key); !exist {
			err = engine.CreateTables(value)
		}
	}
	return err
}
