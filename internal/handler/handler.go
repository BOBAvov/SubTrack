package handler

import (
	"github.com/BOBAvov/sub_track/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"log/slog"
)

type Handler struct {
	services *service.Service
	log      *slog.Logger
}

func NewHandler(services *service.Service, logger *slog.Logger) *Handler {
	return &Handler{services: services, log: logger}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("isDateValid", isDateValid); err != nil {
			log.Fatalf("isDateValid validator registered %v", err)
		}
	} else {
		log.Fatalf("Vlidatetor error!")
	}

	api := router.Group("/api")
	{
		subs := api.Group("/subs")
		{
			subs.POST("/", h.createSubscription)
			subs.GET("/", h.getAllSubscriptions)
			subs.GET("/:id", h.getByIdSubscription)
			subs.PUT("/:id", h.updateSubscription)
			subs.DELETE("/:id", h.deleteSubscription)
		}
		total := api.Group("/total")
		{
			total.POST("/", h.totalSum)
		}
	}
	return router
}
