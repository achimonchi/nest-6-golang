package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	tok := Token{
		UserId: "1",
		Email:  "reyhan@gmail.com",
	}

	myTok, err := CreateToken(&tok)
	if err != nil {
		t.Errorf("fail to generate token with error :%s", err.Error())
		return
	}

	fmt.Println("mytoken :", myTok)
}

func TestVerifyToken(t *testing.T) {
	tokString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOiIxIiwiZW1haWwiOiJyZXloYW5AZ21haWwuY29tIn19.Y6KuJ5HcT8Tk2ZkWPJ7679xgVEeyLBGADsdiJ23NJi0"

	myTok, err := VerifyToken(tokString)

	assert.Nil(t, err)
	// if err != nil {
	// 	t.Errorf("fail to verify token with error :%s", err.Error())
	// 	return
	// }

	fmt.Println("mytoken :", myTok)
}
