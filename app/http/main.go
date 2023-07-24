package main

import (
	"fmt"
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/auth/delivery"
	repository2 "github.com/Kora-Dance/koradance-backend/internal/auth/repository"
	"github.com/Kora-Dance/koradance-backend/internal/auth/repository/postgres"
	redis2 "github.com/Kora-Dance/koradance-backend/internal/auth/repository/redis"
	"github.com/Kora-Dance/koradance-backend/internal/auth/usecase"
	delivery2 "github.com/Kora-Dance/koradance-backend/internal/choreo/delivery"
	repository3 "github.com/Kora-Dance/koradance-backend/internal/choreo/repository"
	postgres2 "github.com/Kora-Dance/koradance-backend/internal/choreo/repository/postgres"
	redis3 "github.com/Kora-Dance/koradance-backend/internal/choreo/repository/redis"
	"github.com/Kora-Dance/koradance-backend/internal/choreo/repository/s3"
	usecase2 "github.com/Kora-Dance/koradance-backend/internal/choreo/usecase"
	delivery5 "github.com/Kora-Dance/koradance-backend/internal/choreographer/delivery"
	repository5 "github.com/Kora-Dance/koradance-backend/internal/choreographer/repository"
	postgres4 "github.com/Kora-Dance/koradance-backend/internal/choreographer/repository/postgres"
	usecase5 "github.com/Kora-Dance/koradance-backend/internal/choreographer/usecase"
	"github.com/Kora-Dance/koradance-backend/internal/common/router"
	"github.com/Kora-Dance/koradance-backend/internal/domain/auth"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreographer"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/internal/domain/learning_history"
	"github.com/Kora-Dance/koradance-backend/internal/domain/purchase"
	delivery3 "github.com/Kora-Dance/koradance-backend/internal/learning_history/delivery"
	repository6 "github.com/Kora-Dance/koradance-backend/internal/learning_history/repository"
	postgres5 "github.com/Kora-Dance/koradance-backend/internal/learning_history/repository/postgres"
	redis4 "github.com/Kora-Dance/koradance-backend/internal/learning_history/repository/redis"
	usecase3 "github.com/Kora-Dance/koradance-backend/internal/learning_history/usecase"
	repository8 "github.com/Kora-Dance/koradance-backend/internal/like_save/repository"
	postgres7 "github.com/Kora-Dance/koradance-backend/internal/like_save/repository/postgres"
	repository4 "github.com/Kora-Dance/koradance-backend/internal/music/repository"
	postgres3 "github.com/Kora-Dance/koradance-backend/internal/music/repository/postgres"
	delivery4 "github.com/Kora-Dance/koradance-backend/internal/purchase/delivery"
	repository7 "github.com/Kora-Dance/koradance-backend/internal/purchase/repository"
	postgres6 "github.com/Kora-Dance/koradance-backend/internal/purchase/repository/postgres"
	redis5 "github.com/Kora-Dance/koradance-backend/internal/purchase/repository/redis"
	usecase4 "github.com/Kora-Dance/koradance-backend/internal/purchase/usecase"
	"github.com/Kora-Dance/koradance-backend/pkg/aws"
	"github.com/Kora-Dance/koradance-backend/pkg/cryptography"
	"github.com/Kora-Dance/koradance-backend/pkg/jwtauth"
	"github.com/Kora-Dance/koradance-backend/pkg/middleware"
	"github.com/Kora-Dance/koradance-backend/pkg/panics"
	"github.com/Kora-Dance/koradance-backend/pkg/repository"
	"github.com/Kora-Dance/koradance-backend/pkg/secure_otp"
	"github.com/Kora-Dance/koradance-backend/pkg/slackwebhook"
	"github.com/Kora-Dance/koradance-backend/pkg/storekit"
	"github.com/Kora-Dance/koradance-backend/pkg/whatsapp"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

var (
	appName = "kora"
)

type AppUseCase struct {
	authUC            auth.UserAuthUseCase
	choreoUC          choreo.ChoreoUseCase
	learningHistoryUC learning_history.LearningHistoryUseCase
	choreoPurchaseUC  purchase.ChoreoPurchaseUseCase
	choreographerUC   choreographer.ChoreographerUseCase
}

type AppHandler struct {
	handlers []common.APIPathProvider
}

type AppModule struct {
	otpModule      *secure_otp.SecureOtpModule
	waModule       *whatsapp.WhatsappModule
	slackModule    *slackwebhook.SlackWebhookModule
	cryptoModule   *cryptography.CryptographyModule
	jwtModule      *jwtauth.JwtAuthModule
	storeKitModule *storekit.StoreKitModule
	nrAgent        *newrelic.Application
	middlewareM    middleware.MiddlewareInterface
	dbCli          *sqlx.DB
	redisCli       *redis.Client
	awsM           *aws.AWSClient
}

func InitAppModule(cfg *helper.AppConfig) (appModule *AppModule) {
	newRelicAgent, err := newrelic.NewApplication(
		newrelic.ConfigAppName(fmt.Sprintf("%s-application", appName)),
		newrelic.ConfigLicense(cfg.MonitoringConf.NewRelicKey),
		newrelic.ConfigAppLogForwardingEnabled(cfg.MonitoringConf.EnableLogForwarding),
		newrelic.ConfigDistributedTracerEnabled(cfg.MonitoringConf.EnableDistributedTracing),
		newrelic.ConfigCodeLevelMetricsEnabled(cfg.MonitoringConf.EnableCodeLevelMetrics),
	)
	if err != nil {
		log.Fatalf("Failed to init new relic with err : %s\n", err.Error())
	}
	appModule = &AppModule{}
	appModule.slackModule, err = slackwebhook.NewSlackWebhookModule()
	if err != nil {
		log.Println("Failed to init slack module")
	}
	appModule.nrAgent = newRelicAgent
	appModule.jwtModule, err = jwtauth.NewJwtAuthModule(cfg.JWTConf)
	if err != nil {
		log.Fatalf("Failed to init JWT auth module with err : %s\n", err.Error())
	}
	appModule.storeKitModule, err = storekit.NewStoreKitModule(cfg.StoreKitVerifyConfig)
	if err != nil {
		log.Fatalf("Failed to init store kit module with err : %s\n", err.Error())
	}
	appModule.cryptoModule = cryptography.NewCryptographyModule()
	appModule.middlewareM = middleware.NewMiddlewareModule(appModule.jwtModule)
	appModule.dbCli = InitDBCLient(cfg.DBConf)
	appModule.redisCli = InitRedisClient(cfg.RediConf)
	appModule.waModule = whatsapp.NewWhatsappModule(cfg.DBConf, appModule.slackModule)
	appModule.otpModule = secure_otp.NewSecureOtpModule(cfg.SecureOtpConfig, appModule.redisCli)
	appModule.awsM, err = aws.NewAWSModule(cfg.AWSConfig)
	if err != nil {
		log.Fatalf("Failed to init AWS module with err : %s\n", err.Error())
	}
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
	choreoS3Repo := s3.NewS3ChoreoContentRepository(module.awsM, config.AWSConfig.S3Config)
	choreoRepo := repository3.NewChoreoRepository(choreoPostgresrepo, choreoRedisRepo, choreoS3Repo)

	// Init learning history repo
	learnHistoryPostgresrepo := postgres5.NewPostgresLearningHistoryRepository(module.dbCli)
	learnHistoryRedisRepo := redis4.NewRedisLearningHistoryRepository(module.redisCli)
	learnHistoryRepo := repository6.NewLearningHistoryRepository(learnHistoryPostgresrepo, learnHistoryRedisRepo)

	// Init purchase repo
	choreoPurchasePostgresRepo := postgres6.NewPostgresChoreoPurchaseRepository(module.dbCli)
	choreoPurchaseRedisRepo := redis5.NewRedisChoreoPurchaseRepository(module.redisCli)
	choreoPurchaseRepo := repository7.NewChoreoPurchaseRepository(choreoPurchasePostgresRepo, choreoPurchaseRedisRepo)

	// Init choreo like repo
	likeSaveChoreoPostgresRepo := postgres7.NewPostgresLikeSaveRepository(module.dbCli)
	likeSaveRepo := repository8.NewLikeSaveRepository(likeSaveChoreoPostgresRepo, nil)

	// Init base repo
	repoDS := repository.NewRepository(authRepo, choreoRepo, musicRepo, choreographRepo, learnHistoryRepo, choreoPurchaseRepo, likeSaveRepo)
	appRepo = repository.NewBaseRepository(repoDS, config)
	return appRepo
}

func InitHandler(useCase *AppUseCase, appModule *AppModule, config *helper.AppConfig) (appHandler *AppHandler) {
	appHandler = &AppHandler{}
	appHandler.handlers = append(appHandler.handlers, delivery.NewUserAuthHandler(appModule.middlewareM, config.HandlerConf, useCase.authUC))
	appHandler.handlers = append(appHandler.handlers, delivery2.NewChoreoHandler(appModule.middlewareM, config.HandlerConf, useCase.choreoUC))
	appHandler.handlers = append(appHandler.handlers, delivery3.NewLearningHistoryHandler(appModule.middlewareM, config.HandlerConf, useCase.learningHistoryUC))
	appHandler.handlers = append(appHandler.handlers, delivery4.NewChoreoPurchaseHandler(appModule.middlewareM, config.HandlerConf, useCase.choreoPurchaseUC))
	appHandler.handlers = append(appHandler.handlers, delivery5.NewChoreographerHandler(appModule.middlewareM, config.HandlerConf, useCase.choreographerUC))
	return appHandler
}

func InitAppUseCase(appRepo common.BaseRepository, appModule *AppModule) (appUC *AppUseCase) {
	appUC = &AppUseCase{}
	appUC.authUC = usecase.NewUserAuthUseCase(appRepo, appModule.jwtModule, appModule.cryptoModule, appModule.otpModule, appModule.waModule)
	appUC.choreoUC = usecase2.NewChoreoUseCase(appRepo)
	appUC.learningHistoryUC = usecase3.NewLearningHistoryUseCase(appRepo)
	appUC.choreoPurchaseUC = usecase4.NewChoreoPurchaseUseCase(appRepo, appModule.storeKitModule)
	appUC.choreographerUC = usecase5.NewChoreographerUseCase(appRepo)
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

func InitRouter(appHandler *AppHandler, appModule *AppModule, appConfig *helper.AppConfig) (ginEngine *gin.Engine) {
	ginRouter := gin.Default()
	ginRouter.Use(panics.CaptureGinHandler())
	ginRouter.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "service are up and running",
		})
	})
	ginRouter.GET("/testpanic", func(c *gin.Context) {
		panic("test panic")
	})
	ginRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Kora is up and running !",
		})
	})
	koraRouter := router.NewRouter(ginRouter, appModule.nrAgent)
	for _, handler := range appHandler.handlers {
		handler.RegisterPath(koraRouter)
	}
	return ginRouter
}

func main() {
	webhookUrl, err := helper.GetSlackWebhookAlertUrl()
	if err == nil {
		panics.SetOptions(&panics.Options{
			Env:             helper.Getenv(),
			SlackWebhookURL: webhookUrl,
			Tags:            panics.Tags{"host": helper.GetHostname(), "datacenter": "aws"},
		})
	}
	log.Println("Initializing config")
	cfg := helper.InitConfig(appName)
	log.Println("Initializing modules")
	appModule := InitAppModule(cfg)
	_ = helper.SendServiceStartAlert(appModule.slackModule)
	log.Println("Initializing repository")
	appRepo := InitRepository(appModule, cfg)
	log.Println("Initializing usecase")
	appUC := InitAppUseCase(appRepo, appModule)
	log.Println("Initializing handler")
	appHandler := InitHandler(appUC, appModule, cfg)
	log.Println("Initializing server")
	ginRouter := InitRouter(appHandler, appModule, cfg)
	err = ginRouter.Run(fmt.Sprintf(":%s", cfg.ServerConf.Port))
	if err != nil {
		log.Fatalln("Failed to start server", err)
	}
	log.Println("App successfully initialized")

	if err != nil {
		log.Println("error serving http ", err)
	}
	defer func() {
		appModule.waModule.Stop()
	}()
}
