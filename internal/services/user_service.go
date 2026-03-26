package services

import (
	"errors"
	"log"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/repository"
	"github.com/ViitoJooj/clown-crm/pkg/cryptography"
	"github.com/ViitoJooj/clown-crm/pkg/jwtTokens"
	"github.com/golang-jwt/jwt/v4"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(user *domain.User) (*domain.User, error) {
	existing, err := s.repo.FindUserByEmail(user.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}
	if existing != nil {
		log.Println("User already exists")
		return nil, errors.New("invalid credentials")
	}

	newUser, err := domain.NewUser(
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *UserService) Login(email, password string) (*domain.User, string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		log.Println(err)
		return nil, "", errors.New("internal error")
	}
	if user == nil {
		log.Println("User not found")
		return nil, "", errors.New("invalid credentials")
	}

	if !cryptography.CheckPasswordHash(password, user.Password) {
		log.Println("Invalid password")
		return nil, "", errors.New("invalid credentials")
	}

	token, err := jwtTokens.GenerateToken(user.UUID)
	if err != nil {
		log.Println(err)
		return nil, "", errors.New("internal error")
	}

	return user, token, nil
}

func (s *UserService) AccessToken(tokenString string) (*domain.User, error) {
	token, err := jwtTokens.ValidateToken(tokenString)
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Println("Invalid token claims")
		return nil, errors.New("invalid token")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		log.Println("Invalid user ID in token claims")
		return nil, errors.New("invalid token")
	}
	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}
	if user == nil {
		log.Println("User not found for token")
		return nil, errors.New("invalid token")
	}
	return user, nil
}

func (s *UserService) ViewUser(user *domain.User) (*domain.User, error) {
	return s.repo.FindUserByID(user.UUID)
}

func (s *UserService) DeleteUser(user *domain.User) error {
	return s.repo.DeleteUserById(user.UUID)
}
