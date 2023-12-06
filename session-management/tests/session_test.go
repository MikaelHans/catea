package tests

// session_test.go

import (
	"context"
	// "encoding/json"
	"testing"
	"time"

	pb "github.com/MikaelHans/catea/session-management/api"
	session "github.com/MikaelHans/catea/session-management/internal"
	// "github.com/MikaelHans/catea/session-management/internal/repository"
	"github.com/MikaelHans/catea/session-management/pkg/structs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "github.com/MikaelHans/catea/session-management/internal/session"
)

// MockRedisClient is a mock implementation of the Redis client

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

// Set is a mock implementation of the Set method.
func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func TestSetSession_Success(t *testing.T) {
	// Create a mock Redis client
	mockClient := new(MockRedisClient)

	// Set expectations for the Set method
	mockClient.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Replace the ConnectToRedisClient function with the mock client
	// repository.ConnectToRedisClient = func() repository.RedisClient {
	// 	return mockClient
	// }

	// Create a session server instance
	server := &session.Server{}

	// Call the SetSession method
	sessionData := &pb.SessionData{
		SessionID: &pb.SessionID{Sessionid: "some_session_id"},
		// Add other necessary fields based on your actual protobuf definition
	}
	response, err := server.SetSession(context.Background(), sessionData)

	// Assertions
	assert.NoError(t, err, "SetSession should not return an error")
	assert.NotNil(t, response, "Response should not be nil")

	// Assert that the Set method on the mock client was called with the correct arguments
	mockClient.AssertCalled(t, "Set", mock.Anything, "some_session_id", mock.Anything, time.Hour)
}

func TestGetSessionInfo_Success(t *testing.T) {
	// Create a mock Redis client
	mockClient := new(MockRedisClient)

	// Set expectations for the Get method
	mockClient.On("Get", mock.Anything, mock.Anything).Return("some_data", nil)

	// Replace the ConnectToRedisClient function with the mock client
	// repository.ConnectToRedisClient = func() repository.RedisClient {
	// 	return mockClient
	// }

	// Create a session server instance
	server := &session.Server{}

	// Call the GetSessionInfo method
	sessionID := &pb.SessionID{Sessionid: "some_session_id"}
	response, err := server.GetSessionInfo(context.Background(), sessionID)

	// Assertions
	assert.NoError(t, err, "GetSessionInfo should not return an error")
	assert.NotNil(t, response, "Response should not be nil")
	assert.Equal(t, "some_data", response.Data, "Data should match the mock response")

	// Assert that the Get method on the mock client was called with the correct arguments
	mockClient.AssertCalled(t, "Get", mock.Anything, "some_session_id")
}

func TestStoreSessionDataToRedis_Success(t *testing.T) {
	// Create a mock Redis client
	mockClient := new(MockRedisClient)

	// Set expectations for the Set method
	mockClient.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Replace the ConnectToRedisClient function with the mock client
	// repository.ConnectToRedisClient = func() repository.RedisClient {
	// 	return mockClient
	// }

	// Call the storeSessionDataToRedis function directly (assuming it's not exposed in your package)
	currentTime := time.Now()

	// Format the time as a string
	timeString := currentTime.Format("2006-01-02 15:04:05")
	memberData := structs.Member{
		Email:        "test@example.com",
		Firstname:    "John",
		Lastname:     "Doe",
		Member_Since: timeString,
	}
	token := "some_token"
	err := session.StoreSessionDataToRedis(memberData, token, context.Background())

	// Assertions
	assert.NoError(t, err, "storeSessionDataToRedis should not return an error")

	// Assert that the Set method on the mock client was called with the correct arguments
	mockClient.AssertCalled(t, "Set", mock.Anything, "some_token", mock.Anything, time.Hour)
}
