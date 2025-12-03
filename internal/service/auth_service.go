package service

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"senior-fullstack-2025/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	Redis     *redis.Client
	JWTSecret string
}

func NewAuthService(redisClient *redis.Client, jwtSecret string) *AuthService {
	return &AuthService{
		Redis:     redisClient,
		JWTSecret: jwtSecret,
	}
}

// Method for user registration
func (s *AuthService) RegisterUser(username, realname, email, password string) error {
	switch {
	case username == "":
		return fmt.Errorf("username is required")
	case realname == "":
		return fmt.Errorf("realname is required")
	case email == "":
		return fmt.Errorf("email is required")
	case password == "":
		return fmt.Errorf("password is required")

	default:
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*]{8,}$`)
	if !passwordRegex.MatchString(password) {
		return fmt.Errorf("password must be at least 8 characters long and contain only alphanumeric or special characters")
	}

	hashedPassword := hashSHA1(password)

	value := model.User{
		RealName: realname,
		Email:    email,
		Password: hashedPassword,
	}
	keyFormat := fmt.Sprintf("login_%s", username)

	userJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = s.Redis.Set(context.Background(), keyFormat, userJSON, 0).Err()

	if err != nil {
		return err

	}

	return nil
}

// Method for user login
func (s *AuthService) LoginUser(username, password string) (string, error) {
	switch {
	case username == "":
		return "", fmt.Errorf("username is required")
	case password == "":
		return "", fmt.Errorf("password is required")
	default:
	}

	hashedPassword := hashSHA1(password)

	keyFormat := fmt.Sprintf("login_%s", username)
	val, err := s.Redis.Get(context.Background(), keyFormat).Result()

	if err == redis.Nil {
		return "", errors.New("user not found")
	} else if err != nil {
		return "", err
	}

	var user model.User
	err = json.Unmarshal([]byte(val), &user)

	if err != nil {
		return "", errors.New("failed to parse user data")
	}

	if user.Password != hashedPassword {
		return "", errors.New("invalid password")
	}

	// Generate JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Expiration: 72 hours
	})

	// Sign and get the complete encoded token as a string using the secret
	token, err := claims.SignedString([]byte(s.JWTSecret))

	return token, err
}

// Helper for sha1 hash
func hashSHA1(input string) string {
	hash := sha1.New()
	hash.Write([]byte(input))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
