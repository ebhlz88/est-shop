package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ebhlz88/est-shop/common"
	"github.com/ebhlz88/est-shop/models"
	"github.com/ebhlz88/est-shop/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HandleUser(s common.Store) common.ApiFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if r.Method == "POST" {
			return CreateUser(w, r, s)

		}
		if r.Method == http.MethodGet {
			return GetUser(w, r, s)
		}
		return nil
	}
}

func HandleUserById(s common.Store) common.ApiFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if r.Method == http.MethodGet {
			return GetUserById(w, r, s)
		}
		return nil
	}
}
func HandleLogin(s common.Store) common.ApiFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if r.Method == http.MethodPost {
			return Login(w, r, s)
		}
		return nil
	}
}
func CreateUser(w http.ResponseWriter, r *http.Request, s common.Store) error {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return utils.WriteJson(w, http.StatusBadRequest, models.APIError{
			Error: err.Error(),
		})
	}
	passwordHash, err := HashPassword(user.Password)
	if err != nil {
		fmt.Println("Error Hashing password")
	}
	err = s.CreateUser(user.Name, user.UserName, passwordHash, user.Number, time.Now())
	if err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, models.APIError{
			Error: err.Error(),
		})
	}
	return utils.WriteJson(w, http.StatusCreated, models.APISuccessMessage{
		Message: "successfully created",
	})
}
func Login(w http.ResponseWriter, r *http.Request, s common.Store) error {
	type User struct {
		Username string
		Password string
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return utils.WriteJson(w, http.StatusBadRequest, models.APIError{
			Error: err.Error(),
		})
	}
	fmt.Print(user.Username)
	userFromDb, err := s.GetUserByUsername(user.Username)
	if err != nil {
		return utils.WriteJson(w, http.StatusNotFound, models.APIError{
			Error: err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password)); err != nil {
		return utils.WriteJson(w, http.StatusUnauthorized, models.APIError{
			Error: "username or password is incorrect",
		})
	}
	token, err := GenerateJwtToken([]byte(os.Getenv("JWT_SECRET")), userFromDb.ID)
	if err != nil {
		fmt.Println(err)
	}
	return utils.WriteJson(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request, s common.Store) error {
	users, err := s.GetUser()
	if err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, models.APIError{
			Error: err.Error(),
		})
	}
	return utils.WriteJson(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request, s common.Store) error {
	id := GetUserId(r)
	users, err := s.GetUserById(id)
	if err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, models.APIError{
			Error: err.Error(),
		})
	}
	return utils.WriteJson(w, http.StatusOK, users)
}

func GetUserId(r *http.Request) int {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("error converting id")
	}
	return id
}

func GenerateJwtToken(secret []byte, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(secret)
}

func ValidateJWT(tokenString string, secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Ensure token method is HS256
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func WithJWTAuth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("x-jwt-token")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := ValidateJWT(tokenStr, []byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			PermissionDenied(w)
			return
		}
		HandlerFunc(w, r)
	}
}

func PermissionDenied(w http.ResponseWriter) error {
	return utils.WriteJson(w, http.StatusUnauthorized, models.APIError{
		Error: "Permission Denied",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckHashPassword(passwordHash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}
