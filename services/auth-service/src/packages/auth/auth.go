package auth

import (
	"errors"
	"log"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/xams-backend/services/auth-service/src/internal/models"
	"github.com/xams-backend/services/auth-service/src/internal/users/usecase"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	JWT_SECRET_ENV = "JWT_SECRET"
	JWT_EXPIRATION = 24 * time.Hour
	SALTY = 14
)

type Authentication struct {
	userUsecase usecase.IUserUsecase
	jwtSecret   []byte
}

func NewAuthentication(uc usecase.IUserUsecase) IAuthentication {	
	jwtSecret := os.Getenv(JWT_SECRET_ENV)
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	return &Authentication{
		userUsecase: uc,
		jwtSecret:   []byte(jwtSecret),
	}
}

func (auth *Authentication) generateJWT(user *models.User) (string, error) {
	claims := &models.UserClaims{
		UserID: user.UserID,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRATION)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token_issuer",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	ss, err := token.SignedString(auth.jwtSecret)
	if err != nil {
		log.Printf("Failed to sign JWT token: %v", err)
		return "", errors.New("failed to sign JWT token")
	}

	return ss, nil
}

func (auth *Authentication) Login(email string, password string) (*models.UserResponse, error) {
	user, err := auth.userUsecase.GetUserByEmail(email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Password mismatch: %v", err)
		return nil, errors.New("invalid password")
	}

	token, err := auth.generateJWT(user)
	if err != nil {
		return nil, err
	}

	userResponse := models.UserResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		Prename:   user.Prename,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BranchID:  user.BranchID,
		Role:      user.Role,
		Token:     token,
	}

	log.Printf("User %s logged in successfully", user.UserID)
	return &userResponse, nil
}

func (auth *Authentication) Register(user *models.User) (*models.UserResponse, error) {
	if err := auth.userUsecase.IsUserAlreadyExists(user.UserID); errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("User with ID %s already exists", user.UserID)
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), SALTY)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, errors.New("error hashing password")
	}
	user.Password = string(hashedPassword)

	if err := auth.userUsecase.CreateUser(user); err != nil {
		log.Printf("Cannot create user: %v", err)
		return nil, errors.New("cannot create user")
	}

	createdUser, err := auth.userUsecase.GetUser(user.UserID)
	if err != nil {
		log.Printf("Cannot find created user: %v", err)
		return nil, errors.New("cannot find created user")
	}

	token, err := auth.generateJWT(createdUser)
	if err != nil {
		return nil, err
	}

	userResponse := models.UserResponse{
		UserID:    createdUser.UserID,
		Email:     createdUser.Email,
		Prename:   createdUser.Prename,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		BranchID:  createdUser.BranchID,
		Role:      createdUser.Role,
		Token:     token,
	}

	log.Printf("User %s registered successfully", createdUser.UserID)
	return &userResponse, nil
}
