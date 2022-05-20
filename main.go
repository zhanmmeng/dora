package main

import (
	"dora/Backstage/common"
	_ "dora/Backstage/controller"
	_ "dora/Backstage/model"
	"dora/Backstage/router"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
)

type User struct {
	gorm.Model
	Name     string
	Phone    string
	Password string
}

func main() {

	InitConfig()
	db := common.InitDB()
	defer db.Close()

	// 1.创建路由
	r := gin.Default()

	r = router.CollectRouter(r)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/Backstage/config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(main)
	}
}