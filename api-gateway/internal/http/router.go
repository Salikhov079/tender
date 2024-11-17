package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/dilshodforever/tender/internal/http/docs"
	"github.com/dilshodforever/tender/internal/http/handlers"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description Dilshod's API Gateway
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handlers.HTTPHandler) *gin.Engine {
	r := gin.Default()

	// Middleware setup
	ca, err := casbin.NewEnforcer("internal/pkg/config/model.conf", "internal/pkg/config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// r.Use(middleware.NewAuth(ca))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Authentication"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Swagger documentation
	url := ginSwagger.URL("/swagger/doc.json") // Adjusted path for Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	rt := r.Group("/api")
	// Tender Endpoints
	tender := rt.Group("/client/tenders")
	{
		tender.POST("", h.CreateTender)
		tender.POST("/:id/award/:bid_id", h.TenderAward)
		tender.DELETE("/:id", h.DeleteTender)
		tender.GET("", h.ListTenders)
		tender.PUT("/:id", h.UpdateTender)
	}

	// User Tender Endpoints
	user := rt.Group("/users")
	{
		user.GET("/:id/tenders", h.ListUserTenders)
		user.GET("/:id/bids", h.ListContractorBids)
	}

	// Tender Bid Endpoints
	bids := rt.Group("/tenders/:id/bids")
	{
		bids.POST("", h.SubmitBid)
		bids.GET("", h.GetAllBidsByTenderId)
	}

	// General Bid Endpoints
	bid := rt.Group("/bid")
	{
		bid.GET("/list", h.ListBids)
	}

	return r
}
