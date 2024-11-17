package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"log"

	"github.com/casbin/casbin/v2"
	"github.com/dilshodforever/tender/internal/http/token"
	"github.com/dilshodforever/tender/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

type JwtRoleAuth struct {
	enforcer   *casbin.Enforcer
	jwtHandler token.JWTHandler
}

func NewAuth(enforce *casbin.Enforcer) gin.HandlerFunc {
	auth := JwtRoleAuth{
		enforcer: enforce,
	}
	return func(ctx *gin.Context) {
		allow, err := auth.CheckPermission(ctx)
		if err != nil {
			valid, _ := err.(jwt.ValidationError)
			if valid.Errors == jwt.ValidationErrorExpired {
				ctx.AbortWithStatusJSON(http.StatusForbidden, "Invalid token !!!")

			} else {
				ctx.AbortWithStatusJSON(401, "Access token expired")
			}
		} else if !allow {
			ctx.AbortWithStatusJSON(http.StatusForbidden, "Permission denied")

		}
	}

}

func (a *JwtRoleAuth) GetRole(r *gin.Context) (string, error) {
	var (
		claims jwt.MapClaims
		err    error
	)
	jwtToken := r.Request.Header.Get("Authorization")

	if jwtToken == "" {
		return "unauthorized", nil
	} else if strings.Contains(jwtToken, "Basic") {
		return "unauthorized", nil
	}
	a.jwtHandler.Token = jwtToken
	a.jwtHandler.SigningKey = config.Load().TokenKey
	claims, err = a.jwtHandler.ExtractClaims()

	if err != nil {
		log.Println("Error while extracting claims: ", err)
		return "unauthorized", err
	}
	fmt.Println("role: ", claims["role"])
	return claims["role"].(string), nil
}

func (a *JwtRoleAuth) CheckPermission(r *gin.Context) (bool, error) {
	role, err := a.GetRole(r)
	if err != nil {
		log.Println("Error while getting role from token: ", err)
		return false, err
	}
	method := r.Request.Method
	path := r.FullPath()
	allowed, err := a.enforcer.Enforce(role, path, method)

	if err != nil {
		log.Println("Error while comparing role from csv list: ", err)
		return false, err
	}

	return allowed, nil
}

func GetUserId(r *gin.Context) string {
	jwtToken := r.Request.Header.Get("Authorization")

	claims, err := token.ExtractClaim(jwtToken)
	if err != nil {
		panic(err)
	}
	id := claims["id"].(string)
	return id
}

func GetStorageId(r *gin.Context) string {
	jwtToken := r.Request.Header.Get("Authorization")

	claims, err := token.ExtractClaim(jwtToken)
	if err != nil {
		panic(err)
	}
	id := claims["storage_id"].(string)
	return id
}


// RateLimitMiddleware limits requests to 5 per minute per contractor
func RateLimitMiddleware(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id") // Foydalanuvchi ID sini olish
		tenderID := c.Param("id")        // Tender ID sini olish
		key := "rate_limit:" + tenderID + ":" + userID

		// Redis dan so'rovlar sonini tekshirish
		val, err := client.Get(context.Background(), key).Result()
		if err == redis.Nil {
			// Redisda ma'lumot mavjud emas bo'lsa, yangisini qo'shish
			client.Set(context.Background(), key, "1", time.Minute)
		} else if err != nil {
			// Redis bilan bog'lanishda xatolik
			c.JSON(500, gin.H{"error": "Rate limit check failed"})
			c.Abort()
			return
		} else {
			count, _ := strconv.Atoi(val)
			if count >= 5 {
				// Agar limitga yetilgan bo'lsa, so'rovni rad etish
				c.JSON(429, gin.H{"error": "Rate limit exceeded"})
				c.Abort()
				return
			}
			// Aks holda, sanashni davom ettirish
			client.Incr(context.Background(), key)
		}

		// So'rovni davom ettirish
		c.Next()
	}
}
