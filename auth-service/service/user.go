package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	s "github.com/dilshodforever/nasiya-savdo/storage"
)

type UserService struct {
	stg s.InitRoot
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg s.InitRoot) *UserService {
	return &UserService{stg: stg}
}

func (c *UserService) Register(ctx context.Context, userR *pb.User) (*pb.User, error) {
	user, err := c.stg.User().Register(userR)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (c *UserService) Login(ctx context.Context, login *pb.UserLogin) (*pb.User, error) {
	user, err := c.stg.User().Login(login)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (c *UserService) GetById(ctx context.Context, id *pb.ById) (*pb.User, error) {
	user, err := c.stg.User().GetById(id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (c *UserService) GetAll(ctx context.Context, filter *pb.UserFilter) (*pb.AllUsers, error) {
	users, err := c.stg.User().GetAll(filter)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return users, nil
}

func (c *UserService) Update(ctx context.Context, user *pb.User) (*pb.User, error) {
	user, err := c.stg.User().Update(user)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (c *UserService) Delete(ctx context.Context, id *pb.ById) (*pb.User, error) {
	user, err := c.stg.User().Delete(id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}
