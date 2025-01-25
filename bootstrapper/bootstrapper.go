package bootstrapper

import (
	"log/slog"
	"main/config"
	"main/internal/database"
	logger "main/internal/lib/logger"

)

var Secret string = ""

type RootBootstrapper struct {
	Infrastructure struct {
		Logger *slog.Logger
		// Server *restapi.Server
		DB     database.DB
	}
	// Controller   controller.Controller
	Config       *config.Config
	// Handlers     handlers.Handlers
	// Repository   repo.Repository
	// Service      service.Service
}

type RootBoot interface {
	registerRepositoriesAndServices(db database.DB)
	registerAPIServer(cfg config.Config) error
	RunAPI() error
}

func New() RootBoot {
	return &RootBootstrapper{
		Config: config.NewConfig(),
	}
}

func (r *RootBootstrapper) RunAPI() error {
	r.Infrastructure.Logger = logger.NewLogger(r.Config.LogLevel)

	r.registerRepositoriesAndServices(r.Infrastructure.DB)
	err := r.registerAPIServer(*r.Config)
	if err != nil {
		return err
	}

	return err
}

func (r *RootBootstrapper) registerRepositoriesAndServices(db database.DB) {
	logger := r.Infrastructure.Logger
	r.Infrastructure.DB = database.NewDB().NewConn(*r.Config, logger)
	// r.Repository = repo.NewUserRepo(r.Infrastructure.DB, logger)
	// r.Service = service.New(r.Repository, r.ExternalRepo, logger)
}

func (r *RootBootstrapper) registerAPIServer(cfg config.Config) error {
	// logger := r.Infrastructure.Logger

	// r.Handlers = handlers.New(r.Controller, logger)
	// r.Handlers.Link(api)
	// if r.Handlers == nil {
	// 	return err
	// }

	// r.Infrastructure.Server = restapi.NewServer(api)
	// r.Infrastructure.Server.Port = cfg.ServerPort
	// r.Infrastructure.Server.ConfigureAPI()
	// if err := r.Infrastructure.Server.Serve(); err != nil {
	// 	log.Fatalln(err)
	// }

	return nil
}
