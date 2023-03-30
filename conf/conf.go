package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"lan_api/model"
	"lan_api/util"
)

type DataStruct struct {
	MySql struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}
	LogLever string
	Redis    struct {
		Db string
	}
}

var Data = new(DataStruct)

// init 初始化配置项
func init() {
	// 获取配置文件
	if initConf() != nil {
		util.Log().Panic("init error")
		return
	}

	// 设置日志级别
	util.BuildLogger(Data.LogLever)
	//err := errors.New("test log")
	//util.LogD("失败 logD %+v", err)
	util.LogD("Data", Data)
	// 读取翻译文件
	//if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
	//	util.Log().Panic("翻译文件加载失败", err)
	//}

	// 连接数据库
	dataMysql := Data.MySql
	mysqlString := dataMysql.User + ":" +
		dataMysql.Password + "@/" +
		dataMysql.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	model.Database(mysqlString)
	//cache.Redis()
}

func initConf() error {
	data, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, Data)
	if err != nil {
		return err
	}
	return nil
}
