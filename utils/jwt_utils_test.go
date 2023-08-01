package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"goLangJwtPrac/structures"
	"testing"
)

func TestCreateAccessToken(t *testing.T) {
	user := &structures.User{
		Username:    "testuser",
		Authorities: "user",
	}

	accessToken, err := CreateAccessToken(user)
	if err != nil {
		assert.Nil(t, err, err)
	}
	fmt.Println("Access Token:", accessToken)
}

func TestVerifyAccessToken(t *testing.T) {
	var testToken = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyIiwiYXV0aG9yaXRpZXMiOiJ1c2VyIiwiZXhwIjoxNjkwODY4MjU5fQ.KmtscnUuKviJGptvW_hhYCyeV4CVUE14Rk0UsdWj2EQ-HRkOmGUGgrp08wOavCkbBHUbiQ6U3Wv2cuxMz5KHfA"
	// verifyAccessToken 테스트
	valid, err := VerifyAccessToken(testToken)
	if err != nil {
		assert.Nil(t, err, "Failed to verify token: %v", err)
	}
	fmt.Println("Token is valid:", valid)
}
