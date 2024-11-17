package postgres

import (
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
)

type InitRoot interface {
	User() User
}

type User interface {
	Register(user *pb.User) (*pb.User, error)
	Login(user *pb.UserLogin) (*pb.User, error)
	GetById(id *pb.ById) (*pb.User, error)
	GetAll(filter *pb.UserFilter) (*pb.AllUsers, error)
	Update(user *pb.User) (*pb.User, error)
	Delete(id *pb.ById) (*pb.User, error)
}
