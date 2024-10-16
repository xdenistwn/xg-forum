package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xdenistwn/xg-forum.git/internal/configs"
	"github.com/xdenistwn/xg-forum.git/internal/handlers/memberships"
	membershipRepo "github.com/xdenistwn/xg-forum.git/internal/repository/memberships"
	membershipSvc "github.com/xdenistwn/xg-forum.git/internal/service/memberships"
	"github.com/xdenistwn/xg-forum.git/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		config *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal initiasi config.", err)
	}

	config = configs.Get()
	log.Println("config", config)

	db, err := internalsql.Connect(config.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(config, membershipRepo)

	// membership group
	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	// listen port
	r.Run(config.Service.Port)
}
