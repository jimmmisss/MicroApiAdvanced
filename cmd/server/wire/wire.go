//go:build wireinject
// +build wireinject

package wire

import (
	"microApiAdvanced/internal/handler"
	"microApiAdvanced/internal/repository"
	"microApiAdvanced/internal/server"
	"microApiAdvanced/internal/service"
	"microApiAdvanced/pkg/app"
	"microApiAdvanced/pkg/helper/sid"
	"microApiAdvanced/pkg/jwt"
	"microApiAdvanced/pkg/log"
	"microApiAdvanced/pkg/server/http"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
