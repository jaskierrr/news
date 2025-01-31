package bootstrapper

import (
	"log/slog"

	"gopkg.in/reform.v1"

	"main/config"
	"main/internal/database"
	logger "main/internal/lib/logger"
	"main/internal/usecases"
)

var Secret string = ""

type RootBootstrapper struct {
	Infrastructure struct {
		Logger *slog.Logger
		// Server *restapi.Server
		DB *reform.DB
	}
	// Controller   controller.Controller
	Config *config.Config
	// Handlers     handlers.Handlers
	// Repository   repo.Repository
	Usecases usecases.Usecases
}

type RootBoot interface {
	registerRepositoriesAndServices(db *reform.DB)
	// registerAPIServer(cfg config.Config)
	RunAPI()
}

func New() RootBoot {
	return &RootBootstrapper{
		Config: config.NewConfig(),
	}
}

func (r *RootBootstrapper) RunAPI() {
	r.Infrastructure.Logger = logger.NewLogger(r.Config.LogLevel)

	r.registerRepositoriesAndServices(r.Infrastructure.DB)
	// r.registerAPIServer(*r.Config)

	// return r.Handlers
}

func (r *RootBootstrapper) registerRepositoriesAndServices(db *reform.DB) {
	logger := r.Infrastructure.Logger
	r.Infrastructure.DB = database.NewDB(*r.Config, logger)
	// r.Repository = repo.NewUserRepo(r.Infrastructure.DB, logger)
	// r.Service = service.New(r.Repository, r.ExternalRepo, logger)
}

// func (r *RootBootstrapper) registerAPIServer(cfg config.Config) {
// 	logger := r.Infrastructure.Logger

// 	r.Handlers = handlers.New(r.Usecases, logger)
// }
