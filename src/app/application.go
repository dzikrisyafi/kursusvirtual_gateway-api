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
	router.Use(cors.Default())
	enrolls := http.NewEnrollsHandler(enrolls.NewService(rest.NewRestUsersRepository(), rest.NewRestCoursesRepository()))
	cohorts := http.NewCohortHandler(cohort.NewService(rest.NewRestUsersRepository(), rest.NewRestCoursesRepository()))

	enrollsGroup := router.Group("/enrolls")
	enrollsGroup.Use(middleware.Auth())
	{
		router.POST("/", enrolls.Create)
		router.PUT("/:enroll_id", enrolls.Update)
		router.DELETE("/:enroll_id", enrolls.Delete)
	}

	cohortsGroup := router.Group("/cohorts")
	cohortsGroup.Use(middleware.Auth())
	{
		router.POST("/", cohorts.Create)
		router.PUT("/:cohort_id", cohorts.Update)
		router.DELETE("/:cohort_id", cohorts.Delete)
	}

	logger.Info("start the application...")
	router.Run(":8010")
}
