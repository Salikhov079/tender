package api

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/dilshodforever/nasiya-savdo/api/handler"
	"github.com/dilshodforever/nasiya-savdo/api/middleware"
	_ "github.com/dilshodforever/nasiya-savdo/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title auth service API
// @version 1.0
// @description created by salikhov
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy: ", err)
		panic(err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(middleware.NewAuth(e))
	
	{
		r.POST("/register", h.Register)
		r.POST("/login", h.LoginUser)
	}
	
	u := r.Group("/user")
	{
		u.PUT("/update/:id", h.UpdateUser)
		u.DELETE("/delete/:id", h.DeleteUser)
		u.GET("/getall", h.GetAllUser)
		u.GET("/getbyid/:id", h.GetbyIdUser)
	}

	{
		u.GET("/get_profil", h.GetProfil)
		u.PUT("/update_profil", h.UpdateProfil)
		u.DELETE("/delete_profil", h.DeleteProfil)
	}

	{
		u.GET("/refresh-token", h.RefreshToken)
	}

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	return r
}
