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

	"project/xihe-statistics/controller"
	"project/xihe-statistics/docs"
	"project/xihe-statistics/infrastructure/pgsql"
	"project/xihe-statistics/infrastructure/repositories"
)

type Service struct {
	Log *logrus.Entry

	Port    int
	Timeout time.Duration
}

func StartWebServer(port int, timeout time.Duration) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logRequest())

	setRouter(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	defer interrupts.WaitForGracefulShutdown()

	interrupts.ListenAndServe(srv, timeout)
}

//setRouter init router
func setRouter(engine *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "xihe-statistics"
	docs.SwaggerInfo.Description = ""

	// collections := config.Conf.Mongodb.MongodbCollections

	// // infrastructure.repositories -> domain.repository (NewxxxxRepository)
	// bigModelRecord := repositories.NewBigModelRecordRepository(
	// 	// infrastructure.mongodb -> infrastructure.repositories (mapper)
	// 	mongodb.NewBigModelMapper(collections.BigModel),
	// )

	bigModelRecord := repositories.NewBigModelRecordRepository(
		// infrastructure.mongodb -> infrastructure.repositories (mapper)
		pgsql.NewBigModelMapper(pgsql.BigModelRecord{}),
	)

	// repoRecord := repositories.NewUserWithRepoRepository(
	// 	mongodb.NewUserWithRepoMapper(collections.Repo),
	// )

	repoRecord := repositories.NewUserWithRepoRepository(
		pgsql.NewUserWithRepoMapper(pgsql.UserWithRepo{}),
	)

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
