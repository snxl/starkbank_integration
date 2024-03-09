package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynqmon"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/snxl/stark_bank_integration/src/application/docs"
	healthcheck "github.com/snxl/stark_bank_integration/src/application/handler/health_check"
	invoiceevent "github.com/snxl/stark_bank_integration/src/application/handler/invoice_event"
	"github.com/snxl/stark_bank_integration/src/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of requests",
		},
		[]string{"method"},
	)
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(addr string) error {

	docs.SwaggerInfo.Title = "Service docs"
	docs.SwaggerInfo.Description = "service documentation"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	v1 := r.Group("/v1")
	{
		webhook := v1.Group("/webhook")
		{
			webhook.POST("/starkbank", invoiceevent.NewInvoicePaidEventHandler().Run)
		}
	}

	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitor",
		RedisConnOpt: config.GetAsynq().RedisOpt,
	})

	prometheus.MustRegister(httpRequestsTotal)
	r.Any("/monitor/*a", gin.WrapH(h))
	r.GET("/healthcheck", healthcheck.NewHealthCheckHandler().Run)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(recoveryMiddleware())
	return r.Run(fmt.Sprintf(":%s", addr))
}

func recoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(500, gin.H{"message": "Erro interno"})
				c.Abort()
			}
		}()
		c.Next()
	}
}
