package handlers

import (
	"net/http"
	"strconv"

	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateTender godoc
// @Summary Create a new tender
// @Description Create a new tender with the provided details
// @Tags 04-Tender
// @Accept json
// @Produce json
// @Param request body pb.CreateTenderRequest true "Tender details"
// @Success 200 {object} pb.TenderResponse "Tender created successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Server error"
// @Router /tenders [POST]
func (h *HTTPHandler) CreateTender(c *gin.Context) {
	var req pb.CreateTenderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.TenderService.CreateTender(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// TenderAward godoc
// @Summary Award a tender
// @Description Award a tender to a specific bid
// @Tags Tender
// @Accept json
// @Produce json
// @Param id path string true "Tender ID"
// @Param bid_id path string true "Bid ID"
// @Success 200 {object} pb.TenderResponse "Tender awarded successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Tender or bid not found"
// @Failure 500 {object} string "Server error"
// @Router /tenders/{id}/award/{bid_id} [POST]
func (h *HTTPHandler) TenderAward(c *gin.Context) {
	var req pb.CreatTenderAwardRequest
	req.TenderId=c.Param("id")
	req.BidId=c.Param("bid_id")
	resp, err := h.TenderService.TenderAward(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteTender godoc
// @Summary Delete a tender
// @Description Delete a tender by its ID
// @Tags 04-Tender
// @Accept json
// @Produce json
// @Param id path string true "Tender ID"
// @Success 200 {object} pb.TenderResponse "Tender deleted successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Tender not found"
// @Failure 500 {object} string "Server error"
// @Router /tenders/{id} [DELETE]
func (h *HTTPHandler) DeleteTender(c *gin.Context) {
	tenderID := c.Param("id")

	resp, err := h.TenderService.DeleteTender(c, &pb.TenderIdRequest{Id: tenderID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListTenders godoc
// @Summary List tenders
// @Description List tenders with optional filtering by title and deadline
// @Tags Tender
// @Accept json
// @Produce json
// @Param title query string false "Filter by title"
// @Param deadline query string false "Filter by deadline"
// @Param limit query int false "Number of items to retrieve"
// @Param offset query int false "Starting position of items to retrieve"
// @Success 200 {object} pb.ListTendersResponse "List of tenders"
// @Failure 500 {object} string "Server error"
// @Router /tenders [GET]
func (h *HTTPHandler) ListTenders(c *gin.Context) {
	title := c.Query("title")
	deadline := c.Query("deadline")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Prepare request object
	req := &pb.ListTendersRequest{
		Title:    title,
		Deadline: deadline,
		Limit:    int32(limit),
		Offset:   int32(offset),
	}

	// Call the gRPC service
	resp, err := h.TenderService.ListTenders(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateTender godoc
// @Summary Update a tender
// @Description Update an existing tender with the provided details
// @Tags 04-Tender
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param request body pb.UpdateTenderRequest true "Updated tender details"
// @Success 200 {object} pb.TenderResponse "Tender updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Tender not found"
// @Failure 500 {object} string "Server error"
// @Router /tenders/{id} [PUT]
func (h *HTTPHandler) UpdateTender(c *gin.Context) {
	var req pb.UpdateTenderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id=c.Param("id")
	resp, err := h.TenderService.UpdateTender(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListUserTenders godoc
// @Summary List tenders for a specific user
// @Description List all tenders created by a specific user
// @Tags 04-Tender
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} pb.ListTendersResponse "List of user tenders"
// @Failure 500 {object} string "Server error"
// @Router /users/{id}/tenders [GET]
func (h *HTTPHandler) ListUserTenders(c *gin.Context) {
	userID := c.Param("id")
	resp, err := h.TenderService.ListUserTenders(c, &pb.TenderIdRequest{Id: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
