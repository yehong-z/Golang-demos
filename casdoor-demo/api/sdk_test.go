package api

import (
	"fmt"
	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"testing"
)

func TestGetToken(t *testing.T) {
	Init()
	code := "809b11ff9eb77e1d3b30"
	state := "casdoor"
	token, err := auth.GetOAuthToken(code, state)
	if err != nil {
		t.Error(err)
	}

	acToken := token.AccessToken
	fmt.Println(acToken)
	reToken := token.RefreshToken

	token, err = auth.RefreshOAuthToken(reToken)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token.AccessToken)

	Jwt, _ := auth.ParseJwtToken(token.AccessToken)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(Jwt.ExpiresAt)
}
