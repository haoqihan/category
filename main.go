package main

import (
	"github.com/haoqihan/category/common"
	"github.com/haoqihan/category/domain/repository"
	service2 "github.com/haoqihan/category/domain/service"
	"github.com/haoqihan/category/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	category "github.com/haoqihan/category/proto/category"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		// 设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加 consul 作为注册中心
		micro.Registry(consulRegister),

	)
	// 获取mysql配置,路径中不用带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	// 链接数据库
	db, err := gorm.Open(
		"mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// 禁止复表
	db.SingularTable(true)

	// 初始化service
	service.Init()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{
		CategoryDataService: categoryDataService,
	})

	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
