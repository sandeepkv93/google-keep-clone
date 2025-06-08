package services

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "google-keep-clone/internal/config"
    "google-keep-clone/internal/models"
    "google-keep-clone/internal/repositories"
    "google-keep-clone/internal/validators"
)

type AuthService struct {
    userRepo *repositories.UserRepository
    config   *config.Config
}

type Claims struct {
    UserID string `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

func NewAuthService(userRepo *repositories.UserRepository, config *config.Config) *AuthService {
    return &AuthService{
        userRepo: userRepo,
        config:   config,
    }
}

func (s *AuthService) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func (s *AuthService) CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (s *AuthService) GenerateToken(user *models.User) (string, error) {
    claims := &Claims{
        UserID: user.ID.String(),
        Email:  user.Email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.config.JWTSecret), nil
    })

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    return nil, err
}

func (s *AuthService) Register(req *validators.RegisterRequest) (*models.User, string, error) {
    // Check if user already exists
    existingUser, _ := s.userRepo.GetByEmail(req.Email)
    if existingUser != nil {
        return nil, "", errors.New("user already exists")
    }

    // Hash password
    hashedPassword, err := s.HashPassword(req.Password)
    if err != nil {
        return nil, "", errors.New("failed to hash password")
    }

    // Create user
    user := &models.User{
        Email:    req.Email,
        Password: hashedPassword,
        Name:     req.Name,
        Provider: "local",
    }

    if err := s.userRepo.Create(user); err != nil {
        return nil, "", errors.New("failed to create user")
    }

    // Generate token
    token, err := s.GenerateToken(user)
    if err != nil {
        return nil, "", errors.New("failed to generate token")
    }

    return user, token, nil
}

func (s *AuthService) Login(req *validators.LoginRequest) (*models.User, string, error) {
    // Get user by email
    user, err := s.userRepo.GetByEmail(req.Email)
    if err != nil {
        return nil, "", errors.New("invalid credentials")
    }

    // Check password
    if !s.CheckPasswordHash(req.Password, user.Password) {
        return nil, "", errors.New("invalid credentials")
    }

    // Generate token
    token, err := s.GenerateToken(user)
    if err != nil {
        return nil, "", errors.New("failed to generate token")
    }

    return user, token, nil
}