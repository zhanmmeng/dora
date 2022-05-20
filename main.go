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
)

type User struct {
	gorm.Model
	Name     string
	Phone    string
	Password string
}

func main() {

	db := common.InitDB()
	defer db.Close()

	// 1.创建路由
	r := gin.Default()

	r = router.CollectRouter(r)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	panic(r.Run(":8000"))
}


