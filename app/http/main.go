package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"kora-backend/app/helper"
	"kora-backend/internal/auth/delivery"
	repository2 "kora-backend/internal/auth/repository"
	"kora-backend/internal/auth/repository/postgres"
	redis2 "kora-backend/internal/auth/repository/redis"
	"kora-backend/internal/auth/usecase"
	delivery2 "kora-backend/internal/choreo/delivery"
	repository3 "kora-backend/internal/choreo/repository"
	postgres2 "kora-backend/internal/choreo/repository/postgres"
	redis3 "kora-backend/internal/choreo/repository/redis"
	usecase2 "kora-backend/internal/choreo/usecase"
	repository5 "kora-backend/internal/choreographer/repository"
	postgres4 "kora-backend/internal/choreographer/repository/postgres"
	"kora-backend/internal/common/cryptography"
	"kora-backend/internal/common/jwtauth"
	"kora-backend/internal/common/middleware"
	"kora-backend/internal/common/repository"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/common"
	repository4 "kora-backend/internal/music/repository"
	postgres3 "kora-backend/internal/music/repository/postgres"
	"log"
	"net/http"
)

var (
	appName = "kora"
)

type AppUseCase struct {
	authUC   auth.UserAuthUseCase
	choreoUC choreo.ChoreoUseCase
}

type AppHandler struct {
	handlers []common.APIPathProvider
}

type AppModule struct {
	cryptoModule *cryptography.CryptographyModule
	jwtModule    *jwtauth.JwtAuthModule
	nrAgent      *newrelic.Application
	middlewareM  *middleware.MiddlewareModule
	dbCli        *sqlx.DB
	redisCli     *redis.Client
}

func InitAppModule(cfg *helper.AppConfig) (appModule *AppModule) {
	newRelicAgent, err := newrelic.NewApplication(
		newrelic.ConfigAppName(fmt.Sprintf("%s-application", appName)),
		newrelic.ConfigLicense(cfg.MonitoringConf.NewRelicKey),
		newrelic.ConfigAppLogForwardingEnabled(cfg.MonitoringConf.EnableLogForwarding),
		newrelic.ConfigDistributedTracerEnabled(cfg.MonitoringConf.EnableDistributedTracing),
	)
	if err != nil {
		log.Fatalf("Failed to init new relic with err : %s\n", err.Error())
	}
	appModule = &AppModule{}
	appModule.nrAgent = newRelicAgent
	appModule.jwtModule, err = jwtauth.NewJwtAuthModule(cfg.JWTConf)
	if err != nil {
		log.Fatalf("Failed to init JWT auth module with err : %s\n", err.Error())
	}
	appModule.cryptoModule = cryptography.NewCryptographyModule()
	appModule.middlewareM = middleware.NewMiddlewareModule(appModule.jwtModule)
	appModule.dbCli = InitDBCLient(cfg.DBConf)
	appModule.redisCli = InitRedisClient(cfg.RediConf)
	return appModule
}

func InitRepository(module *AppModule, config *helper.AppConfig) (appRepo common.BaseRepository) {
	// Init user auth repo
	userAuthRedisRepo := redis2.NewRedisUserAuthRepository(module.redisCli)
	userAuthPostgresRepo := postgres.NewPostgresUserAuthRepository(module.dbCli)
	authRepo := repository2.NewUserAuthRepository(userAuthPostgresRepo, userAuthRedisRepo)

	// Init music repo
	musicPostgresRepo := postgres3.NewPostgresMusicRepository(module.dbCli)
	musicRepo := repository4.NewMusicRepository(musicPostgresRepo, nil)

	// Init choreographer repo
	choreographPostgresRepo := postgres4.NewPostgresChoreographerRepository(module.dbCli)
	choreographRepo := repository5.NewChoreographerRepository(choreographPostgresRepo, nil)

	// Init choreo repo
	choreoPostgresrepo := postgres2.NewPostgresChoreoRepository(module.dbCli)
	choreoRedisRepo := redis3.NewRedisChoreoRepository(module.redisCli)
	choreoRepo := repository3.NewChoreoRepository(choreoPostgresrepo, choreoRedisRepo)

	// Init base repo
	repoDS := repository.NewRepository(authRepo, choreoRepo, musicRepo, choreographRepo)
	appRepo = repository.NewBaseRepository(repoDS, config)
	return appRepo
}

func InitHandler(useCase *AppUseCase, appModule *AppModule, config *helper.AppConfig) (appHandler *AppHandler) {
	appHandler = &AppHandler{}
	appHandler.handlers = append(appHandler.handlers, delivery.NewUserAuthHandler(appModule.middlewareM, config.HandlerConf, useCase.authUC))
	appHandler.handlers = append(appHandler.handlers, delivery2.NewChoreoHandler(appModule.middlewareM, config.HandlerConf, useCase.choreoUC))
	return appHandler
}

func InitAppUseCase(appRepo common.BaseRepository, appModule *AppModule) (appUC *AppUseCase) {
	appUC = &AppUseCase{}
	appUC.authUC = usecase.NewUserAuthUseCase(appRepo, appModule.jwtModule, appModule.cryptoModule)
	appUC.choreoUC = usecase2.NewChoreoUseCase(appRepo)
	return appUC
}

func InitDBCLient(cfg *helper.DatabaseConfig) (cli *sqlx.DB) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", cfg.DriverName, cfg.Username, cfg.Password, cfg.Hostname, cfg.Port, cfg.DBName)
	cli, err := sqlx.Connect(
		cfg.DriverName,
		dsn,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return cli
}

func InitRedisClient(cfg *helper.RedisConfig) (cli *redis.Client) {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Hostname, cfg.Port),
		DB:   0, // use default DB
	})
}

func InitRouter(appHandler *AppHandler, appModule *AppModule) (router *gin.Engine) {
	router = gin.Default()
	router.Use(nrgin.Middleware(appModule.nrAgent))
	for _, handler := range appHandler.handlers {
		handler.RegisterPath(router)
	}
	return router
}

func main() {
	log.Println("Initializing config")
	cfg := helper.InitConfig(appName)
	log.Println("Initializing modules")
	appModule := InitAppModule(cfg)
	log.Println("Initializing repository")
	appRepo := InitRepository(appModule, cfg)
	log.Println("Initializing usecase")
	appUC := InitAppUseCase(appRepo, appModule)
	log.Println("Initializing handler")
	appHandler := InitHandler(appUC, appModule, cfg)
	log.Println("Initializing server")
	router := InitRouter(appHandler, appModule)
	log.Println("App successfully initialized")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Kora is up and running !",
		})
	})
	err := router.Run()
	if err != nil {
		return
	}
}
