package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opensourceways/community-robot-lib/interrupts"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"project/xihe-statistics/config"
	"project/xihe-statistics/controller"
	"project/xihe-statistics/docs"
	"project/xihe-statistics/infrastructure/gitlab"
	"project/xihe-statistics/infrastructure/pgsql"
	"project/xihe-statistics/infrastructure/repositories"
)

type Service struct {
	Log *logrus.Entry

	Port    int
	Timeout time.Duration
}

func StartWebServer(port int, timeout time.Duration, cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logRequest())

	setRouter(r, cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	defer interrupts.WaitForGracefulShutdown()

	interrupts.ListenAndServe(srv, timeout)
}

//setRouter init router
func setRouter(engine *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "xihe-statistics"
	docs.SwaggerInfo.Description = ""

	bigModelRecord := repositories.NewBigModelRecordRepository(
		pgsql.NewBigModelMapper(pgsql.BigModelRecord{}),
	)

	repoRecord := repositories.NewUserWithRepoRepository(
		pgsql.NewUserWithRepoMapper(pgsql.UserWithRepo{}),
	)

	registerRecord := repositories.NewRegisterRecordRepository(
		pgsql.NewRegisterRecordMapper(pgsql.RegisterRecord{}),
	)

	fileUploadRecord := repositories.NewFileUploadRecordRepository(
		pgsql.NewFileUploadRecordMapper(pgsql.FileUploadRecord{}),
	)

	downloadRecord := repositories.NewDownloadRecordRepository(
		pgsql.NewDownloadRecordMapper(pgsql.DownloadRecord{}),
	)

	trainRecord := repositories.NewTrainRecordRepository(
		pgsql.NewTrainRecordMapper(pgsql.TrainRecord{}),
	)

	gitlabRecord := repositories.NewGitLabRecordRepository(
		pgsql.NewGitLabRecordMapper(pgsql.GitLabRecord{}),
	)

	platform := gitlab.NewGitlabStatistics(cfg)

	// controller -> gin
	v1 := engine.Group(docs.SwaggerInfo.BasePath)
	{
		// domain.repository -> app -> controller (NewxxxxService | AddxxxxController)
		controller.AddRouterForBigModelRecordController(
			v1, bigModelRecord,
		)

		controller.AddRouterForRepoRecordController(
			v1, repoRecord,
		)

		controller.AddRouterForD1Controller(
			v1, bigModelRecord, repoRecord,
		)

		controller.AddRouterForRegisterRecordController(
			v1, registerRecord,
		)

		controller.AddRouterForFileUploadRecordController(
			v1, fileUploadRecord,
		)

		controller.AddRouterForDownloadRecordController(
			v1, downloadRecord, platform, gitlabRecord,
		)

		controller.AddRouterForTrainRecordController(
			v1, trainRecord,
		)
	}

	engine.UseRawPath = true
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		logrus.Infof(
			"| %d | %d | %s | %s |",
			c.Writer.Status(),
			endTime.Sub(startTime),
			c.Request.Method,
			c.Request.RequestURI,
		)
	}
}
