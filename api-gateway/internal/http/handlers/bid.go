package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/dilshodforever/tender/internal/http/middleware"
	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// SubmitBid handles creating a new bid
// @Summary      Submit Bid
// @Description  Submit a bid for a tender
// @Tags         Bid
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Tender ID"
// @Param        bid body pb.SubmitBidRequest true "Bid details"
// @Success      200 {object} pb.BidResponse "Bid submitted successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while submitting bid"
// @Router       /tenders/{id}/bids [post]
func (h *HTTPHandler) SubmitBid(ctx *gin.Context) {
	middleware.RateLimitMiddleware(h.Redis)
	var req pb.SubmitBidRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	req.TenderId = ctx.Param("id")
	id := middleware.GetUserId(ctx)
	req.ContractorId = id

	// Calling the service to submit the bid
	res, err := h.BidService.SubmitBid(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// ListBids handles listing bids with optional filters
// @Summary      List Bids
// @Description  List all bids with optional filtering by price, delivery time, limit, and offset
// @Tags         Bid
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        price query float64 false "Price filter"
// @Param        delivery_time query int32 false "Delivery time filter"
// @Param        limit query int32 false "Limit the number of bids"
// @Param        offset query int32 false "Offset for pagination"
// @Success      200 {object} pb.ListBidsResponse "Bids retrieved successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while retrieving bids"
// @Router       /bid/list [get]
func (h *HTTPHandler) ListBids(ctx *gin.Context) {
	// Parsing query parameters
	price, err := strconv.ParseFloat(ctx.DefaultQuery("price", "0"), 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid price parameter"})
		return
	}

	deliveryTime, err := strconv.Atoi(ctx.DefaultQuery("delivery_time", "0"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid delivery_time parameter"})
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Create the request struct
	req := pb.ListBidsRequest{
		Price:        float32(price),
		DeliveryTime: int32(deliveryTime),
		Limit:        int32(limit),
		Offset:       int32(offset),
	}

	// Generate a cache key based on the presence of filter parameters
	var cacheKey string
	if price == 0 && deliveryTime == 0 {
		// No filters applied, generic cache key
		cacheKey = "bids:no_filter"
	} else {
		// Filters applied, include them in the cache key
		cacheKey = fmt.Sprintf("bids:price:%f:delivery_time:%d:limit:%d:offset:%d", price, deliveryTime, limit, offset)
	}

	// Try to get the cached data from Redis
	cachedData, err := h.Redis.Get(ctx, cacheKey).Result()
	if err != nil && err != redis.Nil {
		// If Redis connection fails or another error occurs, log it
		slog.Error("Failed to get data from Redis: %v", err)
	}

	var res *pb.ListBidsResponse

	if err == nil {
		// Data found in Redis, unmarshal it
		fmt.Println("Cache hit (Redis)")

		// If cached data is found, unmarshal it
		if err := json.Unmarshal([]byte(cachedData), &res); err != nil {
			// Error while parsing the cached data
			ctx.JSON(500, gin.H{"error": "Failed to parse cached data"})
			return
		}
	} else if err == redis.Nil {
		// Data not found in Redis, fetch it from the service
		fmt.Println("Cache miss (postgres)")

		// Call the service to fetch data from MongoDB or another source
		res, err = h.BidService.ListBids(ctx, &req)
		if err != nil {
			// Error while fetching the data from the service
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Marshal the response before caching it
		data, err := json.Marshal(res)
		if err != nil {
			// Error while marshaling the response
			slog.Error("Failed to marshal response: %v", err)
		} else {
			// Store the data in Redis with a 30-minute expiration
			err := h.Redis.Set(ctx, cacheKey, string(data), 30*time.Minute).Err()
			if err != nil {
				// Error while setting data in Redis
				slog.Error("Failed to set data in Redis: %v", err)
			}
		}
	} else {
		// In case of any other Redis errors
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Return the response (from cache or service)
	ctx.JSON(200, res)
}

// GetAllBidsByTenderId handles fetching all bids by tender ID
// @Summary      Get All Bids by Tender ID
// @Description  Get all bids associated with a specific tender
// @Tags         Bid
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Tender ID"
// @Success      200 {object} pb.ListBidsResponse "Bids retrieved successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while retrieving bids"
// @Router       /tenders/{id}/bids [get]
func (h *HTTPHandler) GetAllBidsByTenderId(ctx *gin.Context) {
	id := ctx.Param("id")
	// Calling the service to get all bids by tender ID
	res, err := h.BidService.GetAllBidsByTenderId(ctx, &pb.GetAllByid{Id: id})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// ListContractorBids handles listing bids by contractor ID
// @Summary      List Contractor Bids
// @Description  List all bids placed by a specific contractor
// @Tags         Bid
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Contractor ID"
// @Success      200 {object} pb.GetAllBidsByUserIdRequest "Contractor's bids retrieved successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while retrieving bids"
// @Router       /users/{id}/bids [get]
func (h *HTTPHandler) ListContractorBids(ctx *gin.Context) {
	id := ctx.Param("id")

	// Calling the service to get all bids by contractor ID
	res, err := h.BidService.ListContractorBids(ctx, &pb.GetAllByid{Id: id})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}
