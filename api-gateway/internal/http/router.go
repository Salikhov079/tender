package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dilshodforever/tender/internal/http/handlers"
	_ "github.com/dilshodforever/tender/internal/http/docs"

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

	// Tender Endpoints
tender := r.Group("/tenders")
{
    tender.POST("/", h.CreateTender)
    tender.POST("/:id/award/:bid_id", h.TenderAward)
    tender.DELETE("/:id", h.DeleteTender)
    tender.GET("/", h.ListTenders)
    tender.PUT("/:id", h.UpdateTender)
}

// User Tender Endpoints
user := r.Group("/users")
{
    user.GET("/:id/tenders", h.ListUserTenders)
    user.GET("/:id/bids", h.ListContractorBids)
}

// Tender Bid Endpoints
bids := r.Group("/tenders/:id/bids")
{
    bids.POST("/", h.SubmitBid)
    bids.GET("/", h.GetAllBidsByTenderId)
}

// General Bid Endpoints
bid := r.Group("/bid")
{
    bid.GET("/list", h.ListBids)
}


	return r
}
