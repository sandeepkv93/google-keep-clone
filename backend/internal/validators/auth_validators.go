package validators

import (
    "errors"
    "regexp"
    "strings"
)

type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
    Name     string `json:"name" validate:"required,min=2,max=100"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type GoogleLoginRequest struct {
    Token string `json:"token" validate:"required"`
}

func ValidateStruct(req interface{}) error {
    switch v := req.(type) {
    case *RegisterRequest:
        return validateRegisterRequest(v)
    case *LoginRequest:
        return validateLoginRequest(v)
    case *GoogleLoginRequest:
        return validateGoogleLoginRequest(v)
    default:
        return errors.New("unknown request type")
    }
}

func validateRegisterRequest(req *RegisterRequest) error {
    if err := validateEmail(req.Email); err != nil {
        return err
    }
    
    if err := validatePassword(req.Password); err != nil {
        return err
    }
    
    if err := validateName(req.Name); err != nil {
        return err
    }
    
    return nil
}

func validateLoginRequest(req *LoginRequest) error {
    if err := validateEmail(req.Email); err != nil {
        return err
    }
    
    if req.Password == "" {
        return errors.New("password is required")
    }
    
    return nil
}

func validateGoogleLoginRequest(req *GoogleLoginRequest) error {
    if req.Token == "" {
        return errors.New("token is required")
    }
    
    return nil
}

func validateEmail(email string) error {
    if email == "" {
        return errors.New("email is required")
    }
    
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return errors.New("invalid email format")
    }
    
    return nil
}

func validatePassword(password string) error {
    if password == "" {
        return errors.New("password is required")
    }
    
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters long")
    }
    
    return nil
}

func validateName(name string) error {
    if name == "" {
        return errors.New("name is required")
    }
    
    name = strings.TrimSpace(name)
    if len(name) < 2 {
        return errors.New("name must be at least 2 characters long")
    }
    
    if len(name) > 100 {
        return errors.New("name must be less than 100 characters long")
    }
    
    return nil
}