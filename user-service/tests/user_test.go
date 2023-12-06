// user_test.go
package tests

import (
	"context"
	"testing"

	pb "github.com/MikaelHans/catea/user-service/api"
	"github.com/stretchr/testify/assert"
	"github.com/MikaelHans/catea/user-service/internal/user"
)



func TestLogin_Success(t *testing.T) {
	// Set up a test case where login is successful
	server := &user.Server{}
	loginCredentials := &pb.LoginCredentials{
		Email: "mikael.hans25@gmail.com",
		Pass:  "punten",
	}

	response, err := server.Login(context.Background(), loginCredentials)
	
	// Assertions
	assert.NoError(t, err, "Login should not return an error")
	assert.NotEmpty(t, response.Token, "Token should not be empty on successful login")
}

func TestLogin_Failure(t *testing.T) {
	// Set up a test case where login fails
	server := &user.Server{}
	loginCredentials := &pb.LoginCredentials{
		Email: "nonexistent@example.com",
		Pass:  "wrong_password",
	}

	response, err := server.Login(context.Background(), loginCredentials)

	// Assertions
	assert.Error(t, err, "Login should return an error for invalid credentials")
	assert.Empty(t, response.Token, "Token should be empty on failed login")
}

func TestSignUp_Success(t *testing.T) {
	// Set up a test case where signup is successful
	server := &user.Server{}
	signupCredentials := &pb.SignupCredentials{
		Email: "newussdfdfsdfer@example.com",
		Pass:  "new_user_password",
	}

	response, err := server.SignUp(context.Background(), signupCredentials)

	// Assertions
	assert.NoError(t, err, "SignUp should not return an error")
	assert.Equal(t, "OK", response.Message, "Message should be 'OK' on successful signup")
}

func TestSignUp_Failure(t *testing.T) {
	// Set up a test case where signup fails (e.g., due to existing email)
	server := &user.Server{}
	signupCredentials := &pb.SignupCredentials{
		Email: "test@example.com", // Assuming this email already exists
		Pass:  "existing_user_password",
	}

	response, err := server.SignUp(context.Background(), signupCredentials)

	// Assertions
	assert.Error(t, err, "SignUp should return an error for existing email")
	assert.Empty(t, response.Message, "Message should be empty on failed signup")
}
