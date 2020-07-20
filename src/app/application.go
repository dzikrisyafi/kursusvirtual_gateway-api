package app

import (
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/http"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/repository/rest"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/services/cohort"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/services/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_middleware/middleware"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	enrolls := http.NewEnrollsHandler(enrolls.NewService(rest.NewRestUsersRepository(), rest.NewRestCoursesRepository()))
	cohorts := http.NewCohortHandler(cohort.NewService(rest.NewRestUsersRepository(), rest.NewRestCoursesRepository()))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Content-Length"},
	}))

	enrollsGroup := router.Group("/enrolls")
	enrollsGroup.Use(middleware.Auth())
	{
		enrollsGroup.POST("/", enrolls.Create)
		enrollsGroup.PUT("/:enroll_id", enrolls.Update)
		enrollsGroup.DELETE("/:enroll_id", enrolls.Delete)
	}

	cohortsGroup := router.Group("/cohorts")
	cohortsGroup.Use(middleware.Auth())
	{
		cohortsGroup.POST("/", cohorts.Create)
		cohortsGroup.PUT("/:cohort_id", cohorts.Update)
		cohortsGroup.DELETE("/:cohort_id", cohorts.Delete)
	}

	logger.Info("start the application...")
	router.Run(":8010")
}
