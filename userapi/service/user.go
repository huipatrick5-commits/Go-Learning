package service

import (
    "userapi/repository"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (repository.User, error) {
    return s.repo.GetByID(id)
}

func (s *UserService) AddUser(user repository.User) error {
    return s.repo.AddUser(user)
}

func (s *UserService) DelUser(id int) error {
    return s.repo.DelByID(id)
}

func (s *UserService) ModUser(user repository.User) error {
    return s.repo.ModUser(user)
}
