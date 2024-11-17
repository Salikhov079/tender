package handler

import (
	"strconv"
	"time"

	token "github.com/dilshodforever/nasiya-savdo/api/token"
	"github.com/dilshodforever/nasiya-savdo/config"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

// RegisterUser handles the creation of a new user
// @Summary Register User
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Create body pb.User true "Create"
// @Success 201 {string} string "Create Successful"
// @Failure 400 {string} string "Validation Error"
// @Failure 409 {string} string "Duplicate Email"
// @Router /register [post]
func (h *Handler) Register(ctx *gin.Context) {
	user := &pb.User{}

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "invalid input data"})
		return
	}

	if user.Username == "" || user.Email == "" {
		ctx.JSON(400, gin.H{"message": "username or email cannot be empty"})
		return
	}

	if !IsEmailValid(user.Email) {
		ctx.JSON(400, gin.H{"message": "invalid email format"})
		return
	}

	if user.Role != "client" && user.Role != "admin" && user.Role != "contractor" {
		ctx.JSON(400, gin.H{"message": "invalid role"})
		return
	}

	users, err := h.User.GetAll(ctx, &pb.UserFilter{})
	if err != nil {
		ctx.JSON(500, gin.H{"message": "failed to retrieve users"})
		return
	}

	if !IsEmailUniq(users, user.Email) {
		ctx.JSON(400, gin.H{"message": "Email already exists"})
		return
	}

	if !IsUserNameUniq(users, user.Username) {
		ctx.JSON(400, gin.H{"message": "Username already exists"})
		return
	}

	user, err = h.User.Register(ctx, user)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "failed to register user"})
		return
	}

	t := token.GenerateJWTToken(user)
	h.redis.Set(user.Id, t.RefreshToken, 30*24*time.Hour)

	ctx.JSON(201, gin.H{
		"token":   t,
		"message": "User registered successfully",
	})
}

// UpdateUser handles updating an existing user
// @Summary Update User
// @Description Update an existing user
// @Tags Admin Managment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Param Update body pb.User true "Update"
// @Success 200 {string} string "Update Successful"
// @Failure 400 {string} string "Error while updating user"
// @Router /user/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	user := pb.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err = h.User.Update(ctx, &user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success!",
	})
}

// DeleteUser handles the deletion of a user
// @Summary Delete User
// @Description Delete an existing user
// @Tags Admin Managment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "Delete Successful"
// @Failure 400 {string} string "Error while deleting user"
// @Router /user/delete/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}

	_, err := h.User.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success!",
	})
}

// GetAllUser handles retrieving all users
// @Summary Get All Users
// @Description Get all users
// @Tags Admin Managment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query pb.UserFilter true "Query parameter"
// @Success 200 {object} pb.AllUsers "Get All Successful"
// @Failure 400 {string} string "Error while retrieving users"
// @Router /user/getall [get]
func (h *Handler) GetAllUser(ctx *gin.Context) {
	filter := pb.UserFilter{
		Limit:    parseQueryInt32(ctx, "limit", 10), // Default limit 10
		Offset:   parseQueryInt32(ctx, "offset", 0), // Default offset 0
		Username: ctx.Query("full_name"),
		Email:    ctx.Query("email"),
	}

	res, err := h.User.GetAll(ctx, &filter)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}

func parseQueryInt32(ctx *gin.Context, key string, defaultValue int32) int32 {
	valueStr := ctx.Query(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return int32(value)
}

// GetByIdUser handles retrieving a user by ID
// @Summary Get User By ID
// @Description Get a user by ID
// @Tags Admin Managment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} pb.User "Get By ID Successful"
// @Failure 400 {string} string "Error while retrieving user"
// @Router /user/getbyid/{id} [get]
func (h *Handler) GetbyIdUser(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}

	res, err := h.User.GetById(ctx, &id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}

// LoginUser handles user login
// @Summary Login User
// @Description Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Create body pb.UserLogin true "Login"
// @Success 200 {object} string "Login Successful"
// @Failure 401 {string} string "Invalid username or password"
// @Failure 400 {string} string "Validation Error"
// @Router /login [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	user := pb.UserLogin{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid input data"})
		return
	}

	if user.Username == "" || user.Password == "" {
		ctx.JSON(400, gin.H{"message": "Username and password are required"})
		return
	}

	res, err := h.User.Login(ctx, &pb.UserLogin{Username: user.Username, Password: user.Password})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = invalid username or password" {
			ctx.JSON(401, gin.H{"message": "Invalid username or password"})
			return
		}
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}

	t := token.GenerateJWTToken(res)
	h.redis.Set(res.Id, t.RefreshToken, 30*24*time.Hour)

	ctx.JSON(200, gin.H{
		"token":   t,
		"message": "Login successful",
	})
}

// GetProfil handles retrieving a user Profil
// @Summary Get User Profil
// @Description Get a user Profil
// @Tags User Managment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.User "Get Profil Successful"
// @Failure 400 {string} string "Error while retrieving user"
// @Router /user/get_profil [get]
func (h *Handler) GetProfil(ctx *gin.Context) {
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)

	res, err := h.User.GetById(ctx, &pb.ById{Id: id})
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}

// UpdateUser handles updating an existing user
// @Summary Update Profil
// @Description Update an existing user
// @Tags User Managment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Update body pb.User true "Update"
// @Success 200 {string} string "Update Successful"
// @Failure 400 {string} string "Error while updating user"
// @Router /user/update_profil [put]
func (h *Handler) UpdateProfil(ctx *gin.Context) {
	user := pb.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	cnf := config.Load()
	user.Id, _ = token.GetIdFromToken(ctx.Request, &cnf)
	_, err = h.User.Update(ctx, &user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Updated successfuly!",
	})
}

// DeleteProfil handles the deletion of a Profil
// @Summary Delete Profil
// @Description Delete an existing Profil
// @Tags User Managment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {string} string "Delete Successful"
// @Failure 400 {string} string "Error while deleting user"
// @Router /user/delete_profil [delete]
func (h *Handler) DeleteProfil(ctx *gin.Context) {
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)

	_, err := h.User.Delete(ctx, &pb.ById{Id: id})
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Deleted successfuly!",
	})
}

// RefreshToekn handles the deletion of a Token
// @Summary refresh Toekn
// @Description refresh an existing Token
// @Tags User Managment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {string} string "refresh Successful"
// @Failure 400 {string} string "Error while refreshed token"
// @Router /user/refresh-token [get]
func (h *Handler) RefreshToken(ctx *gin.Context) {
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)
	tok, err := h.redis.Get(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err = token.ExtractClaim(&cnf, tok)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.User.GetById(ctx, &pb.ById{Id: id})
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	access := token.GenerateJWTToken(user).AccessToken

	ctx.Request.Header.Set("Authorization", access)

	ctx.JSON(200, gin.H{"token: ": access})
}
