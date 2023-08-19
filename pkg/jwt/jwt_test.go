package jwt

import (
	"fmt"
	"testing"
)

func TestJWT_CreateToken(t *testing.T) {
	jwt := NewJWT([]byte("signkey"))
	fmt.Println(jwt)
	token, err := jwt.CreateToken(CustomClaims{ID: 1})
	fmt.Println(token, err)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlck5hbWUiOiJseWoifQ.FVeQiFJHOPCa8U_Y62OH4YohdfPisurEBcuNiKVaoa4
}

func TestJWT_ParseToken(t *testing.T) {
	jwt := NewJWT([]byte("signkey"))
	fmt.Println(jwt)
	claims, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlck5hbWUiOiJseWoifQ.FVeQiFJHOPCa8U_Y62OH4YohdfPisurEBcuNiKVaoa4")
	fmt.Println(claims, err)
}
