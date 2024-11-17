package handler

import (
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	r "github.com/dilshodforever/nasiya-savdo/storage/redis"
)

type Handler struct {
	User  pb.UserServiceClient
	redis r.InMemoryStorageI
}

func NewHandler(us pb.UserServiceClient, rdb r.InMemoryStorageI) *Handler {
	return &Handler{us, rdb}
}
