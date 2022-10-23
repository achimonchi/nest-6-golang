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
	tokString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3N1ZWQiOiIyMDIyLTEwLTIzVDA5OjI5OjUzLjQyMDEyNCswNzowMCIsInBheWxvYWQiOnsidXNlcl9pZCI6IjEiLCJlbWFpbCI6InJleWhhbkBnbWFpbC5jb20ifX0.jK_wETyTKE1bAwEOl0rMy3zOwHhTh1SCuSooyIaqdw4"

	myTok, err := VerifyToken(tokString)

	assert.Nil(t, err)

	fmt.Println("mytoken :", myTok)
}
