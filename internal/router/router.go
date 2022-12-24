package router

import (
	"github.com/gin-gonic/gin"
	"sleep-tracker/external"
	"sleep-tracker/internal/handler"
	"sleep-tracker/internal/middleware"
	"sleep-tracker/internal/repository"
	"sleep-tracker/internal/service"
)

type Router struct {
	Configs     *external.Configs
	DataSources *external.DataSources
}

func NewRouter(configs *external.Configs, dataSources *external.DataSources) *Router {
	return &Router{
		Configs:     configs,
		DataSources: dataSources,
	}
}

func (r *Router) Init() (engine *gin.Engine, err error) {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	userRepository := repository.NewUserRepository(r.DataSources.PostgreSQL)
	sleepRepository := repository.NewSleepRepository(r.DataSources.PostgreSQL)
	userService := service.NewUserService(&service.UserServiceConfig{
		UserRepository: userRepository,
	})
	sleepService := service.NewSleepService(&service.SleepServiceConfig{
		SleepRepository: sleepRepository,
		UserRepository:  userRepository,
	})
	userHandler := handler.NewUserHandler(&handler.UserHandlerConfig{
		UserService:  userService,
		SleepService: sleepService,
		MaxBodyBytes: r.Configs.MaxBodyBytes,
	})

	adminRepository := repository.NewUserRepository(r.DataSources.PostgreSQL)
	adminService := service.NewAdminService(&service.AdminServiceConfig{
		AdminRepository: adminRepository,
		UserRepository:  userRepository,
	})
	adminHandler := handler.NewAdminHandler(&handler.AdminHandlerConfig{
		AdminService: adminService,
		UserService:  userService,
		MaxBodyBytes: r.Configs.MaxBodyBytes,
	})

	routerGroup := router.Group(r.Configs.URLS.APIURL)

	ug := routerGroup.Group(r.Configs.URLS.UserURL)
	ug.POST("signup", userHandler.Signup)
	ug.POST("signin", userHandler.Signin)
	ug.POST("sleep", userHandler.CreateSleep)
	ug.DELETE("sleep", userHandler.DeleteSleep)
	ug.PATCH("sleep", userHandler.UpdateSleep)

	ag := routerGroup.Group(r.Configs.URLS.AdminURL)
	ag.DELETE("", adminHandler.Delete)
	ag.PATCH("", adminHandler.Update)
	ag.GET("", adminHandler.GetUsers)

	return router, nil
}
