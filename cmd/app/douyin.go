package app

import (
	"flag"
	"github.com/RaymondCode/simple-demo/database"

	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/conf"
	"github.com/RaymondCode/simple-demo/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	help = flag.Bool("h", false, "to show help")
	env  = flag.String("env", "dev", "env: dev | test | prod")
)

func InitDouYinService() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}

	// load config
	err := conf.ConfigLoad(*env)
	if err != nil {
		log.Errorf("main(): in ConfigLoad() error:%s", err.Error())
		common.AbnormalExit()
	}

	// init datasource
	err = initDatabase()
	if err != nil {
		log.Errorf("main(): in common_database.GetInstanceConnection() error:%s", err.Error())
		common.AbnormalExit()
	}

	//init router
	r := gin.Default()
	router.InitRouter(r)
	serverPort := conf.ServerConfig.Port
	err = r.Run("0.0.0.0:" + serverPort)
	if err != nil {
		return
	}

}

// 初始化DB
func initDatabase() error {
	err := database.GetInstanceConnection().Init()
	if err != nil {
		return err
	}
	return nil
}
