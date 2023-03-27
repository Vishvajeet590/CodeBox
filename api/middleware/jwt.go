package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

type CustomClaims struct {
	UserId   string   `json:"userId"`
	UserType string   `json:"userType"`
	Scopes   []string `json:"scopes"`
	jwt.StandardClaims
}

func TokenAuthMiddleware(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(scope, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func TokenValid(api_scope string, c *gin.Context) error {
	r := c.Request
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return fmt.Errorf("not Enough Permission")
	}

	// Setting Username to the context
	c.Set("userId", claims["userId"])

	for _, role := range scopes {
		if role == api_scope {
			return nil
		}
	}

	return fmt.Errorf("Not Enough Permission")
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	token_string := ExtractToken(r)

	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bear_token := r.Header.Get("Authorization")
	str_rrr := strings.Split(bear_token, " ")
	if len(str_rrr) == 2 {
		return str_rrr[1]
	}
	return ""
}

func GenerateToken(userId string, userType string, scope []string) (string, error) {
	mySigningKey := []byte(os.Getenv("ACCESS_SECRET"))
	//todo change key
	claims := CustomClaims{
		userId,
		userType,
		scope,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(40 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err

}
