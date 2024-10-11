package usecase

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/Thehien0811/order-product-microservice/internal/user/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func (uc implUseCase) SignIn(ctx context.Context, input SignInInput) (DetailUser, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return DetailUser{}, errors.New("Hash password failed")
	}

	opt := repository.CreateUserOption{
		Username: input.Username,
		Password: string(hashedPassword),
	}
	u, err := uc.repo.CreateUser(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}

	res := DetailUser{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password,
	}
	return res, nil
}

func (uc implUseCase) SignUp(ctx context.Context, input SignUpInput) (string, error) {
	user, err := uc.repo.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return "", errors.New("Username doesn't exist")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("Username doesn't exist")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: input.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		log.Fatal("JWT_SECRET not set in .env")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", errors.New("Error generating token")
	}
	return tokenString, nil

}
